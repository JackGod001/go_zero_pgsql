package rsa2

import (
	"testing"
)

func TestGenerateRSAKeys(t *testing.T) {
	privateKeyPEM, publicKeyPEM, err := GenerateRSAKeys(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA keys: %v", err)
	}

	if privateKeyPEM == nil || len(privateKeyPEM) == 0 {
		t.Error("Generated private key PEM is nil or empty")
	}

	if publicKeyPEM == nil || len(publicKeyPEM) == 0 {
		t.Error("Generated public key PEM is nil or empty")
	}

	// 创建一个RSA服务实例
	rsaService, err := NewRSAService(privateKeyPEM, publicKeyPEM)
	if err != nil {
		t.Fatalf("Failed to create RSAService with generated keys: %v", err)
	}

	// 创建要签名的数据
	testData := []byte("Hello, RSA!")
	// 测试签名
	signature, err := rsaService.Sign(testData)
	if err != nil {
		t.Fatalf("Failed to sign data: %v", err)
	}
	// 测试验证签名
	err = rsaService.VerifySign(testData, signature)
	if err != nil {
		t.Errorf("Failed to verify signature: %v", err)
	}
}
