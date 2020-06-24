package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func main() {

	var lFlag = flag.Bool("l", false, "File Attribute")
	var rFlag = flag.Bool("r", false, "sort reverse")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if *rFlag {
		length := len(fileInfos)

		reverseFileInfos := []os.FileInfo{}
		for i := length - 1; i >= 0; i-- {
			reverseFileInfos = append(reverseFileInfos, fileInfos[i])
		}

		fileInfos = reverseFileInfos
	}

	if *lFlag {
		fmt.Println("合計", len(fileInfos))
		for _, fileInfo := range fileInfos {

			fmt.Printf("%v", fileInfo.Mode())

			fmt.Printf(" %v", fileInfo.Size())

			if stat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
				uid := strconv.Itoa(int(stat.Uid))
				u, err := user.LookupId(uid)
				if err != nil {
					fmt.Printf(" %v", uid)
				} else {
					fmt.Printf(" %v", u.Username)
				}
				gid := strconv.Itoa(int(stat.Gid))
				g, err := user.LookupGroupId(gid)
				if err != nil {
					fmt.Printf(" %v", gid)
				} else {
					fmt.Printf(" %v", g.Name)
				}
			}

			fmt.Printf(" %v ", fileInfo.ModTime().Format("Jan 2 15:04"))

			fmt.Println(fileInfo.Name())
		}
		return
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}
