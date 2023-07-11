> 搜索引擎搜 JavaPub 
> 这里希望通过区块的交易信息特征，解析出 uniswap V2 上的流动池价格比，减轻对节点的压力。




[toc]




> 必要解决的问题

TODO
1. 指定代币，如何获得流动池地址
2. 通过流动池地址如何获得代币
3. 通过流动池交易如何获得交易双方的数量
4. uniswapV3 的交易信息如何获取



### 识别流动池地址

#### 方案1

> ⚠️注意：这里是针对 Uniswap V2 的代币流动池。注意验证后使用

添加流动性，俩种方式：

创建流动池的uniswap合约地址： 0x7a250d5630b4cf539739df2c5dacb4c659f2488d

Function： addLiquidityETH
Function： addLiquidity

数据解析参考案例地址：
1. addLiquidityETH https://etherscan.io/tx/0x2af98e853d9c38976a8e3d135844999d0c327fd6439b49abf0143c6c761a7225
2. addLiquidity https://etherscan.io/tx/0x25692501b6940dfb8da3f5ce74a43404c93222133efd83c031025879137760dd


**1. 判断哪一条是添加流动池的交易哈希**

```bash
（addLiquidityETH：
    to="0x7a250d5630b4cf539739df2c5dacb4c659f2488d" and 
    input = "0xf305d719000000000000000000000000" + "%" and  （input前缀）
    len(input) = 394
）
或者
（addLiquidity：
    to="0x7a250d5630b4cf539739df2c5dacb4c659f2488d" and 
    input = "0xe8e33700000000000000000000000000" + "%" and  （input前缀）
    len(input) = 522
）
```



> 满足就可以判断这是一笔添加流动池的交易

