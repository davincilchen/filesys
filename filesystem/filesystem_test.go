package filesystem

import (
	"fmt"
	"testing"

	"github.com/afero"
)

func TestFileSystem_Initialize(t *testing.T) {

	fs := &FileSystem{}

	hasCache, err := fs.CheckCache()
	//log.Println(hasCache, err)
	if err == nil {
		t.Error("Shold have error because of cache is not initialized")
	}
	if hasCache == true {
		t.Error("[hasCache] Shold be false because of cache is not initialized")
	}

	fs.Initialize(nil)

	hasCache, err = fs.CheckCache()
	if err != nil {
		t.Error("Shold not have error because of cache is initialized")
	}
	if hasCache != true {
		t.Error("[hasCache] Shold be true because of cache is initialized")
	}
}

func TestFileSystem_Reinitialize(t *testing.T) {

	testFS := &afero.MemMapFs{}
	fsutil := &afero.Afero{Fs: testFS}
	filename := "this_exists.txt"
	testFS.Create(filename)

	data := "Programming today is a race between software engineers striving to " +
		"build bigger and better idiot-proof programs, and the Universe trying " +
		"to produce bigger and better idiots. So far, the Universe is winning."

	if err := fsutil.WriteFile(filename, []byte(data), 0644); err != nil {
		t.Error("WriteFile ", filename, " : ", err)
	}

	_, err := fsutil.ReadFile(filename)
	if err != nil {
		t.Error("ReadFile : error unexpected, can't read file [", filename, "]")
	}

	//..//
	fs := &FileSystem{}
	fs.Initialize(fsutil)

	raw, err := fs.Get(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		data, ok := raw.([]byte)
		if ok {
			res1 := string(data[:])
			fmt.Println("get data before Uninitialize() is -> ", res1)
		} else {
			t.Error("check type is not ok")
		}
	}

	testFS.Remove(filename) // ignore error

	// --------------------------------------- //
	fs.Uninitialize()
	// --------------------------------------- //
	_, err = fs.Get(filename)
	if err != nil {
		t.Error("Filesystem.Get() should have error after Filesystem.Uninitialize() and remove file.")
	}
}

func TestFileSystem_Uninitialize(t *testing.T) {

	fs := &FileSystem{}

	fs.Initialize(nil)
	hasCache, err := fs.CheckCache()
	if err != nil {
		t.Error("Shold not have error because of cache is initialized")
	}
	if hasCache != true {
		t.Error("[hasCache] Shold be true because of cache is initialized")
	}

	fs.Uninitialize()
	hasCache, err = fs.CheckCache()

	if err == nil {
		t.Error("Shold have error because of cache is uninitialized")
	}
	if hasCache == true {
		t.Error("[hasCache] Shold be false because of cache is uninitialized")
	}
}

func TestFileSystem_CheckCache(t *testing.T) {

	TestFileSystem_Initialize(t)
}

func TestFileSystem_CheckOrInitCache(t *testing.T) {

	fs := &FileSystem{}
	fs.CheckOrInitCache()
	hasCache, err := fs.CheckCache()
	if err != nil {
		t.Error("Shold not have error because of cache is initialized")
	}
	if hasCache != true {
		t.Error("[hasCache] Shold be true because of cache is initialized")
	}
}

func TestFileSystem_Get(t *testing.T) {
	testFS := &afero.MemMapFs{}
	fsutil := &afero.Afero{Fs: testFS}

	filename := "this_exists.txt"
	testFS.Create(filename)

	data := "Programming today is a race between software engineers striving to " +
		"build bigger and better idiot-proof programs, and the Universe trying " +
		"to produce bigger and better idiots. So far, the Universe is winning."

	if err := fsutil.WriteFile(filename, []byte(data), 0644); err != nil {
		t.Error("WriteFile ", filename, " : ", err)
	}

	_, err := fsutil.ReadFile(filename)
	if err != nil {
		t.Error("ReadFile : error unexpected, can't read file [", filename, "]")
	}

	//..//
	fs2 := &FileSystem{}
	fs2.Initialize(fsutil)

	var res1, res2 string = "", ""
	// .. 1 read .. //
	fmt.Println("******************************************")
	raw, err := fs2.Get(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		data, ok := raw.([]byte)
		if ok {
			res1 = string(data[:])
			//fmt.Println("string of data is -> ", res1)
		} else {
			t.Error("check type is not ok")
		}
	}

	testFS.Remove(filename) // ignore error

	// .. 2 read after remove file .. //
	fmt.Println("******************************************")
	raw2, err := fs2.Get(filename)
	if err != nil {
		t.Error("Filesystem get error:", err)
	} else {
		data, ok := raw2.([]byte)
		if ok {
			res2 = string(data[:])
			//fmt.Println("string of data is -> ", res2)
		} else {
			t.Error("check type is not ok")
		}
	}

	if res1 != res2 {
		t.Error("res1 != res2 after file remove. Remove file should not change cache. res1 will equal res2.")
	}
}
