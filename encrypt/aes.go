package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// 16, 24, and 32-bit strings correspond to AES-128, AES-192, and AES-256 encryption methods respectively
// PKCS7 padding mode
func Pkcs7Padding(cipherText []byte, blockSize int) []byte {
	p := blockSize - len(cipherText)%blockSize
	pt := bytes.Repeat([]byte{byte(p)}, p)
	return append(cipherText, pt...)
}

// Reverse of padding, removing padding strings
func Pkcs7UnPadding(originData []byte) ([]byte, error) {
	l := len(originData)
	if l == 0 {
		return nil, errors.New("Encrypted string error")
	} else {
		unp := int(originData[l-1])
		return originData[:(l - unp)], nil
	}
}

// Implement encryption
func AesEcr(originData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := block.BlockSize()
	originData = Pkcs7Padding(originData, size)
	blockMode := cipher.NewCBCEncrypter(block, key[:size])
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return crypted, nil
}

// Implement decryption
func AesDec(cypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:size])
	originData := make([]byte, len(cypted))
	blockMode.CryptBlocks(originData, cypted)
	originData, err = Pkcs7UnPadding(originData)
	if err != nil {
		return nil, err
	}
	return originData, err
}

// encrypted base64
func AESEnc(data, key []byte) (string, error) {
	res, err := AesEcr(data, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), err
}

// decrypt
func AESDec(encStr string, key []byte) ([]byte, error) {
	decodeString, err := base64.StdEncoding.DecodeString(encStr)
	if err != nil {
		return nil, err
	}
	return AesDec(decodeString, key)
}