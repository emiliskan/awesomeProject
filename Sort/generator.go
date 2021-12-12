package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const MaxNum int = 999_999_999
const ChunkSize int = 10_000_000

func generatePhoneFiles() {
	num := 1_000_000_000

	phonesFile, _ := os.OpenFile("phones", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	var phonesChunk []string

	dataWriter := bufio.NewWriter(phonesFile)

	for i := 0; i < MaxNum; i++ {

		phone := "+79" + strconv.Itoa(num)[1:]
		phonesChunk = append(phonesChunk, phone)

		if len(phonesChunk) > ChunkSize {
			writeChunkToFile(dataWriter, phonesChunk)
			phonesChunk = nil
		}

		num++
	}

	writeChunkToFile(dataWriter, phonesChunk)

	err := dataWriter.Flush()
	if err != nil {
		return
	}
	phonesFile.Close()

}

func writeChunkToFile(dataWriter *bufio.Writer, chunk []string) {
	rand.Shuffle(len(chunk), func(i, j int) {
		chunk[i], chunk[j] = chunk[j], chunk[i]
	})
	data := strings.Join(chunk, "\n")
	_, _ = dataWriter.WriteString(data + "\n")
}
