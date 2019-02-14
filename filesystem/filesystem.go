package filesystem

//TODO: lock for rontine if necessarry
import (
	"fmt"
	"io/ioutil"
	"log"

	cache "github.com/patrickmn/go-cache"
)

//File is a
type File interface {
	ReadFile(filename string) ([]byte, error)
}

//Ioutil is a
type Ioutil struct {
}

//ReadFile is a
func (io *Ioutil) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

//FileSystem is a
type FileSystem struct {
	cache *cache.Cache
	file  File
}

const (
	//DefaultCacheExpiration is a
	DefaultCacheExpiration = cache.NoExpiration
)

//Initialize is a
func (fs *FileSystem) Initialize(file File) error {
	fs.cache = cache.New(DefaultCacheExpiration, 0)
	if file == nil {
		fs.file = &Ioutil{}
	} else {
		fs.file = file
	}
	//spew.Dump(fs.cache)
	return nil
}

//Reinitialize is a
func (fs *FileSystem) Reinitialize() error {
	if fs.cache == nil {
		fs.Initialize(nil)
	} else {
		fs.cache.Flush()
	}
	return nil
}

//Uninitialize is a
func (fs *FileSystem) Uninitialize() error {
	fs.cache = nil

	return nil
}

//CheckCache is a
func (fs *FileSystem) CheckCache() (bool, error) {

	if fs.cache == nil {
		return false, fmt.Errorf("Cache is nil")
	}

	return true, nil
}

//CheckOrInitCache is a
func (fs *FileSystem) CheckOrInitCache() (bool, error) {

	found, err := fs.CheckCache()
	if err != nil {
		//log.Println(err)
		fs.Initialize(nil)
	}

	return found, nil
}

//Get is a
func (fs *FileSystem) Get(key string) (interface{}, error) {

	fs.CheckOrInitCache()
	// _, err := fs.CheckCache()
	// if err != nil {
	// 	return nil, err
	// }

	data, found := fs.cache.Get(key)
	if found {
		return data, nil
	}

	log.Println("Data not found for key [", key, "]", "Reload from file.")

	if fs.file == nil {
		return nil, fmt.Errorf("No ReadFile Interface")
	}
	raw, err := fs.file.ReadFile(key) //raw, err := ioutil.ReadFile(key)
	if err != nil {
		return nil, err
	}

	fs.cache.Set(key, raw, cache.DefaultExpiration)

	return raw, nil
}
