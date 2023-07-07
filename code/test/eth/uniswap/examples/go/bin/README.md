abigen 是合约 abi 转换为 golang 的工具，目前在 mac 下使用。

命令：

https://etherscan.io/address/0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11#code

```bash
../../bin/abigen --abi=pair.abi.json --type=pair_contract --pkg=uniswap --out=pair_uniswap.go
```