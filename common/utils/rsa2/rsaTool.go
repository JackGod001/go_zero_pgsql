package rsa2

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RSAService 结构体用于封装RSA操作
type RSAService struct {
	//私钥可选
	privateKey *rsa.PrivateKey
	//公钥必须
	publicKey *rsa.PublicKey
}

// NewRSAService 创建一个新的RSAService实例
func NewRSAService(privateKeyPEM []byte, publicKeyPEM []byte) (*RSAService, error) {
	rsaService := &RSAService{}

	if privateKeyPEM != nil {
		// 解析私钥
		block, _ := pem.Decode(privateKeyPEM)
		if block == nil {
			return nil, errors.New("failed to parse PEM block containing the private key")
		}

		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaService.privateKey = privateKey
	}

	// 解析公钥
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	var ok bool
	rsaService.publicKey, ok = pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}

	return rsaService, nil
}

// Sign 使用私钥对数据进行签名 rsa2采用 sha-256散列
func (r *RSAService) Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)
	return rsa.SignPKCS1v15(rand.Reader, r.privateKey, crypto.SHA256, hashed[:])
}

// VerifySign 使用公钥验证签名
func (r *RSAService) VerifySign(data []byte, signature []byte) error {
	hashed := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(r.publicKey, crypto.SHA256, hashed[:], signature)
}

// GenerateRSAKeys 生成RSA密钥对
func GenerateRSAKeys(bits int) ([]byte, []byte, error) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	// 编码私钥为PEM格式
	privateKeyDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyDER,
	}
	privateKeyPEM := pem.EncodeToMemory(privateKeyBlock)

	// 从私钥中提取公钥
	publicKey := &privateKey.PublicKey

	// 编码公钥为PEM格式
	publicKeyDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	publicKeyBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyDER,
	}
	publicKeyPEM := pem.EncodeToMemory(publicKeyBlock)

	return privateKeyPEM, publicKeyPEM, nil
}