![在这里插入图片描述](https://img-blog.csdnimg.cn/ceb7f9656b3f4a2881725448e5fab9fe.png)


```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "accessList": [],
        "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
        "blockNumber": "0x1041d31",
        "chainId": "0x1",
        "from": "0xfbfeaf0da0f2fde5c66df570133ae35f3eb58c9a",
        "gas": "0x32900e",
        "gasPrice": "0x6a657e2dc",
        "hash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
        "input": "0xf305d7190000000000000000000000006982508145454ce325ddbe47a25d4ec3d2311933000000000000000000000000000000000000134f796fac973ecf304af6000000000000000000000000000000000000000000134f796fac973ecf304af60000000000000000000000000000000000000000000000000000001bc16d674ec80000000000000000000000000000fbfeaf0da0f2fde5c66df570133ae35f3eb58c9a0000000000000000000000000000000000000000000000000000000064399277",
        "maxFeePerGas": "0x87c76f677",
        "maxPriorityFeePerGas": "0x5f5e100",
        "nonce": "0x4",
        "r": "0x323062a14fbf82fc2521b62ed404d38b3069ff563de63b0c761d188071aa78d3",
        "s": "0x620e132b168ebfb36fc60192095e6e44b3323026612e506847c3217411509389",
        "to": "0x7a250d5630b4cf539739df2c5dacb4c659f2488d",
        "transactionIndex": "0xb6",
        "type": "0x2",
        "v": "0x0",
        "value": "0x1bc16d674ec80000"
    }
}
```

**2. 获取流动池地址**


```bash
1. 取 logs 第一个对象；
2. len(data)=130；
3. data.sub[0,66] 等于流动池合约地址；
```


![在这里插入图片描述](https://img-blog.csdnimg.cn/b23cb92de2a5494ca409ec344480cf31.png)


```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
        "blockNumber": "0x1041d31",
        "contractAddress": null,
        "cumulativeGasUsed": "0xdeabee",
        "effectiveGasPrice": "0x6a657e2dc",
        "from": "0xfbfeaf0da0f2fde5c66df570133ae35f3eb58c9a",
        "gasUsed": "0x296d14",
        "logs": [
            {
                "address": "0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f",
                "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
                "blockNumber": "0x1041d31",
                "data": "0x000000000000000000000000a43fe16908251ee70ef74718545e4fe6c5ccec9f00000000000000000000000000000000000000000000000000000000000275da",
                "logIndex": "0xf4",
                "removed": false,
                "topics": [
                    "0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9",
                    "0x0000000000000000000000006982508145454ce325ddbe47a25d4ec3d2311933",
                    "0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
                ],
                "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
                "transactionIndex": "0xb6"
            },
            {
                "address": "0x6982508145454ce325ddbe47a25d4ec3d2311933",
                "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
                "blockNumber": "0x1041d31",
                "data": "0x000000000000000000000000000000000000134f796fac973ecf304af6000000",
                "logIndex": "0xf5",
                "removed": false,
                "topics": [
                    "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                    "0x000000000000000000000000fbfeaf0da0f2fde5c66df570133ae35f3eb58c9a",
                    "0x000000000000000000000000a43fe16908251ee70ef74718545e4fe6c5ccec9f"
                ],
                "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
                "transactionIndex": "0xb6"
            },
            {
                "address": "0x6982508145454ce325ddbe47a25d4ec3d2311933",
                "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
                "blockNumber": "0x1041d31",
                "data": "0x000000000000000000000000000000000000016e614438831900c82f5a000000",
                "logIndex": "0xf6",
                "removed": false,
                "topics": [
                    "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925",
                    "0x000000000000000000000000fbfeaf0da0f2fde5c66df570133ae35f3eb58c9a",
                    "0x0000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d"
                ],
                "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
                "transactionIndex": "0xb6"
            },
            {
                "address": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
                "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
                "blockNumber": "0x1041d31",
                "data": "0x0000000000000000000000000000000000000000000000001bc16d674ec80000",
                "logIndex": "0xf7",
                "removed": false,
                "topics": [
                    "0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c",
                    "0x0000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d"
                ],
                "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
                "transactionIndex": "0xb6"
            },
            {
                "address": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
                "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
                "blockNumber": "0x1041d31",
                "data": "0x0000000000000000000000000000000000000000000000001bc16d674ec80000",
                "logIndex": "0xf8",
                "removed": false,
                "topics": [
                    "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                    "0x0000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d",
                    "0x000000000000000000000000a43fe16908251ee70ef74718545e4fe6c5ccec9f"
                ],
                "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
                "transactionIndex": "0xb6"
            },
            {
                "address": "0xa43fe16908251ee70ef74718545e4fe6c5ccec9f",
                "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
                "blockNumber": "0x1041d31",
                "data": "0x00000000000000000000000000000000000000000000000000000000000003e8",
                "logIndex": "0xf9",
                "removed": false,
                "topics": [
                    "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                    "0x0000000000000000000000000000000000000000000000000000000000000000",
                    "0x0000000000000000000000000000000000000000000000000000000000000000"
                ],
                "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
                "transactionIndex": "0xb6"
            },
            {
                "address": "0xa43fe16908251ee70ef74718545e4fe6c5ccec9f",
                "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
                "blockNumber": "0x1041d31",
                "data": "0x0000000000000000000000000000000000000000001726ad4325f321caa6ed95",
                "logIndex": "0xfa",
                "removed": false,
                "topics": [
                    "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                    "0x0000000000000000000000000000000000000000000000000000000000000000",
                    "0x000000000000000000000000fbfeaf0da0f2fde5c66df570133ae35f3eb58c9a"
                ],
                "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
                "transactionIndex": "0xb6"
            },
            {
                "address": "0xa43fe16908251ee70ef74718545e4fe6c5ccec9f",
                "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
                "blockNumber": "0x1041d31",
                "data": "0x000000000000000000000000000000000000134f796fac973ecf304af60000000000000000000000000000000000000000000000000000001bc16d674ec80000",
                "logIndex": "0xfb",
                "removed": false,
                "topics": [
                    "0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"
                ],
                "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
                "transactionIndex": "0xb6"
            },
            {
                "address": "0xa43fe16908251ee70ef74718545e4fe6c5ccec9f",
                "blockHash": "0xdeb266f17a263a050035bd68619c928955a01deaeaa45367e360a3129ffef81c",
                "blockNumber": "0x1041d31",
                "data": "0x000000000000000000000000000000000000134f796fac973ecf304af60000000000000000000000000000000000000000000000000000001bc16d674ec80000",
                "logIndex": "0xfc",
                "removed": false,
                "topics": [
                    "0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f",
                    "0x0000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d"
                ],
                "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
                "transactionIndex": "0xb6"
            }
        ],
        "logsBloom": "0x000000080000000000000000800000000000000000000000200100000000010000000000000000000000000000000000020000000800000000000000002000020000000800000000000000080000008000000000040000000000000080000000000000000200000000000000000008000000004020080000000000100000000000000000000080000040004000000400000000010000000800000002000000000200000008000000000000000080000000000000000000000000000000800001010000120000000000000000000a00400000000000000010000000000080200000102000010000000a0000000000000000000000000000400000000000000000",
        "status": "0x1",
        "to": "0x7a250d5630b4cf539739df2c5dacb4c659f2488d",
        "transactionHash": "0x273894b35d8c30d32e1ffa22ee6aa320cc9f55f2adbba0583594ed47c031f6f6",
        "transactionIndex": "0xb6",
        "type": "0x2"
    }
}
```

#### 方案2

...

### 获取流动池对应的代币对

通过流动池交易的 data 参数来识别。

https://etherscan.io/tx/0x25692501b6940dfb8da3f5ce74a43404c93222133efd83c031025879137760dd

![在这里插入图片描述](https://img-blog.csdnimg.cn/f9893b835add491d92fc8a7107662b7b.png)

![在这里插入图片描述](https://img-blog.csdnimg.cn/ab32511a72de4db38b4e3ad5a536158f.png)



### 获得当前流动池代币价格比



