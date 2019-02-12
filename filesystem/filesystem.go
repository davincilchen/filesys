package filesystem

import (
	"time"

	"github.com/davecgh/go-spew/spew"

	cache "github.com/patrickmn/go-cache"
)

const (
	// For use with functions that take an expiration time.
	NoExpiration time.Duration = -1
	// For use with functions that take an expiration time. Equivalent to
	// passing in the same expiration duration as was given to New() or
	// NewFrom() when the cache was created (e.g. 5 minutes.)
	DefaultExpiration time.Duration = 0
)

type FileSystem struct {
	cache *cache.Cache
}

func (fs *FileSystem) init(k string, x interface{}, d time.Duration) {
	fs.cache = cache.New(DefaultExpiration, 0)
	spew.Dump(fs.cache)
}
