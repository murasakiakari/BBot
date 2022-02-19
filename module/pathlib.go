package module

import (
	"os"
	"path/filepath"
)

var CurrentWorkingDirectory = Path(os.Args[0]).Dir()

type Path string

func (path Path) Abs() Path {
	pathString, _ := filepath.Abs(string(path))
	return Path(pathString)
}

func (path Path) Base() string {
	return filepath.Base(string(path))
}

func (path Path) Dir() Path {
	return Path(filepath.Dir(string(path)))
}

func (path Path) Ext() string {
	return filepath.Ext(string(path))
}

func (path Path) IsExist() bool {
	_, err := os.Stat(string(path))
	return !os.IsNotExist(err)
}

func (path Path) Join(element ...string) Path {
	tempPath := make([]string, len(element) + 1)
	tempPath[0] = string(path)
    copy(tempPath[1:], element)
	return Path(filepath.Join(tempPath...))
}

func (path Path) ReadFile() ([]byte, error) {
	return os.ReadFile(string(path))
}
