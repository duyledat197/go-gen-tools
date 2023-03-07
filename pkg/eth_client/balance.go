package eth_client

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Balance struct {
	BlockNumber uint64
	Value       *big.Float
}

func (c *ETHClient) GetBalance(ctx context.Context, address string) (*Balance, error) {
	blockNumber, err := c.client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	balance, err := c.client.BalanceAt(ctx, common.HexToAddress(address), nil)

	fBalance, ok := new(big.Float).SetString(balance.String())
	if !ok {
		return nil, fmt.Errorf("unable to set balance to float")
	}

	val := fBalance.Quo(fBalance, big.NewFloat(math.Pow10(18)))
	return &Balance{
		BlockNumber: blockNumber,
		Value:       val,
	}, nil
}
