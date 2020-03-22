package resource

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

//ErrInvalidResource occurs when resource is not file
var ErrInvalidResource = errors.New("resource is not file")

//ErrInvalidDirectoryStructure occurs when there are missing some folders. Probably user is not in geography root
var ErrInvalidDirectoryStructure = errors.New("invalid directory structure")

type Existence interface {
	Exist() bool
}

//Resource represents path to given resource (file or directory only)
type Resource struct {

	//Directory should end with "/" symbol
	Directory string

	//FileName should contain extension
	FileName string
}

//New returns new Resource
func New(dir, file string) Resource {
	return Resource{
		Directory: dir,
		FileName:  file,
	}
}

//Write save p bytes to given location
func (r Resource) Write(p []byte) (n int, err error) {
	if !r.isFile() {
		return 0, ErrInvalidResource
	}

	return len(p), ioutil.WriteFile(path.Clean(r.Directory+r.FileName), p, 0741)
}

//Exist checks if resource exists
func (r Resource) Exist() bool {
	resourcePath := ""

	if r.isDir() {
		resourcePath = r.Directory
	}

	if r.isFile() {
		resourcePath = r.Directory + r.FileName
	}

	_, err := os.Stat(resourcePath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

//DirExist checks if directory of given resource (file or directory itself) exists.
func (r Resource) DirExist() bool {
	_, err := os.Stat(r.Directory)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

//Unlink removes resource from file system
func (r Resource) Unlink() error {
	if r.isDir() {
		return os.RemoveAll(r.Directory)
	}

	if r.isFile() {
		return os.Remove(r.Directory + r.FileName)
	}

	return ErrInvalidResource
}

//Execute unlink resource according to context.
func (r Resource) Execute(ctx context.Context, ch1 chan error) {
	select {
	case <-ctx.Done():
		return
	default:
		ch1 <- r.Unlink()
	}
}

//CheckDirStructure checks if user is in geography root folder
func CheckDirStructure(dirs []Resource) error {
	for _, dir := range dirs {
		if !dir.Exist() {
			return ErrInvalidDirectoryStructure
		}
	}

	return nil
}

//isFile checks if resource is file eg. has Directory and FileName defined
func (r Resource) isFile() bool {
	if len(r.FileName) != 0 && len(r.Directory) != 0 {
		return true
	}

	return false
}

//isDir checks if resource is directory eg. has valid Directory
func (r Resource) isDir() bool {
	if len(r.Directory) != 0 && len(r.FileName) == 0 {
		return true
	}

	return false
}
