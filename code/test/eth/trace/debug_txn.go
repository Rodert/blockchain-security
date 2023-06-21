package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
	获取内部交易

	交易信息回放


*/
func main() {
	url := "https://eth-mainnet.g.alchemy.com/v2/Pk_IKdH-I_qEQ4uwQThNxkB-CD_6RHaY"

	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"debug_traceTransaction\",\"params\":[\"0x8fc90a6c3ee3001cdcbbb685b4fbe67b1fa2bec575b15b0395fea5540d0901ae\", {\"tracer\": \"callTracer\"}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
