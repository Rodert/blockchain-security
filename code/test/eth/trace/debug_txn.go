package trace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
	获取内部交易
	交易信息回放
- 如何获取以太坊内部交易（ETH转账）
*/
func GetInterTxn(txn string) []InternalTxn {
	time.Sleep(time.Second * 3)

	url := "https://eth-mainnet.g.alchemy.com/v2/dAnEgTX45tYXUXfNwtnXGRpJ-YiDBL7d"
	param := "{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"debug_traceTransaction\",\"params\":[\"%s\", {\"tracer\": \"callTracer\"}]}"
	payload := strings.NewReader(fmt.Sprintf(param, txn))
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("authority", "eth-mainnet.g.alchemy.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://composer.alchemy.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "macOS")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	utilsFun := UtilsFun{}

	if err := json.Unmarshal(body, &utilsFun); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("###\n%+v\n", utilsFun)
	var list []InternalTxn
	for _, call := range utilsFun.Result.Calls {
		if call.Type == "CALL" && call.Value != "0x0" {
			list = append(list, InternalTxn{
				From:  call.From,
				Gas:   call.Gas,
				To:    call.To,
				Value: call.Value,
			})
		}
	}
	fmt.Printf("###\n%+v\n", list)
	return list
}

// 内部交易 internal txn
type InternalTxn struct {
	From  string `json:"from"`
	Gas   string `json:"gas"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type UtilsFun struct {
	Jsonrpc string         `json:"jsonrpc"`
	Id      int            `json:"id"`
	Result  UtilsFunResult `json:"result"`
}

type UtilsFunResult struct {
	From    string                `json:"from"`
	Gas     string                `json:"gas"`
	GasUsed string                `json:"gasUsed"`
	To      string                `json:"to"`
	Input   string                `json:"input"`
	Output  string                `json:"output"`
	Calls   []UtilsFunResultCalls `json:"calls"`
	Value   string                `json:"value"`
	Type    string                `json:"type"`
}

type UtilsFunResultCalls struct {
	From    string                `json:"from"`
	Gas     string                `json:"gas"`
	GasUsed string                `json:"gasUsed"`
	To      string                `json:"to"`
	Input   string                `json:"input"`
	Output  string                `json:"output"`
	Calls   []UtilsFunResultCalls `json:"calls"`
	Value   string                `json:"value"`
	Type    string                `json:"type"`
}
