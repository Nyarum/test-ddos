package protocol

import (
	"bufio"
	"strings"
)

func Read(r *bufio.Reader) (string, error) {
	s, err := r.ReadString(delim)
	if err != nil {
		return "", err
	}

	s = strings.TrimSuffix(s, string(delim))

	return s, nil
}
