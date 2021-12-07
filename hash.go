package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Md5 - Calculate the md5 hash of a string
func Md5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

// Md5File - Calculates the md5 hash of a given file
func Md5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return "", err
	}

	var size int64 = 1048576 // 1M
	hash := md5.New()

	if fi.Size() < size {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return "", err
		}
		hash.Write(data)
	} else {
		b := make([]byte, size)
		for {
			n, err := f.Read(b)
			if err != nil {
				break
			}

			hash.Write(b[:n])
		}
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Sha1 - Calculate the sha1 hash of a string
func Sha1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha1File - Calculates the md5 hash of a given file
func Sha1File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func FileSha256(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", errors.New("Open file failure: " + file)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	h := sha256.New()
	if _, err = io.Copy(h, f); err != nil {
		return "", err
	}
	selfSum := fmt.Sprintf("%x", h.Sum(nil))

	return selfSum, nil
}
