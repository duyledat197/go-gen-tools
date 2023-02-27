package ethclient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type ETHClient struct {
	Address string
	Client  *ethclient.Client
	// abi       abi.ABI
	// contract  ethereum.Address
	Logger *zap.Logger
}

func (c *ETHClient) Init(ctx context.Context) error {
	client, err := ethclient.DialContext(ctx, c.Address)
	if err != nil {
		return err
	}
	c.Client = client
	return nil
}

func (c *ETHClient) CreateWallet() (*EthereumWallet, error) {
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

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &EthereumWallet{
		PrivateKey:       privateKey,
		PrivateKeyString: privKeyStr,
		PublicKey:        publicKeyECDSA,
		PublicKeyString:  pubKeyStr,
		Address:          address,
	}, nil
}

func (c *ETHClient) GetBalance(ctx context.Context, address string) (*Balance, error) {
	blockNumber, err := c.Client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	balance, err := c.Client.BalanceAt(ctx, common.HexToAddress(address), nil)

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
