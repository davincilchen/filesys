package filesystem

import (
	"github.com/davecgh/go-spew/spew"

	cache "github.com/patrickmn/go-cache"
)

type FileSystem struct {
	cache *cache.Cache
}

func (fs *FileSystem) initialize() error {
	fs.cache = cache.New(cache.NoExpiration, 0)
	spew.Dump(fs.cache)
	return nil
}

func (fs *FileSystem) Reinitialize() error {
	if fs.cache == nil {
		fs.initialize()
	} else {
		fs.cache.Flush()
	}
	return nil
}

func (fs *FileSystem) uninitialize() error {
	fs.cache = nil
	return nil
}
