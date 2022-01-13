package server

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"net"
	"test-ddos/lib/protocol"
	"test-ddos/lib/quotes"

	"github.com/bwesterb/go-pow"
)

// Server represents methods to setup server with resolving word-of-widsmon algo
type Server struct {
	Quotes        quotes.Quotes
	POWDifficulty uint32
}

// NewServer creates default server
func NewServer(qs quotes.Quotes) Server {
	return Server{qs, 32}
}

// SetupAndServe setup a tcp server and handles client data
func (s Server) SetupAndServe(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Can't accept connect, err - %v", err)
			continue
		}

		go func() {
			if err := s.handleClient(conn); err != nil {
				fmt.Printf("handle connect failure, err - %v", err)
			}
			conn.Close()
		}()
	}
}

// handleClient handles client connection and procedure of pow verification
func (s Server) handleClient(conn net.Conn) error {
	r, w := bufio.NewReader(conn), conn

	challenge := pow.NewRequest(s.POWDifficulty, s.generateNonce())
	if err := protocol.Write(w, challenge); err != nil {
		return err
	}

	proof, err := protocol.Read(r)
	if err != nil {
		return err
	}

	verified, err := pow.Check(challenge, proof, nil)
	if err != nil {
		return err
	}

	if !verified {
		return fmt.Errorf("can't verify exchange way")
	}

	return protocol.Write(w, s.Quotes.Rand())
}

func (s Server) generateNonce() []byte {
	nonce := make([]byte, 8)
	rand.Read(nonce)
	return nonce
}
