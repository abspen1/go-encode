package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abspen1/go-encode/encode"
)

func main() {
	arg := os.Args[1]
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	lines := 0
	start := time.Now()
	for scanner.Scan() {
		str = scanner.Text()
		if len(str) > 0 {
			if lines != 0 {
				fmt.Println()
			} else {
				lines++
			}
			encode.Run(str, arg)
		} else {
			fmt.Println()
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
