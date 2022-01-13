package client

import (
	"bufio"
	"net"
	"test-ddos/lib/protocol"

	"github.com/bwesterb/go-pow"
)

// Client represents methods to connect and resolve word-of-widsmon algo
type Client struct {
}

// NewClient creates default client
func NewClient() Client {
	return Client{}
}

// ConnectAndGetQuote connects to tcp server and handles server data
func (c Client) ConnectAndGetQuote(addr string) (string, error) {
	ln, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer ln.Close()

	return c.handleData(ln)
}

// handleData handles client connection and procedure of pow verification
func (c Client) handleData(conn net.Conn) (string, error) {
	r, w := bufio.NewReader(conn), conn

	challenge, err := protocol.Read(r)
	if err != nil {
		return "", err
	}

	proof, err := pow.Fulfil(challenge, nil)
	if err != nil {
		return "", err
	}

	err = protocol.Write(w, proof)
	if err != nil {
		return "", err
	}

	return protocol.Read(r)
}
