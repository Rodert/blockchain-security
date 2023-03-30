package spider

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"test/utils"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
)

const (
	FilePath = "/Users/shiyuwang/home/s/data/coinmarketcap"
)

var uri, _ = url.Parse("http://127.0.0.1:4780")
var client = &http.Client{
	Timeout: 20 * time.Second,
	Transport: &http.Transport{
		Proxy: http.ProxyURL(uri),
	},
}

func GetFileList(folder string) (fileList []string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		fileList = append(fileList, file.Name())
		// if file.IsDir() {
		// 	GetFileList(folder + "/" + file.Name())
		// } else {
		// 	fileList = append(fileList, file.Name())
		// }
	}
	return fileList
}

/* 采集html并保存 */
func AnalysisList() error {
	coinList := GetCoinList()
	var existCoinList []string
	for _, v := range GetFileList(FilePath + "/details") {
		existCoinList = append(existCoinList, strings.Split(v, ".")[0])
	}
	fmt.Println(existCoinList)
	fmt.Println(len(existCoinList))

	for _, v := range coinList {
		if utils.SliceContains(existCoinList, v) {
			continue
		}
		if err := GetProjectInfo(v); err != nil {
			fmt.Println(v)
			WriteFileAdd(FilePath+"/errUrl.txt", v+"\n")
		}
	}

	return nil
}

func AnalysisProject() error {

	return nil
}

func GetProjectInfo(coinPath string) error {
	var err error
	url := fmt.Sprintf("https://coinmarketcap.com/currencies/%v/", coinPath)
	fmt.Println(url)
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	if err != nil {
		log.Fatal(err)
		return err
	}

	res, err := client.Do(reqest)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	content := buf.String()
	WriteFile(fmt.Sprintf(FilePath+"/details/%v.html", coinPath), content)

	// utf8Body, _ := DecodeHTMLBody(res.Body, "")

	// doc, err := goquery.NewDocumentFromReader(utf8Body)
	// if err != nil {
	// 	return
	// }
	return nil
}

func WriteFile(filePath, content string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(content)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func WriteFileAdd(filePath, content string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}
}

func GetCoinList() []string {
	var coinList []string

	content, err := os.ReadFile(FilePath + "/coinmarketcap-list.txt")
	if err != nil {
		log.Fatal(err)
	}
	var cryptos Cryptos
	err = json.Unmarshal(content, &cryptos)
	if err != nil {
		return coinList
	}

	for index, v := range cryptos.Values {
		if index >= 9007 {
			break
		}
		coinList = append(coinList, utils.ToStr(v[3]))
	}
	return coinList
}

/* 自动识别编码格式 */
func DecodeHTMLBody(body io.Reader, charset string) (io.Reader, error) {
	if charset == "" {
		charset = detectContentCharset(body)
	}

	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}

	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(body)
	}

	return body, nil
}

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, _ := charset.DetermineEncoding(data, ""); len(name) != 0 {
			return name
		}
	}

	return "utf-8"
}

type (
	/* 币种列表 */
	Cryptos struct {
		Batch  string   `json:"batch"`
		Fields []string `json:"fields"`
		Values [][]any  `json:"values"`
	}
)
