# input

> 通过解析合约创建时的交易 input，判断是否是 ERC 协议标准。


/*
ERC20 是以太坊上最常见的代币标准之一，它定义了代币合约的基本接口。以下是一些 ERC20 合约中一般会实现的必需函数及其函数签名：

totalSupply()：获取代币的总发行量，函数签名为：function totalSupply() public view returns (uint256);

balanceOf(address _owner)：获取指定地址 _owner 的代币余额，函数签名为：function balanceOf(address _owner) public view returns (uint256);

transfer(address _to, uint256 _value)：将指定数量 _value 的代币转移到地址 _to，函数签名为：function transfer(address _to, uint256 _value) public returns (bool);
0xa9059cbb

allowance(address _owner, address _spender)：获取地址 _owner 授权给地址 _spender 可以使用的代币数量，函数签名为：function allowance(address _owner, address _spender) public view returns (uint256);

approve(address _spender, uint256 _value)：授权地址 _spender 可以使用最多 _value 的代币，函数签名为：function approve(address _spender, uint256 _value) public returns (bool);

transferFrom(address _from, address _to, uint256 _value)：从地址 _from 向地址 _to 转移最多 _value 的代币，前提是地址 _from 已经授权地址 _msgSender() 可以使用相应数量的代币，函数签名为：function transferFrom(address _from, address _to, uint256 _value) public returns (bool);
0x23b872dd

注意，实现 ERC20 标准接口时，以上函数的函数签名是必需的，但是函数内部的具体实现方式并没有规定。因此，ERC20 合约的具体实现可以根据业务需求进行定制化，例如添加锁仓机制、增加发行新代币的接口等。

*/

/*
通常，ERC1155 合约会实现 safeBatchTransferFrom 和 safeTransferFrom 函数，而 ERC721 合约会实现 transferFrom 和 safeTransferFrom 函数。
*/


---

# 取链上数据

## infura 

http://cw.hubwiz.com/card/c/infura-api/1/2/12/

每天免费10w次，限制请求频率。

```bash
curl -X POST \
-x http://127.0.0.1:4780 \
-H "Content-Type: application/json" \
--data '{"jsonrpc": "2.0", "id": 1, "method": "eth_blockNumber", "params": []}' \
"https://mainnet.infura.io/v3/67f1fef3d3d0424fb7d99a14782c2322"
```


GetCode：

**请求参数：**
ADDRESS：要查询的地址，必需
BLOCK PARAMETER：区块编号，或者字符串"latest"、 "earliest"、 "pending"

```bash
curl https://mainnet.infura.io/v3/67f1fef3d3d0424fb7d99a14782c2322 \
    -X POST \
    -x http://127.0.0.1:4780 \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_getCode","params": 
    ["0xdAC17F958D2ee523a2206206994597C13D831ec7", "latest"],"id":1}'
```

---

## alchemy

https://docs.alchemy.com/reference/sdk-getcode

每天免费3亿次。

获取合约下有哪些代币
```bash
curl --request POST \
     --url https://eth-mainnet.g.alchemy.com/v2/Pk_IKdH-I_qEQ4uwQThNxkB-CD_6RHaY \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
     "id": 1,
     "jsonrpc": "2.0",
     "method": "alchemy_getTokenBalances",
     "params": [
          "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
     ]
}
'
```


获取代码：

> 代理加速，非必需

```bash
curl --request POST \
     --proxy http://127.0.0.1:4780 \
     --url https://eth-mainnet.g.alchemy.com/v2/Pk_IKdH-I_qEQ4uwQThNxkB-CD_6RHaY \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
     "id": 1,
     "jsonrpc": "2.0",
     "params": [
          "0xdac17f958d2ee523a2206206994597c13d831ec7",
          "latest"
     ],
     "method": "eth_getCode"
}
'
```