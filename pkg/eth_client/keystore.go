package eth_client

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func (c *ETHClient) NewKeystore(ctx context.Context, password string) (*accounts.Account, error) {
	key := keystore.NewKeyStore(c.KeyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := key.NewAccount(password)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *ETHClient) VerifyKeyStore(ctx context.Context, password string)
