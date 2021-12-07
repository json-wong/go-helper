package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

type Aes struct {
	key string // AES cfb key
}

func NewAes(key string) (*Aes, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid aes key")
	}

	return &Aes{key: key}, nil
}

func (a *Aes) Encrypt(text string) ([]byte, error) {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(a.key))
	if err != nil {
		return []byte(""), err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return []byte(""), err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
	// convert to base64
	//return base64.URLEncoding.EncodeToString(ciphertext)
}

func (a *Aes) EncToBase64(text string) (string, error) {
	b, err := a.Encrypt(text)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

// Decrypt from base64 to decrypted string
func (a *Aes) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(a.key))
	if err != nil {
		return []byte(""), err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the cipherText.
	if len(cipherText) < aes.BlockSize {
		return []byte(""), errors.New("cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func (a *Aes) DecFromBase64(cryptoText string) ([]byte, error) {
	b, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return []byte(""), err
	}

	return a.Decrypt(b)
}

func (a *Aes) DecFromBase64ToString(cryptoText string) (string, error) {
	b, err := a.DecFromBase64(cryptoText)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", b), nil
}
