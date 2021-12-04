package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
)

// Md5 - Calculate the md5 hash of a string
func Md5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Md5File - Calculates the md5 hash of a given file
func Md5File(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// Sha1 - Calculate the sha1 hash of a string
func Sha1(s string) string {
	digest := sha1.Sum([]byte(s))
	return hex.EncodeToString(digest[:])
}

// Sha1File - Calculates the md5 hash of a given file
func Sha1File(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	digest := sha1.Sum(data)
	return hex.EncodeToString(digest[:])
}