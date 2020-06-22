package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}
