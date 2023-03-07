package eth_client

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type ETHClient struct {
	Address string
	client  *ethclient.Client
	abi     abi.ABI
	// contract ethereum.Address
	Logger      *zap.Logger
	KeyStoreDir string
	keyStore    *keystore.KeyStore
}

func (c *ETHClient) Connect(ctx context.Context) error {
	client, err := ethclient.DialContext(ctx, c.Address)
	if err != nil {
		return err
	}
	c.keyStore = keystore.NewKeyStore(c.KeyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	c.client = client
	return nil
}

func (c *ETHClient) Stop(ctx context.Context) error {
	c.client.Close()
	return nil
}
