package accounttest_test

import (
	"fmt"
	"test/eth/account"
	"testing"
)

func TestGetPrivateKeyByKeystore(t *testing.T) {
	file := "/Users/shiyuwang/Library/Ethereum/keystore/UTC--2023-03-30T02-49-39.560806000Z--d4738610478cb25391802d1ae2008f91a3d5ca0c"
	privKey, address := account.GetPrivateKeyByKeystore(file, "123456")
	fmt.Println("##########")
	fmt.Println(privKey)
	fmt.Println(address)
}
