package main

import (
	"log"
	"os"
	"path"
	"strconv"
	"todo/cmd"
)

func main() {
	fileName := ".todo"
	dirCur, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	action := os.Args[1]

	fileStorage := path.Join(dirCur, fileName)

	switch action {
	case "add":
		cmd.Add(fileStorage, os.Args[2:])
		break
	case "list":
		cmd.List(fileName)
		break
	case "del":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		cmd.Delete(fileName, id)

	}
	// data, err := os.Stat(path.Join(dirCur, fileName))

	// file, err := os.OpenFile(path.Join(dirCur, fileName), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()
	// fmt.Fprint(file, "Hello")

	// var result string

	// fmt.Fscan(file, &result)
	// fmt.Println(result)

	// fmt.Println(os.Getenv("HOME"))
}
