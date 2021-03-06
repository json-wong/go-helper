package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

var (
	// ErrPrivateKey indicates the invalid private key.
	ErrPrivateKey = errors.New("private key error")
	// ErrPublicKey indicates the invalid public key.
	ErrPublicKey = errors.New("failed to parse PEM block containing the public key")
	// ErrNotRsaKey indicates the invalid RSA key.
	ErrNotRsaKey = errors.New("key type is not RSA")
)

type (
	// RsaDecrypter represents a RSA decrypter.
	RsaDecrypter interface {
		Decrypt(input []byte) ([]byte, error)
		DecryptFromBase64(input string) ([]byte, error)
	}

	// RsaEncrypter represents a RSA encrypter.
	RsaEncrypter interface {
		Encrypt(input []byte) ([]byte, error)
		EncryptToString(input []byte) (string, error)
	}

	rsaBase struct {
		bytesLimit int
	}

	rsaDecrypter struct {
		rsaBase
		privateKey *rsa.PrivateKey
	}

	rsaEncrypter struct {
		rsaBase
		publicKey *rsa.PublicKey
	}
)

// NewRsaDecrypter returns a RsaDecrypter with the given file.
func NewRsaDecrypter(key string) (RsaDecrypter, error) {
	content := []byte(key)
	block, _ := pem.Decode(content)
	if block == nil {
		return nil, ErrPrivateKey
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &rsaDecrypter{
		rsaBase: rsaBase{
			bytesLimit: privateKey.N.BitLen() >> 3,
		},
		privateKey: privateKey,
	}, nil
}

func (r *rsaDecrypter) Decrypt(input []byte) ([]byte, error) {
	return r.crypt(input, func(block []byte) ([]byte, error) {
		return rsaDecryptBlock(r.privateKey, block)
	})
}

func (r *rsaDecrypter) DecryptFromBase64(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, nil
	}

	base64Decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}

	return r.Decrypt(base64Decoded)
}

// NewRsaEncrypter returns a RsaEncrypter with the given key.
func NewRsaEncrypter(key string) (RsaEncrypter, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, ErrPublicKey
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pubKey := pub.(type) {
	case *rsa.PublicKey:
		return &rsaEncrypter{
			rsaBase: rsaBase{
				// https://www.ietf.org/rfc/rfc2313.txt
				// The length of the data D shall not be more than k-11 octets, which is
				// positive since the length k of the modulus is at least 12 octets.
				bytesLimit: (pubKey.N.BitLen() >> 3) - 11,
			},
			publicKey: pubKey,
		}, nil
	default:
		return nil, ErrNotRsaKey
	}
}

func (r *rsaEncrypter) Encrypt(input []byte) ([]byte, error) {
	return r.crypt(input, func(block []byte) ([]byte, error) {
		return rsaEncryptBlock(r.publicKey, block)
	})
}

// EncryptToString AES encode
func (r *rsaEncrypter) EncryptToString(input []byte) (string, error) {
	enc, err := r.Encrypt(input)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(enc), nil
}

func (r *rsaBase) crypt(input []byte, cryptFn func([]byte) ([]byte, error)) ([]byte, error) {
	var result []byte
	inputLen := len(input)

	for i := 0; i*r.bytesLimit < inputLen; i++ {
		start := r.bytesLimit * i
		var stop int
		if r.bytesLimit*(i+1) > inputLen {
			stop = inputLen
		} else {
			stop = r.bytesLimit * (i + 1)
		}
		bs, err := cryptFn(input[start:stop])
		if err != nil {
			return nil, err
		}

		result = append(result, bs...)
	}

	return result, nil
}

func rsaDecryptBlock(privateKey *rsa.PrivateKey, block []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, block)
}

func rsaEncryptBlock(publicKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg)
}

// GenerateRSAKey ??????RSA????????????????????????????????????
// bits ????????????
func GenerateRSAKey(bits int) {
	//GenerateKey?????????????????????????????????random????????????????????????????????????RSA??????
	//Reader?????????????????????????????????????????????????????????
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//????????????
	//??????x509??????????????????ras??????????????????ASN.1 ??? DER???????????????
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//??????pem?????????x509???????????????????????????
	//????????????????????????
	privateFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//????????????pem.Block???????????????
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//????????????????????????
	pem.Encode(privateFile, &privateBlock)

	//????????????
	//?????????????????????
	publicKey := privateKey.PublicKey
	//X509???????????????
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem????????????
	//?????????????????????????????????
	publicFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//????????????pem.Block???????????????
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//???????????????
	pem.Encode(publicFile, &publicBlock)
}
