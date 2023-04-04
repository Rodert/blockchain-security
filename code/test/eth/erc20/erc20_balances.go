package erc20

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// "sdp/internal/utils"
)

type balanceResult struct {
	Balance *big.Int
	Error   error
}

func BatchGetBalances(client *ethclient.Client, accounts []common.Address) (map[common.Address]*big.Int, error) {
	ctx := context.TODO()

	bals := make(chan balanceResult, len(accounts))
	for _, a := range accounts {
		go func(a common.Address) {
			bal, err := client.BalanceAt(ctx, a, nil)
			bals <- balanceResult{Balance: bal, Error: err}
		}(a)
	}

	balances := make(map[common.Address]*big.Int)
	for i := 0; i < len(accounts); i++ {
		result := <-bals
		if result.Error != nil {
			return nil, result.Error
		}
		balances[accounts[i]] = result.Balance
	}

	return balances, nil
}

// ---

// func NewAgent(rawURL string) (*Agent, error) {
// 	client, err := rpc.DialContext(context.Background(), rawURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Agent{client: client}, nil
// }

// type Agent struct {
// 	client *rpc.Client
// }

// func (a *Agent) Close() {
// 	a.client.Close()
// }

// func (agent *Agent) toBlockNumArg(number *big.Int) string {
// 	if number == nil {
// 		return "latest"
// 	}
// 	if number.Cmp(big.NewInt(-1)) == 0 {
// 		return "pending"
// 	}
// 	return hexutil.EncodeBig(number)
// }

// func (agent *Agent) GetBlockNumber() (uint64, error) {
// 	var rawBlockNumber string
// 	err := agent.client.Call(&rawBlockNumber, "eth_blockNumber")
// 	if err != nil {
// 		return 0, err
// 	}

// 	return hexutil.DecodeUint64(rawBlockNumber)
// }

// func (agent *Agent) GetBlockByNumber(blockNumber int64) (*Block, error) {
// 	bn := big.Int{}
// 	block := &Block{}
// 	err := agent.client.CallContext(
// 		context.Background(),
// 		block,
// 		"eth_getBlockByNumber",
// 		agent.toBlockNumArg(bn.SetInt64(blockNumber)),
// 		true,
// 	)
// 	if err != nil {
// 		return block, err
// 	}
// 	return block, nil
// }

// func (agent *Agent) batchTransactionReceipts(transactions []Transaction) ([]rpc.BatchElem, error) {
// 	batchElemList := []rpc.BatchElem{}
// 	receiptMethod := "eth_getTransactionReceipt"
// 	for _, transaction := range transactions {
// 		batchElemList = append(batchElemList, rpc.BatchElem{
// 			Method: receiptMethod,
// 			Args:   []any{transaction.Hash},
// 			// Result: &Receipt{},
// 		})
// 	}
// 	if err := agent.client.BatchCallContext(context.Background(), batchElemList); err != nil {
// 		return nil, err
// 	}

// 	return batchElemList, nil
// }

// func (agent *Agent) GetContractTransaction(block *Block) ([]*Receipt, error) {
// 	l := len(block.Transactions)
// 	receipts := make([]*Receipt, 0, l)
// 	if l == 0 {
// 		return receipts, nil
// 	}

// 	batchElemList, err := agent.batchTransactionReceipts(block.Transactions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, batchElem := range batchElemList {
// 		if receipt, ok := batchElem.Result.(*Receipt); ok {
// 			receipts = append(receipts, receipt)
// 		}
// 	}
// 	return receipts, nil
// }

// func (agent *Agent) GetBlockInfoByBlockNumber(blockNumber int64) (Data, error) {
// 	b := Data{
// 		BlockNumber: blockNumber,
// 	}
// 	var err error
// 	// 获取区块信息
// 	b.Block, err = agent.GetBlockByNumber(blockNumber)
// 	if err != nil {
// 		return b, err
// 	}
// 	b.Receipts, err = agent.GetContractTransaction(b.Block)
// 	if err != nil {
// 		return b, err
// 	}
// 	_b := &b
// 	err = _b.Check()
// 	return b, err
// }

// func (agent *Agent) GetGzipDataByBlockNumber(blockNumber int64) ([]byte, error) {
// 	data, err := agent.GetBlockInfoByBlockNumber(blockNumber)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	dataGzip, err := MarshalToJsonWithGzip(&data)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	log.Infof("block number:%d,transaction quantity:%d,after gzip:%dkb", blockNumber, len(data.Block.Transactions), len(dataGzip)/1024)
// 	return dataGzip, nil
// }

// // 压缩
// func MarshalToJsonWithGzip(jsonData *Data) ([]byte, error) {
// 	dataAfterMarshal, err := json.Marshal(jsonData)
// 	if err != nil {
// 		return dataAfterMarshal, err
// 	}
// 	return utils.Encode(dataAfterMarshal)
// }

// // 解压
// func UnmarshalDataFromJsonWithGzip(msg []byte) (*Data, error) {
// 	dataAfterDecode, err := utils.Decode(msg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	data := &Data{}
// 	err = json.Unmarshal(dataAfterDecode, data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }
