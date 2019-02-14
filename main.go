package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/spf13/afero"
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
	deleteFile(name)

	// ----------------------------------- //
	fmt.Println("----------------- new ReadFile interface -----------------")
	testFS := &afero.MemMapFs{}
	fsutil := &afero.Afero{Fs: testFS}

	filename := "this_exists.txt"
	testFS.Create(filename)
	contents, err := fsutil.ReadFile(filename)
	//t.Errorf("test1----- %v %v", contents, err)
	if err != nil {
		fmt.Println("ReadFile : error unexpected, can't read file [", filename, "]")
	}
	fmt.Println("Empty read. contents = ", contents)

	data := "Programming today is a race between software engineers striving to " +
		"build bigger and better idiot-proof programs, and the Universe trying " +
		"to produce bigger and better idiots. So far, the Universe is winning."

	if err := fsutil.WriteFile(filename, []byte(data), 0644); err != nil {
		fmt.Println("WriteFile ", filename, " : ", err)
	}

	contents, err = fsutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile : error unexpected, can't read file [", filename, "]")
	}
	//fmt.Println("Read again. contents = ", contents)
	fmt.Println("string of data is -> ", string(contents[:]))

	//..//
	fs2 := &tfs.FileSystem{}
	fs2.Initialize(fsutil)

	var res1, res2 string = "", ""
	// .. 1 read .. //
	fmt.Println("******************************************")
	raw, err = fs2.Get(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		data, ok := raw.([]byte)
		if ok {
			//fmt.Println("check type is ok (type is []byte)")
			//fmt.Println("data is ", data)
			res1 = string(data[:])
			fmt.Println("string of data is -> ", res1)
		} else {
			fmt.Println("check type is not ok")
		}
	}

	testFS.Remove(filename) // ignore error

	// .. 2 read after remove file .. //
	fmt.Println("******************************************")
	raw2, err := fs2.Get(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		data, ok := raw2.([]byte)
		if ok {
			//fmt.Println("check type is ok (type is []byte)")
			//fmt.Println("data is ", data)
			res2 = string(data[:])
			fmt.Println("string of data is -> ", res2)
		} else {
			fmt.Println("check type is not ok")
		}
	}

	if res1 != res2 {
		fmt.Println("res1 != res2")
	}

}
