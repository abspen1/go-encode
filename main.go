package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abspen1/go-encode/decode"

	"github.com/abspen1/go-encode/encode"
)

func main() {
	scheme := os.Args[1]
	sortFunc := os.Args[2]

	start := time.Now()
	if scheme == "encode" {
		encode.Run(sortFunc)
	} else if scheme == "decode" {
		decode.Run(sortFunc)
	} else {
		fmt.Println("Invalid command", scheme)
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
