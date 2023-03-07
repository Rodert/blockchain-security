package eth

import (
	"context"
	"fmt"
	"test/eth/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Erc20Balance() {
	rawurl := "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		fmt.Println(err)
		return
	}

	addr := "0xda98f2EfF1D348bD58B6E5d636d8f123DF3AF535"
	token20Address := "0xdac17f958d2ee523a2206206994597c13d831ec7"

	// 获取原生币余额
	ethBalance, err := client.BalanceAt(
		context.Background(),
		common.HexToAddress(addr),
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("eth balance: %v\n", ethBalance)

	// 获取代币余额
	token20Balance, err := BalanceToken(
		client,
		token20Address,
		addr,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("token20 balance: %v\n", token20Balance)
}

func BalanceToken(backend bind.ContractBackend, tokenAddress, address string) (string, error) {
	tokenAddr := common.HexToAddress(tokenAddress)
	instance, err := erc20.NewContractName(tokenAddr, backend)
	if err != nil {
		return "", err
	}

	addr := common.HexToAddress(address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, addr)
	if err != nil {
		return "", err
	}

	return bal.String(), nil
}
