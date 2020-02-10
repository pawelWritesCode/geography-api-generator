package templates

import "os"

type Existence interface {
	Exist() bool
}

type Resource struct {
	Directory string
	FileName  string
}

func (r Resource) isFile() bool {
	if len(r.FileName) != 0 && len(r.Directory) != 0 {
		return true
	}

	return false
}

func (r Resource) isDir() bool {
	if len(r.Directory) != 0 && len(r.FileName) == 0 {
		return true
	}

	return false
}

func (r Resource) Exist() bool {
	path := ""

	if r.isDir() {
		path = r.Directory
	}

	if r.isFile() {
		path = r.Directory + r.FileName
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func (r Resource) DirExist() bool {
	_, err := os.Stat(r.Directory)
	if os.IsNotExist(err) {
		return false
	}

	return true
}
