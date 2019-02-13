package filesystem

import (
	"github.com/davecgh/go-spew/spew"

	cache "github.com/patrickmn/go-cache"
)

//FileSystem is a
type FileSystem struct {
	cache *cache.Cache
}

//Initialize is a
func (fs *FileSystem) Initialize() error {
	fs.cache = cache.New(cache.NoExpiration, 0)
	spew.Dump(fs.cache)
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
