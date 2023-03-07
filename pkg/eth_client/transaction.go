package eth_client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *ETHClient) CreateTransaction(ctx context.Context, from, to *Wallet, value *big.Int) (uint64, error) {
	nonce, err := c.client.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		return 0, err
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return 0, err
	}

	var data []byte
	tx := types.NewTransaction(nonce, to.Address, value, gasLimit, gasPrice, data)

	chainID, err := c.client.NetworkID(context.Background())
	if err != nil {
		return 0, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), from.PrivateKey)
	if err != nil {
		return 0, err
	}

	if err := c.client.SendTransaction(context.Background(), signedTx); err != nil {
		return 0, err
	}

	c.Logger.Sugar().Infof("tx sent: %s\n", signedTx.Hash().Hex())
	return nonce, nil
}
