package quotes

import (
	"io/ioutil"
	"math/rand"
	"strings"
)

// Quotes represents quotes as kind type of []string to implement Rand method
type Quotes []string

func (q Quotes) Rand() string {
	key := rand.Int31n(int32(len(q)) - 1)
	return q[key]
}

// LoadQuotes loads file with quotes and loads it to Quotes struct
func LoadQuotes(filePath string) (Quotes, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
