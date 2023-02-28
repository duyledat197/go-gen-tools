package eth_client

import "crypto/ecdsa"

type EthereumWallet struct {
	PrivateKey       *ecdsa.PrivateKey
	PrivateKeyString string
	PublicKey        *ecdsa.PublicKey
	PublicKeyString  string
	Address          string
}
