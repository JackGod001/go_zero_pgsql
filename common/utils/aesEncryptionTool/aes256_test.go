package aesEncryptionTool

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	cipher, err := NewAESCipher()
	if err != nil {
		//fmt.Println("Error initializing AES cipher:", err)
		return
	}

	//生成key文件
	//err = cipher.generateKey()
	//if err != nil {
	//	fmt.Println("Error generating key:", err)
	//	return
	//}
	//const textfile = "/Users/baishaojie/.ssh/private.pem"
	//读取 textfile 内容
	//data, err := ioutil.ReadFile(textfile)
	//if err != nil {
	//	log.Fatal(err)
	//}
	data := "加密内容 哈哈哈"

	// Example usage
	plaintext := data
	encrypted, err := cipher.Encrypt([]byte(plaintext))
	if err != nil {
		//fmt.Println("Error encrypting:", err)
		return
	}

	fmt.Println("Encrypted:", encrypted)

	decrypted, err := cipher.Decrypt(encrypted)
	if err != nil {
		//fmt.Println("Error decrypting:", err)
		return
	}

	fmt.Println("Decrypted:", string(decrypted))
}
