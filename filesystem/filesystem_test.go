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
	// tests := []struct {
	// 	name    string
	// 	fs      *FileSystem
	// 	want    bool
	// 	wantErr bool
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := tt.fs.CheckOrInitCache()
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("FileSystem.CheckOrInitCache() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if got != tt.want {
	// 			t.Errorf("FileSystem.CheckOrInitCache() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestFileSystem_Get(t *testing.T) {
	// type args struct {
	// 	key string
	// }
	// tests := []struct {
	// 	name    string
	// 	fs      *FileSystem
	// 	args    args
	// 	want    interface{}
	// 	wantErr bool
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := tt.fs.Get(tt.args.key)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("FileSystem.Get() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("FileSystem.Get() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }

}
