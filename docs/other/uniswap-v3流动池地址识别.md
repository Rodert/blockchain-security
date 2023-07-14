

[toc]


案例数据

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "accessList": [],
        "blockHash": "0x9b29fac4b6460521cf964ee52c4d5c7036855fde28c797c5b0b2993876450b79",
        "blockNumber": "0x10de9a2",
        "chainId": "0x1",
        "from": "0x99a080de8963341d0029693f5a6949be2f4a5dd6",
        "gas": "0x8cef9",
        "gasPrice": "0x5a1454efb",
        "hash": "0xfab2f1a51bcca936eb6c8067990992548a14bd84d6d4f3f99365e47d2614efb9",
        "input": "0x883164560000000000000000000000008e0e57dcb1ce8d9091df38ec1bfc3b224529754a000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec70000000000000000000000000000000000000000000000000000000000000bb8fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc1e8fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc788000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004a817c800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004a817c80000000000000000000000000099a080de8963341d0029693f5a6949be2f4a5dd60000000000000000000000000000000000000000000000000000000064b0c377",
        "maxFeePerGas": "0x7affd1c72",
        "maxPriorityFeePerGas": "0x5f5e100",
        "nonce": "0x120",
        "r": "0x2b683877d1549084342faf740d2e39ff78cad27e13ac6db2ec68103722d83f99",
        "s": "0x6961d1c1b90dff13362f52cc39e01f61b4f4f1a2fa374ac00ed3de13490cda87",
        "to": "0xc36442b4a4522e871399cd717abdd847ab11fe88",
        "transactionIndex": "0x5e",
        "type": "0x2",
        "v": "0x1",
        "value": "0x0"
    }
}
```

## 解析步骤

1. 判断是添加流动池的交易

Uniswap V3 合约地址 0xc36442b4a4522e871399cd717abdd847ab11fe88

```bash
to=0xc36442b4a4522e871399cd717abdd847ab11fe88 && input=0x88316456(前缀) && len(input)=714
```

2. 解析代币对地址

Input Data 结构

```bash
Function: mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256))
#	Name	Type	Data
0	params.token0	address	0x8e0E57DCb1ce8d9091dF38ec1BfC3b224529754A
0	params.token1	address	0xdAC17F958D2ee523a2206206994597C13D831ec7
0	params.fee	uint24	3000
0	params.tickLower	int24	-278040
0	params.tickUpper	int24	-276600
0	params.amount0Desired	uint256	0
0	params.amount1Desired	uint256	20000000000
0	params.amount0Min	uint256	0
0	params.amount1Min	uint256	20000000000
0	params.recipient	address	0x99A080dE8963341D0029693f5a6949bE2F4a5dd6
0	params.deadline	uint256	1689305975
```

基于 input 解析添加到流动池的代币

token0 = input[10,74]  -- 0000000000000000000000008e0e57dcb1ce8d9091df38ec1bfc3b224529754a

token1 = input[75,138] -- 000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7





## 参考

1. Uniswap V3 交易 https://etherscan.io/tx/0xa5cb03f359691a57e27c29678204ac7623e7950684583fff48514338ad00502b
2. Uniswap V3 Swap 函数签名 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67
3. Uniswap V3 合约地址 https://etherscan.io/address/0xc36442b4a4522e871399cd717abdd847ab11fe88#code


### 添加流动池案例

1. https://etherscan.io/tx/0xfab2f1a51bcca936eb6c8067990992548a14bd84d6d4f3f99365e47d2614efb9
2. https://etherscan.io/tx/0x5850f95fd86a2e77196e0ecd71007da371fc44e902f344d5f4cb1feb046fe73e
3. https://etherscan.io/tx/0xdd257ef9c850b699881e6bb5a0e6da1559950f61cbfb52c8fb1e774b0e82be24
4. 
