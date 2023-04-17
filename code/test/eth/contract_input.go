package eth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

/* 方法签名转换 */
func InputAnalysis() {

	data := []byte("hello")

	/* ERC20 */
	data = []byte("totalSupply()")
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

}
