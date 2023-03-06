package main

/* 解析 input 判断合约协议 */

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

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	data := []byte("hello")

	/* ERC20 */
	// data = []byte("totalSupply()")
	// data = []byte("balanceOf(address)")
	// data = []byte("transfer(address,uint256)")
	// data = []byte("allowance(address,address)")
	// data = []byte("approve(address,uint256)")
	// data = []byte("transferFrom(address,address,uint256)")

	/* ERC721 */
	// data = []byte("transferFrom(address,address,uint256)")
	// data = []byte("safeTransferFrom(address,address,uint256,bytes)")

	/* ERC1155 */
	// data = []byte("TransferSingle(address,address,address,uint256,uint256)") // 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62
	// data = []byte("TransferBatch(address,address,address,uint256[],uint256[])") // 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb
	// data = []byte("ApprovalForAll(address,address,bool)") // 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31
	// data = []byte("URI(string,uint256)") // 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b
	// data = []byte("safeTransferFrom(address,address,uint256,uint256,bytes)")
	// data = []byte("safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)")
	// data = []byte("setApprovalForAll(address,bool)")
	// data = []byte("onERC1155Received(address,address,uint256,uint256,bytes)")
	// data = []byte("onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)")

	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301
}
