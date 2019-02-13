package main

import (
	"fmt"
	"io/ioutil"
	"log"

	tfs "github.com/tronfs/filesystem"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("hello FileSystem")
	fs := &tfs.FileSystem{}

	name := "setting.txt"
	raw, err := fs.Get(name)
	if err != nil {
		log.Println("[first] err: ", err)
	} else {
		log.Println("[first] raw: ", raw)
	}

	d1 := []byte("hello\ngo\n")
	err = ioutil.WriteFile(name, d1, 0644)
	check(err)

	for i := 0; i < 2; i++ {
		log.Println("-----------------", i+1, "-----------------")
		raw, err = fs.Get(name)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(raw)
		}
	}

}
