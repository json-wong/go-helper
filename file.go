package helper

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// FileGetContents - Reads entire file into a string
func FileGetContents(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// FilePutContents - Write data to a file
func FilePutContents(filename string, data []byte) error {
	if dir := filepath.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// FileExists - Checks whether a file or directory exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Filemtime - Gets file modification time
func Filemtime(file string) time.Time {
	fi, err := os.Stat(file)
	if err != nil {
		return time.Time{}
	}
	return fi.ModTime()
}

// IsDir - Tells whether the filename is a directory
func IsDir(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.IsDir()
}

// IsFile Tells whether the filename is a regular file
func IsFile(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.Mode().IsRegular()
}