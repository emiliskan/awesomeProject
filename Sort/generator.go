package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const MaxNum int = 999_999_999
const ChunkSize int = 10_000

func generatePhoneFiles() {
	num := 79_000_000_000

	phonesFile, _ := os.OpenFile("phones", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	var phonesChunk [ChunkSize]string

	chunkCount := 1
	chunkPos := 0
	dataWriter := bufio.NewWriter(phonesFile)

	for i := 0; i < MaxNum; i++ {

		if i/chunkCount >= ChunkSize {
			writeChunkToFile(dataWriter, phonesChunk)
			chunkCount++
			chunkPos = 0
		}

		phone := "+" + strconv.Itoa(num)
		phonesChunk[chunkPos] = phone
		num++
		chunkPos++
	}

	writeChunkToFile(dataWriter, phonesChunk)

	err := dataWriter.Flush()
	if err != nil {
		fmt.Print(err)
	}
	phonesFile.Close()

}

func writeChunkToFile(dataWriter *bufio.Writer, chunk [ChunkSize]string) {
	rand.Shuffle(len(chunk), func(i, j int) {
		chunk[i], chunk[j] = chunk[j], chunk[i]
	})
	data := strings.Join(chunk[:], "\n")
	_, _ = dataWriter.WriteString(data + "\n")
}
