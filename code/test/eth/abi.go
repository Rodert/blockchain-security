package eth

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func Transform() {
	// 将智能合约的十六进制代码转换为字节数组
	contractCodeHex := "0x..."
	contractCodeBytes, err := hex.DecodeString(contractCodeHex)
	if err != nil {
		panic(err)
	}

	// 通过字节数组创建abi对象
	contractAbi, err := abi.JSON(bytes.NewReader(contractCodeBytes))
	if err != nil {
		panic(err)
	}

	// 打印abi对象
	fmt.Println(contractAbi)
}
