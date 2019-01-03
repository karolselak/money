package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Open opens a file
func Open(f string) (*os.File, error) {
	file, err := os.Open(f)
	return file, err
}

// Close cloeses a file
func Close(f io.Closer) error {
	err := f.Close()
	return err
}

// Read reads a file to byte and returns it
func Read(f *os.File) ([]byte, error) {
	bytes, err := ioutil.ReadAll(f)
	return bytes, err
}

// Write writes given bytes to file in fpath
func Write(bytes []byte, fpath string) error {
	err := ioutil.WriteFile(fpath, bytes, 0644)
	return err
}

// Copy copies file in src to dst
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)

	if err != nil {
		return 0, err
	}

	defer source.Close()
	destination, err := os.Create(dst)

	if err != nil {
		return 0, err
	}

	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err

}
