package tron

import (
	"context"
	"encoding/hex"
	"fmt"
	"test/eth/erc20"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Trc20Balance() {
	rawurl := "https://api.trongrid.io/jsonrpc"

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		fmt.Println(err)
		return
	}

	addr, _ := Base58ToHex("TVqN1kf2zEASkLV8ySMoLH3gF3uAnhJjPT")
	token20Address, _ := Base58ToHex("TJcXHanekjutcEU9vG6PWNotNgMDpu6zVu")

	// 获取原生币余额
	trxBalance, err := client.BalanceAt(
		context.Background(),
		common.HexToAddress(addr),
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("trx balance: %v\n", trxBalance)

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

func Base58ToHex(input string) (string, error) {
	decoded, _, err := base58.CheckDecode(input)
	if err != nil {
		return "", err
	}

	return "0x" + hex.EncodeToString(decoded), nil
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
