package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	// Stat retorna informações do arquivo
	// Se não existir arquivo ele retorna um erro

	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("File Size Bytes: ", fileInfo.Size())
	fmt.Println("Last time modified: ", fileInfo.ModTime())
}
