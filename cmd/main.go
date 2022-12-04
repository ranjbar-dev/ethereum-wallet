package main

import (
	"fmt"
	EthereumWallet "github.com/ranjbar-dev/ethereum-wallet"
	"github.com/ranjbar-dev/ethereum-wallet/enums"
	"math/big"
)

var node = EthereumWallet.Node{
	Http: "https://goerli.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
	Ws:   "wss://goerli.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
}

func main() {

	token := &EthereumWallet.Token{ContractAddress: enums.GOERLI_USDC}
	w, _ := EthereumWallet.CreateEthereumWallet(node, "a24031202755246def61140ae1bce297d0c4886b2faea5ce79001748ef97e8ec")
	fmt.Println(w.EstimateTransferERC20Fee(token))
	fmt.Println(w.TransferERC20(token, "0x75c07e7207Bb00Cf60c77f2506D7CE2B8d18bf0f", big.NewInt(10)))
	return
}
