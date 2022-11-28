package test

import (
	ethereumWallet "github.com/ranjbar-dev/ethereum-wallet"
	"github.com/ranjbar-dev/ethereum-wallet/enums"
	"math/big"
)

var node = ethereumWallet.Node{
	Http: "https://goerli.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
	Ws:   "wss://goerli.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
}
var validPrivateKey = "a24031202755246def61140ae1bce297d0c4886b2faea5ce79001748ef97e8ec"
var invalidPrivateKey = "invalid"
var validOwnerAddress = "0x5A2187B6d76a09F649CDC5d69F182697cFBA126B"
var invalidOwnerAddress = "0x5A218742609F649CDC5d69F182697cFBA126B"
var validToAddress = "0x75c07e7207Bb00Cf60c77f2506D7CE2B8d18bf0f"
var invalidToAddress = "0x75c0721f25065D7CE2B8d18bf0f"
var ethAmount = big.NewInt(10000000000000) // 0.0001 ETH
var erc20Amount = big.NewInt(100000)       // 1 USDC

func wallet() *ethereumWallet.EthereumWallet {
	w, _ := ethereumWallet.CreateEthereumWallet(node, validPrivateKey)
	return w
}

func token() *ethereumWallet.Token {
	return &ethereumWallet.Token{
		ContractAddress: enums.GOERLI_USDC,
	}
}

func crawler() *ethereumWallet.Crawler {
	return &ethereumWallet.Crawler{
		Node:      node,
		Addresses: []string{validOwnerAddress},
	}
}
