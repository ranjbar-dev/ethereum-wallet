package ethereumWallet

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ranjbar-dev/ethereum-wallet/enums"
	"github.com/ranjbar-dev/ethereum-wallet/geth"
	token "github.com/ranjbar-dev/ethereum-wallet/geth/contractErc20"
	"math/big"
)

type Token struct {
	ContractAddress enums.ContractAddress
}

func (t *Token) GetName(node Node) (string, error) {

	c, err := geth.GetGETHClient(node.Http)
	if err != nil {
		return "", err
	}

	contractAddress := common.HexToAddress(t.ContractAddress.String())
	tokenInstance, err := token.NewToken(contractAddress, c)
	if err != nil {
		return "", err
	}

	return tokenInstance.Name(&bind.CallOpts{})
}

func (t *Token) GetSymbol(node Node) (string, error) {

	c, err := geth.GetGETHClient(node.Http)
	if err != nil {
		return "", err
	}

	contractAddress := common.HexToAddress(t.ContractAddress.String())
	tokenInstance, err := token.NewToken(contractAddress, c)
	if err != nil {
		return "", err
	}

	return tokenInstance.Symbol(&bind.CallOpts{})
}

func (t *Token) GetDecimals(node Node) (uint8, error) {

	c, err := geth.GetGETHClient(node.Http)
	if err != nil {
		return 0, err
	}

	contractAddress := common.HexToAddress(t.ContractAddress.String())
	tokenInstance, err := token.NewToken(contractAddress, c)
	if err != nil {
		return 0, err
	}

	return tokenInstance.Decimals(&bind.CallOpts{})
}

func (t *Token) GetBalance(node Node, addressHex string) (*big.Int, error) {

	c, err := geth.GetGETHClient(node.Http)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(t.ContractAddress.String())
	address := common.HexToAddress(addressHex)
	tokenInstance, err := token.NewToken(contractAddress, c)
	if err != nil {
		return nil, err
	}

	return tokenInstance.BalanceOf(&bind.CallOpts{}, address)
}
