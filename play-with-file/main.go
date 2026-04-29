package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Printf("error open file %s", err)
		return
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var s = make([]byte, 3)

	for {
		_, errRead := buf.Read(s)
		if errRead != nil {
			if errRead == io.EOF {
				log.Fatal("eof ", errRead)
				break
			}
			log.Fatal("Error read file")
			break
		}
		log.Printf("the file containing: %v", string(s))
	}
}

func testf(a int, f func(a int)) {
	f(a)
}

func ff(a int) {
	log.Println("Test", a)
}
