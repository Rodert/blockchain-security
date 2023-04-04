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
