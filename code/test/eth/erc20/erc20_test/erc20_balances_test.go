package erc20test_test

import (
	"context"
	"fmt"
	"log"
	"test/eth/erc20"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

/*
	已验证

批量获取地址 ERC20 余额
*/
func TestGetX(t *testing.T) {
	rets := []*string{}
	// 以太坊 eth_call 的参数结构体
	type CallArg struct {
		// common.Address 是以太坊依赖包的地址类型，其原型是 [20]byte 数组
		From     common.Address `json:"from"`
		To       common.Address `json:"to"`
		Gas      string         `json:"gas"`
		GasPrice string         `json:"gas_price"`
		Value    string         `json:"value"`
		Data     string         `json:"data"` // 这个就是 data
		Nonce    string         `json:"nonce"`
	}

	nodeUrl := "https://mainnet.infura.io/v3/ee952d078e364b27a23da799090aecae"
	client, _ := rpc.DialContext(context.Background(), nodeUrl)

	batchElemList := []rpc.BatchElem{}
	name := "eth_call"
	methodId := "0x70a08231" // 这个就是 balanceOf 的 methodId

	for i := 0; i < 100000; i++ {
		res := ""
		arg := &CallArg{}
		// 下面是针对访问 balanceOf 时的必须参数，查询余额是不需要油费的，但是发现一些版本的节点又需要指定这个参数，所以下面还是指定一个
		arg.Gas = hexutil.EncodeUint64(300000)
		arg.To = common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
		//  data 参数的组合格式见 “交易参数的说明” 小节中的详解
		arg.Data = methodId + "000000000000000000000000" + "ab5801a7d398351b8be11c439e05c5b3259aec9b"

		batchElemList = append(batchElemList, rpc.BatchElem{
			Method: name,
			Args:   []interface{}{arg, "latest"},
			Result: &res,
		})
		rets = append(rets, &res)
	}

	fmt.Println(batchElemList)
	if err := client.BatchCall(batchElemList); err != nil {
		fmt.Println(err)
	}

	fmt.Println(rets)
	fmt.Println(len(rets))
}

func TestGetBalances(t *testing.T) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ee952d078e364b27a23da799090aecae")
	if err != nil {
		log.Fatalln("dialing mainnet", err)
	}

	accounts := []common.Address{
		common.HexToAddress("0x4e83362442b8d1bec281594cea3050c8eb01311c"),
		common.HexToAddress("0x03c6cf4fafb0caf2c2d2917ef6ab8c3fb1bb5ffb"),
	}

	balances, err := erc20.BatchGetBalances(client, accounts)
	if err != nil {
		log.Fatalln("getting balances", err)
	}

	for a, bal := range balances {
		fmt.Printf("Address: %v \t Balance: %v \n", a.Hex(), bal)
	}
}

func TestGetERC20Balances(t *testing.T) {

	// ERC20BalanceRpcReq 是查询 ERC20 代币的参数集合结构体
	// type ERC20BalanceRpcReq struct {
	// 	ContractAddress string // 合约的以太坊地址
	// 	UserAddress     string // 用户的以太坊地址
	// 	ContractDecimal int    // 合约所对应代币的数位
	// }
	// // 以太坊 eth_call 的参数结构体
	// type CallArg struct {
	// 	// common.Address 是以太坊依赖包的地址类型，其原型是 [20]byte 数组
	// 	From     common.Address `json:"from"`
	// 	To       common.Address `json:"to"`
	// 	Gas      string         `json:"gas"`
	// 	GasPrice string         `json:"gas_price"`
	// 	Value    string         `json:"value"`
	// 	Data     string         `json:"data"` // 这个就是 data
	// 	Nonce    string         `json:"nonce"`
	// }

	// nodeUrl := "https://mainnet.infura.io/v3/ee952d078e364b27a23da799090aecae"
	// client, err := ethclient.Dial(nodeUrl)
	// if err != nil {
	// 	log.Fatalln("dialing mainnet", err)
	// }

	// address := "0xc58AD8Ff428c354bb849d1dCf1EDCcAC3F102C8E"   // 钱包地址
	// contract1 := "0x78021ABD9b06f0456cB9DB95a846C302c34f8b8D" // 合约地址1
	// contract2 := "0xB8c77482e45F1F44dE1745F52C74426C631bDD52" // 合约地址2

	// paramArr := []ERC20BalanceRpcReq{}
	// item := ERC20BalanceRpcReq{}
	// item.ContractAddress = contract1
	// item.UserAddress = address
	// item.ContractDecimal = 18
	// paramArr = append(paramArr, item)
	// item.ContractAddress = contract2
	// paramArr = append(paramArr, item)

	// name := "eth_call"
	// methodId := "0x70a08231" // 这个就是 balanceOf 的 methodId
	// // 结果数组存储的是每个请求的结果指针，也就是引用
	// rets := []*string{}
	// // 获取参数数组的长度，方便在循环中逐个实例化 BatchElem
	// size := len(paramArr)
	// reqs := []rpc.BatchElem{}

	// for i := 0; i < size; i++ {
	// 	ret := ""
	// 	arg := &CallArg{}
	// 	userAddress := paramArr[i].UserAddress
	// 	// 下面是针对访问 balanceOf 时的必须参数，查询余额是不需要油费的，但是发现一些版本的节点又需要指定这个参数，所以下面还是指定一个
	// 	arg.Gas = hexutil.EncodeUint64(300000)
	// 	arg.To = common.HexToAddress(paramArr[i].ContractAddress)
	// 	//  data 参数的组合格式见 “交易参数的说明” 小节中的详解
	// 	arg.Data = methodId + "000000000000000000000000" + userAddress[2:]
	// 	// 实例化每个 BatchElem
	// 	req := rpc.BatchElem{
	// 		Method: name,
	// 		Args:   []interface{}{arg, "latest"},
	// 		// &ret 传入单个请求的结果引用，这样是保证它在函数内部被修改值后，回到函数外来，值仍有效
	// 		Result: &ret,
	// 	}
	// 	reqs = append(reqs, req)  // 将每个 BatchElem 添加到 BatchElem 数组
	// 	rets = append(rets, &ret) // 每个请求的结果引用添加到结果数组中
	// }

	// err = client.GetRpc().BatchCall(reqs) // 传入 BatchElem 数组，发起批量请求
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
