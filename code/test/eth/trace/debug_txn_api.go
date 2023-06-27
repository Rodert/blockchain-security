package trace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	RequestData struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Method  string `json:"method"`
		Params  Params `json:"params"`
	}
	Params struct {
		Hash         string       `json:"hash"`
		Tracer       string       `json:"tracer"`
		TracerConfig TracerConfig `json:"tracerConfig"`
	}
	TracerConfig struct {
		OnlyTopCall bool `json:"onlyTopCall"`
	}
)

func main() {
	url := "https://eth-mainnet.g.alchemy.com/v2/dAnEgTX45tYXUXfNwtnXGRpJ-YiDBL7d"
	reqData := RequestData{
		JSONRPC: "2.0",
		ID:      0,
		Method:  "debug_traceTransaction",
		Params: Params{
			Hash:   "0xaa900ed2777105ae16c4dcf2d19f515c46bb43480ca3029d3a9b214f6cbcaa16",
			Tracer: "callTracer",
			TracerConfig: TracerConfig{
				OnlyTopCall: false,
			},
		},
	}
	jsonData, _ := json.Marshal(reqData)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%+v\n", resp)
	// ...
}
