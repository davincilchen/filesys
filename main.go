package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	tfs "github.com/tronfs/filesystem"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func deleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		fmt.Println("==> fail deleting file  (", path, ")")
		return
	}

	fmt.Println("==> done deleting file (", path, ")")
}

func main() {
	fmt.Println("Hello FileSystem")
	fs := &tfs.FileSystem{}

	name := "setting.txt"
	deleteFile(name)

	//.. testing if no file ..//
	raw, err := fs.Get(name)
	if err != nil {
		fmt.Println("[first] err: ", err)
	} else {
		fmt.Println("[first] raw: ", raw)
	}

	d1 := []byte("hello\ngo\n")
	err = ioutil.WriteFile(name, d1, 0644)
	check(err)

	//.. test twice ..//
	for i := 0; i < 2; i++ {
		fmt.Println("-----------------", i+1, "-----------------")
		raw, err = fs.Get(name)
		if err != nil {
			fmt.Println(err)
		} else {

			switch raw.(type) {
			case int:
				fmt.Println("type= int")
			case float64:
				fmt.Println("type= float64")
			case []byte:
				fmt.Println("type= []byte")
			// case []uint8:
			// 	fmt.Println("type= []uint8")
			default:
				fmt.Println("type= unkonw")
			}

			fmt.Println("reflect.TypeOf(raw)=", reflect.TypeOf(raw))
			fmt.Println("raw= ", raw)

			data, ok := raw.([]byte)
			if ok {
				fmt.Println("check type is ok (type is []byte)")
				fmt.Println("data is ", data)
				fmt.Println("string of data is -> ", string(data[:]))
			} else {
				fmt.Println("check type is not ok")
			}
		}
	}

}
