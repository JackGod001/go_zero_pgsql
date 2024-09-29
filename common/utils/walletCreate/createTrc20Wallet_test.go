package walletCreate

import (
	"fmt"
	"testing"
)

// 测试创建trc20地址
func TestCreateTRC20Address(t *testing.T) {
	t.Log("Test Create TRC20 Address")
	wallet, err := walletCreate.GenerateTRCWallet()
	if err != nil {
		t.Error("Test Create TRC20 Address Error:", err)
		return
	}
	fmt.Println("TRC20 Address:", wallet.Address)
	fmt.Println("TRC20 Private Key:", wallet.PrivateKey)
	fmt.Println("TRC20 MNEMONIC:", wallet.Mnemonic)
	fmt.Println("TRC20 Public Key:", wallet.PublicKey)
}
