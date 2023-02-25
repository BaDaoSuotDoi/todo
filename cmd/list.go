package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func List(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	bf := bufio.NewReader(file)
	for i := 1; ; i++ {
		line, _, err := bf.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		if string(line) == " " {
			break
		}
		fmt.Printf("%d. %s\n", i, line)
	}
}
