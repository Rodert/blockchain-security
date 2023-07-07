package gotest_test

import (
	"fmt"
	uniswap_go "test/eth/uniswap/examples/go"
	"testing"
)

// 代币地址
const (
	DAIAddress    = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
	WETHAddress   = "0x6b175474e89094c44da98b954eedeac495271d0f"
	PEPEAddress   = "0x6982508145454ce325ddbe47a25d4ec3d2311933"
	USDCAddress   = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	SHIB20Address = "0xdfef6416ea3e6ce587ed42aa7cb2e586362cbbfa"
	UNIAddress    = "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
	XYOAddress    = "0x55296f69f40ea6d20e478533c15a6b08b654e758"
)

// 代币流动池
const (
	DAIWETH = "0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11"
)

const (
	InfuraNode1 = "https://mainnet.infura.io/v3/d6624913c27d4f5a9540f071767b4d49"
)

// 获取 Uniswap 价格比
func TestUniswap0(t *testing.T) {
	/* 案例 DAI/WETH 价格比 */
	pm := uniswap_go.GetExchangeAmount(DAIAddress, WETHAddress, DAIWETH, InfuraNode1)
	fmt.Printf("%+v\n", pm)
}

/*
	参数：代币地址 - 流动池地址 - 节点地址
*/
func TestUniswap1(t *testing.T) {
	pm := uniswap_go.GetExchangeAmount(DAIAddress, UNIAddress, "", InfuraNode1)
	fmt.Printf("%+v\n", pm)
}
