

[toc]


## 案例数据

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

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "blockHash": "0x9b29fac4b6460521cf964ee52c4d5c7036855fde28c797c5b0b2993876450b79",
        "blockNumber": "0x10de9a2",
        "contractAddress": null,
        "cumulativeGasUsed": "0x94939a",
        "effectiveGasPrice": "0x5a1454efb",
        "from": "0x99a080de8963341d0029693f5a6949be2f4a5dd6",
        "gasUsed": "0x74c35",
        "logs": [
            {
                "address": "0xdac17f958d2ee523a2206206994597c13d831ec7",
                "blockHash": "0x9b29fac4b6460521cf964ee52c4d5c7036855fde28c797c5b0b2993876450b79",
                "blockNumber": "0x10de9a2",
                "data": "0x00000000000000000000000000000000000000000000000000000004a817c800",
                "logIndex": "0xea",
                "removed": false,
                "topics": [
                    "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                    "0x00000000000000000000000099a080de8963341d0029693f5a6949be2f4a5dd6",
                    "0x000000000000000000000000371a0985d104c706cafa26ffb990f55ba55bf6c5"
                ],
                "transactionHash": "0xfab2f1a51bcca936eb6c8067990992548a14bd84d6d4f3f99365e47d2614efb9",
                "transactionIndex": "0x5e"
            },
            {
                "address": "0x371a0985d104c706cafa26ffb990f55ba55bf6c5",
                "blockHash": "0x9b29fac4b6460521cf964ee52c4d5c7036855fde28c797c5b0b2993876450b79",
                "blockNumber": "0x10de9a2",
                "data": "0x000000000000000000000000c36442b4a4522e871399cd717abdd847ab11fe88000000000000000000000000000000000000000000000000040d1438cd5f947e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004a817c800",
                "logIndex": "0xeb",
                "removed": false,
                "topics": [
                    "0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde",
                    "0x000000000000000000000000c36442b4a4522e871399cd717abdd847ab11fe88",
                    "0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc1e8",
                    "0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc788"
                ],
                "transactionHash": "0xfab2f1a51bcca936eb6c8067990992548a14bd84d6d4f3f99365e47d2614efb9",
                "transactionIndex": "0x5e"
            },
            {
                "address": "0xc36442b4a4522e871399cd717abdd847ab11fe88",
                "blockHash": "0x9b29fac4b6460521cf964ee52c4d5c7036855fde28c797c5b0b2993876450b79",
                "blockNumber": "0x10de9a2",
                "data": "0x",
                "logIndex": "0xec",
                "removed": false,
                "topics": [
                    "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                    "0x0000000000000000000000000000000000000000000000000000000000000000",
                    "0x00000000000000000000000099a080de8963341d0029693f5a6949be2f4a5dd6",
                    "0x00000000000000000000000000000000000000000000000000000000000831b4"
                ],
                "transactionHash": "0xfab2f1a51bcca936eb6c8067990992548a14bd84d6d4f3f99365e47d2614efb9",
                "transactionIndex": "0x5e"
            },
            {
                "address": "0xc36442b4a4522e871399cd717abdd847ab11fe88",
                "blockHash": "0x9b29fac4b6460521cf964ee52c4d5c7036855fde28c797c5b0b2993876450b79",
                "blockNumber": "0x10de9a2",
                "data": "0x000000000000000000000000000000000000000000000000040d1438cd5f947e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004a817c800",
                "logIndex": "0xed",
                "removed": false,
                "topics": [
                    "0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f",
                    "0x00000000000000000000000000000000000000000000000000000000000831b4"
                ],
                "transactionHash": "0xfab2f1a51bcca936eb6c8067990992548a14bd84d6d4f3f99365e47d2614efb9",
                "transactionIndex": "0x5e"
            }
        ],
        "logsBloom": "0x00800000000000000000100000000020800000000000000000000000000000400000000000000000000000000000010000010000020000000400000008000000000000000800800010000018000000000000000000000000000088000000000000000000020000000000000000000800000000000800000000000010000000000000000008000000000000200000000000000000000000000000000000100000000000040000000000000080000000000000008000200000000000000000004040000002000000400000080000000000000020100000000000000000000020000000000000000000000000000000000000000004000002000000000000000800",
        "status": "0x1",
        "to": "0xc36442b4a4522e871399cd717abdd847ab11fe88",
        "transactionHash": "0xfab2f1a51bcca936eb6c8067990992548a14bd84d6d4f3f99365e47d2614efb9",
        "transactionIndex": "0x5e",
        "type": "0x2"
    }
}
```

## 解析步骤

1. 判断是添加流动池的交易

Uniswap V3 合约地址 0xc36442b4a4522e871399cd717abdd847ab11fe88

```bash
if to=0xc36442b4a4522e871399cd717abdd847ab11fe88 && input=0x88316456(前缀) && len(input)=714
    return true
```

2. 解析代币对的 2 个地址

`Input Data` 结构

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

```bash
return token0 = input[10,74]  -- 0000000000000000000000008e0e57dcb1ce8d9091df38ec1bfc3b224529754a
```


```bash
return token1 = input[75,138] -- 000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7
```

3. 解析对子地址

解析规则

```bash
if logs[0] && logs[0].topics[0]='0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef'
    return logs[0].topics[2]
```

结果： `0x0000000000000000000000008ad599c3a0ff1de082011efddc58f1908eb6e6d8`


## 参考

1. Uniswap V3 交易 https://etherscan.io/tx/0xa5cb03f359691a57e27c29678204ac7623e7950684583fff48514338ad00502b
2. Uniswap V3 Swap 函数签名 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67
3. Uniswap V3 合约地址 https://etherscan.io/address/0xc36442b4a4522e871399cd717abdd847ab11fe88#code


### 添加流动池案例

1. https://etherscan.io/tx/0xfab2f1a51bcca936eb6c8067990992548a14bd84d6d4f3f99365e47d2614efb9
2. https://etherscan.io/tx/0x5850f95fd86a2e77196e0ecd71007da371fc44e902f344d5f4cb1feb046fe73e
3. https://etherscan.io/tx/0xdd257ef9c850b699881e6bb5a0e6da1559950f61cbfb52c8fb1e774b0e82be24
4. 
