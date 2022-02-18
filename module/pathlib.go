package module

import (
	"os"
	"path/filepath"
)

var CurrentWorkingDirectory Path = NewPath(os.Args[0]).Dir()

type Path struct {
	path string	
}

func (p Path) Abs() Path {
	p.path, _ = filepath.Abs(p.path)
	return p
}

func (p Path) Base() string {
	return filepath.Base(p.path)
}

func (p Path) Dir() Path {
	p.path = filepath.Dir(p.path)
	return p
}

func (p Path) Ext() string {
	return filepath.Ext(p.path)
}

func (p Path) IsExist() bool {
	_, err := os.Stat(p.path)
	return !os.IsNotExist(err)
}

func (p Path) Join(element ...string) Path {
	tempPath := make([]string, len(element) + 1)
	tempPath[0] = p.path
    copy(tempPath[1:], element)
	p.path = filepath.Join(tempPath...)
	return p
}

func (p Path) ReadFile() ([]byte, error) {
	return os.ReadFile(p.path)
}

func (p Path) String() string {
	return p.path
}

func NewPath(path string) Path {
	return Path{path}
}
