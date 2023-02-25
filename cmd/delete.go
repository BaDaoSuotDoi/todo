package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func Delete(fileName string, id int) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	temp := ""

	br := bufio.NewReader(file)

	for i := 1; i < id; i++ {
		data, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		temp += string(data) + "\n"
	}

	br.ReadLine()
	for i := id + 1; ; i++ {
		data, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		temp += string(data) + "\n"
	}

	temp += " "

	fmt.Println(temp)
	file.WriteAt([]byte(temp), 0)
	// fmt.Fprint(file, temp)

}
