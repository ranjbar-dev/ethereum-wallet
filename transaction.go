package ethereumWallet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ranjbar-dev/ethereum-wallet/geth"
	"math/big"
)

func createTransactionInput(node Node, fromAddressHex string, toAddressHex string, amountInWei *big.Int) (*types.Transaction, error) {

	fromAddress := common.HexToAddress(fromAddressHex)
	toAddress := common.HexToAddress(toAddressHex)

	client, err := geth.GetGETHClient(node.Http)
	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(21000)
	gasFeeCap := big.NewInt(10 * params.GWei)
	gasTipCap := big.NewInt(10 * params.GWei)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &toAddress,
		Value:     amountInWei,
		Gas:       gasLimit,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Data:      nil,
	}), nil
}

func signTransaction(node Node, transaction *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {

	client, err := geth.GetGETHClient(node.Http)
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	signer := types.LatestSignerForChainID(chainID)

	return types.SignTx(transaction, signer, privateKey)
}

func broadcastTransaction(node Node, transaction *types.Transaction) (string, error) {

	client, err := geth.GetGETHClient(node.Http)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), transaction)
	if err != nil {
		return "", err
	}

	fmt.Println(transaction)

	return transaction.Hash().Hex(), nil
}

func createERC20Transaction(c *ethclient.Client, ew *EthereumWallet) (*bind.TransactOpts, error) {

	privateRCDSA, err := ew.PrivateKeyRCDSA()
	if err != nil {
		return nil, fmt.Errorf("RCDSA private key error: %v", err)
	}

	fromAddress := common.HexToAddress(ew.Address)
	n, err := c.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return nil, err
	}

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	signer := types.LatestSignerForChainID(chainID)

	return &bind.TransactOpts{
		From:  fromAddress,
		Nonce: big.NewInt(int64(n)),
		Signer: func(addr common.Address, localTx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(localTx, signer, privateRCDSA)
		},
	}, nil
}
