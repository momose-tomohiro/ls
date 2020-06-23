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
		fmt.Println("合計", len(fileInfos))
		for _, fileInfo := range fileInfos {

			fmt.Printf("%v", fileInfo.Mode())

			fmt.Printf(" %v", fileInfo.Size())

			if stat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
				uid := strconv.Itoa(int(stat.Uid))
				u, err := user.LookupId(uid)
				if err != nil {
					fmt.Printf(uid)
				} else {
					fmt.Printf(u.Username)
				}
				gid := strconv.Itoa(int(stat.Gid))
				g, err := user.LookupGroupId(gid)
				if err != nil {
					fmt.Printf(gid)
				} else {
					fmt.Printf(g.Name)
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
