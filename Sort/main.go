package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	generatePhoneFiles()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
