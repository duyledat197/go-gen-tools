package eth_client

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey       *ecdsa.PrivateKey
	PrivateKeyString string
	PublicKey        *ecdsa.PublicKey
	PublicKeyString  string
	Address          common.Address
}

func (c *ETHClient) CreateWallet() (*Wallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("crypto.GenerateKey: %w", err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKeyStr := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKeyStr := hexutil.Encode(publicKeyBytes)[4:]

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Wallet{
		PrivateKey:       privateKey,
		PrivateKeyString: privKeyStr,
		PublicKey:        publicKeyECDSA,
		PublicKeyString:  pubKeyStr,
		Address:          address,
	}, nil
}
