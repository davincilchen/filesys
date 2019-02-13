package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

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

			switch raw.(type) {
			case int:
				log.Println("type= int")
			case float64:
				log.Println("type= float64")
			case []byte:
				log.Println("type= []byte")
			// case []uint8:
			// 	log.Println("type= []uint8")
			default:
				log.Println("type= unkonw")
			}

			fmt.Println("reflect.TypeOf(raw)=", reflect.TypeOf(raw))
			log.Println("raw= ", raw)

			data, ok := raw.([]byte)
			if ok {
				log.Println("check type is ok (type is []byte)")
				log.Println("data is ", data)
				log.Println("string of data is -> ", string(data[:]))
			} else {
				log.Println("check type is not ok")
			}
		}
	}

}
