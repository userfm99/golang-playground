package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/h2non/filetype.v1"
)

func main() {
	file, err := ioutil.ReadFile("../data.csv")
	if err != nil {
		log.Panic(err)
	}

	whatFile, err := filetype.MatchFile("../test.txt")
	if err != nil {
		log.Panic(err)
	}

	log.Println("Match file", whatFile)

	isCsv := filetype.Is(file, "csv")
	if !isCsv {
		log.Println("file is not csv")
		return
	}
	log.Println("file is csv")
}
