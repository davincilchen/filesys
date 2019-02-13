package filesystem

//TODO: lock for rontine if necessarry
import (
	"fmt"
	"io/ioutil"
	"log"

	cache "github.com/patrickmn/go-cache"
)

//FileSystem is a
type FileSystem struct {
	cache *cache.Cache
}

const (
	//DefaultCacheExpiration is a
	DefaultCacheExpiration = cache.NoExpiration
)

//Initialize is a
func (fs *FileSystem) Initialize() error {
	fs.cache = cache.New(DefaultCacheExpiration, 0)
	//spew.Dump(fs.cache)
	return nil
}

//Reinitialize is a
func (fs *FileSystem) Reinitialize() error {
	if fs.cache == nil {
		fs.Initialize()
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
		fs.Initialize()
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

	log.Println("Data not found")

	raw, err := ioutil.ReadFile(key)
	if err != nil {
		//panic(err)
		return nil, err
	}

	fs.cache.Set(key, raw, cache.DefaultExpiration)

	return raw, nil
}
