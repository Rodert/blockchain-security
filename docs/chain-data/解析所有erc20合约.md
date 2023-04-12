# ETH

解析所有 Erc20 合约

## 基于所有区块数据计算


方案一：

- `to = null || to = ''`，contractAddress 是一个地址，input 包含以下字符串。

合约代码十六进制包含以下字符串

```bash
18160ddd
70a08231
a9059cbb
dd62ed3e
095ea7b3
23b872dd
```

对应的函数

```go
    /* ERC20 */
    // data = []byte("totalSupply()")
    // data = []byte("balanceOf(address)")
    // data = []byte("transfer(address,uint256)")
    // data = []byte("allowance(address,address)")
    // data = []byte("approve(address,uint256)")
    // data = []byte("transferFrom(address,address,uint256)")
    hash := crypto.Keccak256Hash(data)
    fmt.Println(hash.Hex())
```

案例：https://etherscan.io/tx/0x436fc7d21ed4a0a634f41b50ccb42fca12be7322de5bf9a20c97bdccbb5b2a04

## 