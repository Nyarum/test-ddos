package protocol

import "net"

func Write(w net.Conn, s string) error {
	data := append([]byte(s), delim)
	_, err := w.Write(data)

	return err
}
