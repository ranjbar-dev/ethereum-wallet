package main

import (
	"fmt"
	ethereumWallet "github.com/ranjbar-dev/ethereum-wallet"
	"github.com/ranjbar-dev/ethereum-wallet/enums"
	"math/big"
)

//const fromAddress = "0x5A2187B6d76a09F649CDC5d69F182697cFBA126B"
const toAddress = "0x75c07e7207Bb00Cf60c77f2506D7CE2B8d18bf0f"

//const PrivateKey = "a24031202755246def61140ae1bce297d0c4886b2faea5ce79001748ef97e8ec"
//const HTTP_NODE = "https://goerli.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd"
//const WS_NODE = "wss://goerli.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd"

func main() {

	crawl()
}

func crawl() {
	var node = ethereumWallet.Node{
		Http: "https://goerli.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
		Ws:   "wss://goerli.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
	}
	var validPrivateKey = "88414dbb373a211bc157265a267f3de6a4cec210f3a5da12e89630f2c447ad27"
	w, _ := ethereumWallet.CreateEthereumWallet(node, validPrivateKey)

	fmt.Println(w.Transfer("0x4d496ccc28058b1d74b7a19541663e21154f9c84", big.NewInt(1000000)))
}

func transfer() {
	var node = ethereumWallet.Node{
		Http: "https://goerli.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
		Ws:   "wss://goerli.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
	}
	var validPrivateKey = "a24031202755246def61140ae1bce297d0c4886b2faea5ce79001748ef97e8ec"
	//var validToAddress = "0x75c07e7207Bb00Cf60c77f2506D7CE2B8d18bf0f"
	var ethAmount = big.NewInt(1000000) // 1 USDC

	w, _ := ethereumWallet.CreateEthereumWallet(node, validPrivateKey)

	t := &ethereumWallet.Token{ContractAddress: enums.GOERLI_USDC}

	fmt.Println(w.TransferERC20(t, toAddress, ethAmount))
	//fmt.Println(w.Transfer(validToAddress, ethAmount))
}

//func nonce() {
//	client, _ := geth.GetGETHClient(HTTP_NODE)
//
//	n, _ := client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
//	pn, _ := client.PendingNonceAt(context.Background(), common.HexToAddress(fromAddress))
//
//	fmt.Println("===============")
//	fmt.Println("===Nonce==")
//	fmt.Println(n)
//	fmt.Println("===Pending Nonce==")
//	fmt.Println(pn)
//	fmt.Println("===============")
//
//}
//
//func crawl() {
//
//	c := ethereumWallet.Crawler{
//		Node: ethereumWallet.Node{
//			Http: HTTP_NODE,
//			Ws:   WS_NODE,
//		},
//		Addresses: []string{fromAddress},
//	}
//
//	res, err := c.ScanBlocks(10)
//
//	fmt.Println(res)
//	fmt.Println(err)
//
//}
//
//func transfer() {
//
//	node := ethereumWallet.Node{
//		Http: HTTP_NODE,
//		Ws:   WS_NODE,
//	}
//
//	w := ethereumWallet.CreateEthereumWallet(node, PrivateKey)
//
//	txId, err := w.Transfer(toAddress, big.NewInt(100000000000000)) // 0.0001 eth
//
//	fmt.Println(txId)
//	fmt.Println(err)
//
//}
//
//func transferERC20() {
//
//	node := ethereumWallet.Node{
//		Http: HTTP_NODE,
//		Ws:   WS_NODE,
//	}
//
//	w := ethereumWallet.CreateEthereumWallet(node, PrivateKey)
//
//	token := &ethereumWallet.Token{ContractAddress: enums.GOERLI_USDC}
//
//	txId, err := w.TransferERC20(token, toAddress, big.NewInt(1000000)) //  1 USDC
//
//	fmt.Println(txId)
//	fmt.Println(err)
//
//}
