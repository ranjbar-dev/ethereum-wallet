# ethereum-wallet
ethereum wallet package for creating and generating wallet, transferring ETH, getting wallet balance and crawling blocks to find wallet transactions


### Installation
```
go get github.com/ranjbar-dev/ethereum-wallet@v1.0.1
```

### Test
test for `Crawler`, `EthereumWallet` and `Token` located at /test
```
go test ./test
```


### Nodes 
using infura nodes in this repository 
register at infura and get your http and ws node on mainnet, goerli or sepolia network and set it here 
```
node := ethereumWallet.Node{
		Http: "https://goerli.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
		Ws:   "wss://goerli.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
}
```

### Wallet methods
- generating ethereum wallet
```
w := GenerateEthereumWallet(node)
w.Address // strnig 
w.PrivateKey // strnig 
w.PublicKey // strnig 
```
- creating ethereum wallet from private key
```
w := CreateEthereumWallet(node,privateKeyHex)
w.Address // strnig 
w.PrivateKey // strnig 
w.PublicKey // strnig 
```
- getting wallet ETH balance
```
balanceInWei,err := w.Balance()
balanceInWei // int64 
```
- getting wallet ERC20 balance
```
balanceInTokenSubAmount,err := w.BalanceERC20(token)
balanceInTokenSubAmount // int64 
```
- crawl blocks for addresses transactions
```

c := &Crawler{
		Node: node, 
		Addresses: []string{
			"0x5A2187B6d76a09F649CDC5d69F182697cFBA126B", // list of your addresses
		},
	}
	
res, err := c.ScanBlocks(40) // scan latest 40 block on block chain and extract addressess transactions 

Example 
// *
{
    {
        "address": "TY3PJu3VY8xVUc5BjYwJtyRgP7TfivV666",
        "tranasctions": {
            {
                "tx_id": "0xa1d76647e2c9ff4ceeda382137e7af9bb2879fce96cfc743d7138a8d54e532db",
                "from_address": "0x5a2187b6d76a09f649cdc5d69f182697cfba126b",
                "to_address": "0x75c07e7207bb00cf60c77f2506d7ce2b8d18bf0f",
                "amount": "21000000000",
                "confirmations": 2,
            }
        }
    },
    ...
}
* // 
	
```
- transfer ETH from wallet
```
txId, err := w.Transfer(toAddressHex, amount)
txId // string 
```
- transfer ERC20 from wallet
```
txId, err := w.TransferERC20(token, toAddressHex, amount)
txId // string 
```


### Token methods
- declaring token
```
token := &ethereumWallet.Token{
    ContractAddress: enums.GOERLI_USDC
}
```
- Getting token name
```
token.GetName(w.Node, w.AddressHex) // return string,error
``` 
- Getting token symbol
```
token.GetSymbol(w.Node, w.AddressHex) // return string,error
```
- Getting token decimals
```
token.GetDecimals(w.Node, w.AddressHex) // return int64,error
```

```
I simplified this repository github.com/ethereum/go-ethereum to create this package
You can check go it for better examples and functionalities
```

### Supported contracts
check enums/contracts file alternatively you can create your own contract
```
contractAddress := enums.CreateContractAddress("0x509Ee0d083DdF8AC028f2a56731412edD63223B9")
```

### ETH goerli Faucet
check this link https://goerlifaucet.com

### USDC goerli Faucet
check this link https://usdcfaucet.com/
 
### Important
I simplified this repository github.com/ethereum/go-ethereum repository to create this package You can check go it for better examples and functionalities and do not use this package in production, I created this package for education purposes.

### Donation
Address `0xe342C7ff6a91dF208Fb4bE6Eca02e19D624188C1`