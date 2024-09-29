package aesEncryptionTool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

const keyFilePath = "./aes256"

type AESCipher struct {
	key []byte
}

func NewAESCipher() (*AESCipher, error) {
	cipher := &AESCipher{}

	if _, err := os.Stat(keyFilePath); os.IsNotExist(err) {
		if err := cipher.generateKey(); err != nil {
			return nil, err
		}
	}

	key, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, err
	}

	cipher.key = key
	return cipher, nil
}

const (
	keyLength = 16
	symbols   = "!@#$%^&*()_+{}[]|:;<>,.?/~`"
)

func (a *AESCipher) generateKey() error {
	// 生成32个字符长度的随机字符串作为加密密码
	randomBytes := make([]byte, keyLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return err
	}

	// 将随机字节转换为字符串
	keyString := hex.EncodeToString(randomBytes)

	// 替换随机字符串中的字符
	for i := 0; i < len(symbols); i++ {
		index := randInt(0, len(keyString))
		keyString = keyString[:index] + string(symbols[i]) + keyString[index+1:]
	}

	// 将密钥写入文件
	return os.WriteFile(keyFilePath, []byte(keyString), 0600)
}

func randInt(min, max int) int {
	b := make([]byte, 1)
	rand.Read(b)
	return min + int(b[0])%(max-min)
}

func (a *AESCipher) Encrypt(plaintext []byte) (string, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	plaintext, err = pad(plaintext, aes.BlockSize)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (a *AESCipher) Decrypt(cipherTextStr string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherTextStr)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return unpad(ciphertext, aes.BlockSize)
}

func pad(buf []byte, blockSize int) ([]byte, error) {
	padding := blockSize - len(buf)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(buf, padText...), nil
}

func unpad(buf []byte, blockSize int) ([]byte, error) {
	length := len(buf)
	padding := int(buf[length-1])

	if padding < 1 || padding > blockSize {
		return nil, errors.New("invalid padding")
	}

	padText := buf[length-padding:]
	for _, v := range padText {
		if v != byte(padding) {
			return nil, errors.New("invalid padding")
		}
	}

	return buf[:length-padding], nil
}
