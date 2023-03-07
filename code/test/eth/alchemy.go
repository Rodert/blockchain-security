package eth

/*

基于 alchemy 查询 ETH

https://docs.alchemy.com/reference/eth-getcode

面板：https://dashboard.alchemy.com/
*/

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/* 获取合约代码 */
func GetCodeByAlchemy() {
	fmt.Println("### select contract by alchemy.")

	url := "https://eth-mainnet.g.alchemy.com/v2/Pk_IKdH-I_qEQ4uwQThNxkB-CD_6RHaY"

	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"params\":[\"0xdac17f958d2ee523a2206206994597c13d831ec7\",\"latest\"],\"method\":\"eth_getCode\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
