package filesystem

import (
	"testing"
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

	fs.Initialize()

	hasCache, err = fs.CheckCache()
	if err != nil {
		t.Error("Shold not have error because of cache is initialized")
	}
	if hasCache != true {
		t.Error("[hasCache] Shold be true because of cache is initialized")
	}
}

//IO //TODO: close IO?
func TestFileSystem_Reinitialize(t *testing.T) {
	// tests := []struct {
	// 	name    string
	// 	fs      *FileSystem
	// 	wantErr bool
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if err := tt.fs.Reinitialize(); (err != nil) != tt.wantErr {
	// 			t.Errorf("FileSystem.Reinitialize() error = %v, wantErr %v", err, tt.wantErr)
	// 		}
	// 	})
	// }
}

func TestFileSystem_Uninitialize(t *testing.T) {

	fs := &FileSystem{}

	fs.Initialize()
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

//IO //TODO: close IO?
func TestFileSystem_Get(t *testing.T) {

}
