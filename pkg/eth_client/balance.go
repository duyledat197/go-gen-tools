package eth_client

import "math/big"

type Balance struct {
	BlockNumber uint64
	Value       *big.Float
}
