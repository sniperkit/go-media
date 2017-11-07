package ioe

import (
	"io"
	"os"
	"path/filepath"
)

type FS interface {
	Open(name string) (File, error)
}

type File interface {
	io.ReadWriteSeeker
	io.Closer
	Stat() (os.FileInfo, error)
	Readdir(int) ([]os.FileInfo, error)
}

type SFS struct {
	Root string
}

func (fs *SFS) Chdir(dir string) error {
	fs.Root = filepath.Join(fs.Root, dir)
	return nil
}

func (fs *SFS) Open(name string) (File, error) {
	name = filepath.Join(fs.Root, name)
	f, err := os.Open(name)
	return File(f), err
}

func (fs *SFS) Create(name string) (File, error) {
	name = filepath.Join(fs.Root, name)
	f, err := os.Create(name)
	return File(f), err
}

func (fs *SFS) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	name = filepath.Join(fs.Root, name)
	f, err := os.OpenFile(name, flag, perm)
	return File(f), err
}

func (fs *SFS) Remove(name string) error {
	name = filepath.Join(fs.Root, name)
	return os.Remove(name)
}

func (fs *SFS) Stat(name string) (os.FileInfo, error) {
	name = filepath.Join(fs.Root, name)
	return os.Stat(name)
}

func (fs *SFS) Rename(oldpath, newpath string) error {
	oldpath = filepath.Join(fs.Root, oldpath)
	newpath = filepath.Join(fs.Root, newpath)
	return os.Rename(oldpath, newpath)
}