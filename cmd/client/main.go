package main

import (
	"flag"
	"fmt"
	"log"

	"test-ddos/lib/client"
)

func main() {
	hostAddr := flag.String("addr", "server:8999", "Host of TCP server for word-of-wisdom")
	flag.Parse()

	quote, err := client.NewClient().ConnectAndGetQuote(*hostAddr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Quote from verification - ", quote)
}
