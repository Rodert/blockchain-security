package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// str3 := ""
	// str4 := ""
	// fmt.Println(utils.LCS(str3, str4))
	// eth.InputAnalysis()
	// tron.Trc20Balance()
	// eth.GetCodeByAlchemy()
	// spider.AnalysisList()
	// spider.AnalysisProject()
	// Transform()
	// Tr2()
	// Tr3()
	Tr4()
}

func Tr4() {
	contractCodeHex := "6060604052341561000f57600080fd5b6040516103d93803806103d9833981016040528080518201919050505b8060016000509080519060200190610030929190610034565b505b5061006c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a857805160ff19168380011785556100d2565b828001600101855582156100d2579182015b828111156100d15782518255916020019190600101906100b6565b5b5090506100df91906100e3565b5090565b61010591905b808211156101015760008160009055506001016100e9565b5090565b90565b6103f18061011d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80632a1afcd014602d575b600080fd5b60336047565b6040518082815260200191505060405180910390f35b603b60ba565b6040518082815260200191505060405180910390f35b60008183019050805080905060c3565b919050565b60006020828403121561009e578081fd5b813590509291505056fea26469706673582212207a7a315d0391d4307f8b6a32764f20d2b2c61f8f3506db9b5bb5b5da5f28e92064736f6c634300060b0033"
	contractCode, err := hexutil.Decode(contractCodeHex)
	if err != nil {
		log.Fatal(err)
	}
	contractCodeHash := crypto.Keccak256Hash(contractCode)
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ee952d078e364b27a23da799090aecae")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0x123...")
	abiBytes, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	abiStr := string(abiBytes)
	_ = contractCodeHash
	_ = abiStr
	fmt.Println(contractCodeHash)
	fmt.Println(abiStr)
}

func Tr3() {
	// 输入以太坊智能合约的十六进制代码
	contractCodeHex := "6060604052341561000f57600080fd5b6040516103d93803806103d9833981016040528080518201919050505b8060016000509080519060200190610030929190610034565b505b5061006c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a857805160ff19168380011785556100d2565b828001600101855582156100d2579182015b828111156100d15782518255916020019190600101906100b6565b5b5090506100df91906100e3565b5090565b61010591905b808211156101015760008160009055506001016100e9565b5090565b90565b6103f18061011d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80632a1afcd014602d575b600080fd5b60336047565b6040518082815260200191505060405180910390f35b603b60ba565b6040518082815260200191505060405180910390f35b60008183019050805080905060c3565b919050565b60006020828403121561009e578081fd5b813590509291505056fea26469706673582212207a7a315d0391d4307f8b6a32764f20d2b2c61f8f3506db9b5bb5b5da5f28e92064736f6c634300060b0033"

	// 将十六进制字符串转换为字节数组
	contractCode, err := hex.DecodeString(contractCodeHex)
	if err != nil {
		panic(err)
	}

	// 计算合约代码的SHA3哈希
	contractCodeHash := crypto.Keccak256(contractCode)

	// 输出合约代码的哈希值
	fmt.Println("Contract code hash:", hexutil.Encode(contractCodeHash))

	// 计算合约的ABI
	// abi, err := crypto.ABIContractABIGen(
	// 	map[string]interface{}{},
	// 	nil,
	// 	contractCodeHash,
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// 输出合约ABI
	// fmt.Println("Contract ABI:", string(abi))
}

func Tr2() {
	// 将智能合约代码转换为字节数组
	contractByteCode := "608060405234801561001057600080fd5b5061012f806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806306661abd14602d575b600080fd5b60536004356001565b60408051918252519081900360200190f35b60536004356001565b6000600782029050604b565b91905056fea2646970667358221220d9db9cbbd7a05a130331fd7f44ebd1a54a6eb80ec682f07888b9e95dbd508bbb64736f6c634300060b0033"
	contractByteCodeBytes, _ := hex.DecodeString(contractByteCode)

	// 使用Solidity编译器将Solidity代码编译为ABI
	// contractAbi, _ := abi.JSON(strings.NewReader(abiDefinition))
	contractAbi, _ := abi.JSON(bytes.NewReader(contractByteCodeBytes))
	_ = contractAbi
	fmt.Println(contractAbi.Constructor.Name)
	// 将ABI转换为golang结构体
	type Contract struct {
		AddPerson func(string, uint8, uint8)
		GetPeople func() ([][3]interface{}, error)
	}

	contract := Contract{}
	// contractAbi.Unpack(&contract, "AddPerson", contractAbi)
	fmt.Println(contract)
}

func Transform() {
	// 将智能合约的十六进制代码转换为字节数组
	contractCodeHex := "0x606060405260008060146101000a81548160ff0219169083151502179055506000600355600060045534156200003457600080fd5b60405162002d7c38038062002d7c83398101604052808051906020019091908051820191906020018051820191906020018051906020019091905050336000806101000a81548173"
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
