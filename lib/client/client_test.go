package client

import (
	"bufio"
	"net"
	"test-ddos/lib/protocol"
	"testing"

	"github.com/bwesterb/go-pow"
)

const testQuote = "Quote"

func TestClientVerify(t *testing.T) {
	testClient := NewClient()

	srv, cl := net.Pipe()
	defer srv.Close()
	defer cl.Close()

	go func() {
		r := bufio.NewReader(srv)

		challenge := pow.NewRequest(32, []byte{1, 2, 3, 4, 5, 6, 7, 8})

		// Send test challenge
		err := protocol.Write(srv, challenge)
		if err != nil {
			t.Error(err)
		}

		// Read proof
		_, err = protocol.Read(r)
		if err != nil {
			t.Error(err)
		}

		// Send test quote
		protocol.Write(srv, testQuote)
	}()

	q, err := testClient.handleData(cl)
	if err != nil {
		t.Error(err)
	}

	if q != testQuote {
		t.Errorf("Unexpected quote processed in result, got - %v, require - %v", q, testQuote)
	}
}
