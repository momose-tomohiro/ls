package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	var lFlag = flag.Bool("l", false, "File Attribute")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if *lFlag {
		fmt.Println("total", len(fileInfos))
		for _, fileInfo := range fileInfos {

			fmt.Printf("%v", fileInfo.Mode())

			fmt.Printf(" %v", fileInfo.Size())

			fmt.Printf(" %v", os.Getgid())

			fmt.Printf(" %v ", fileInfo.ModTime().Format("Jan 2 15:04"))

			fmt.Println(fileInfo.Name())
		}
		return
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}
