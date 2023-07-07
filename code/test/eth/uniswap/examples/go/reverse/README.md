[toc]


1. 安装 abigen，通过 abi 生成 GoLang 代码

2. 例子

https://etherscan.io/address/0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11#code

```bash
../bin/abigen --abi=pair.abi.json --type=pair_contract --pkg=uniswap --out=pair_uniswap.go
```

