package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Add(fileStorage string, items []string) {
	file, err := os.OpenFile(fileStorage, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(file, strings.Join(items, " "))

	if err != nil {
		log.Fatal(err)
	}

}
