package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"test-ddos/lib/quotes"
	"test-ddos/lib/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	hostAddr := flag.String("listen", ":8999", "Host of TCP server with port, format is ip:port")
	quotesFilePath := flag.String("quotes", "resources/quotes.txt", "Filepath to a file with quotes")
	flag.Parse()

	qs, err := quotes.LoadQuotes(*quotesFilePath)
	if err != nil {
		log.Fatalf("Can't load quotes file, %e", err)
	}

	fmt.Println("Running tcp server on addr:", *hostAddr)

	err = server.NewServer(qs).SetupAndServe(*hostAddr)
	if err != nil {
		log.Fatalln(err)
	}
}
