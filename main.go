package main

import (
	"fmt"

	tfs "github.com/tronfs/filesystem"
)

func main() {
	fmt.Println("hello world")

	fs := &tfs.FileSystem{}

	fs.Reinitialize()

	fs.Reinitialize()
}
