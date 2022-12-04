package ethereumWallet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ranjbar-dev/ethereum-wallet/geth"
	token "github.com/ranjbar-dev/ethereum-wallet/geth/contractErc20"
	"math/big"
)

type Node struct {
	Http string
	Ws   string
}

type EthereumWallet struct {
	Node       Node
	Address    string
	PrivateKey string
	PublicKey  string
}

func GenerateEthereumWallet(node Node) *EthereumWallet {

	privateKey, _ := generatePrivateKey()
	privateKeyHex := convertPrivateKeyToHex(privateKey)

	publicKey, _ := getPublicKeyFromPrivateKey(privateKey)
	publicKeyHex := convertPublicKeyToHex(publicKey)

	address := getAddressFromPublicKey(publicKey)

	return &EthereumWallet{
		Node:       node,
		Address:    address,
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
	}
}

func CreateEthereumWallet(node Node, privateKeyHex string) (*EthereumWallet, error) {

	privateKey, err := privateKeyFromHex(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey, _ := getPublicKeyFromPrivateKey(privateKey)
	publicKeyHex := convertPublicKeyToHex(publicKey)

	address := getAddressFromPublicKey(publicKey)

	return &EthereumWallet{
		Node:       node,
		Address:    address,
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
	}, nil
}

// struct functions

func (ew *EthereumWallet) PrivateKeyRCDSA() (*ecdsa.PrivateKey, error) {
	return privateKeyFromHex(ew.PrivateKey)
}

func (ew *EthereumWallet) PrivateKeyBytes() ([]byte, error) {

	priv, err := ew.PrivateKeyRCDSA()
	if err != nil {
		return []byte{}, err
	}

	return crypto.FromECDSA(priv), nil
}

// private key

func generatePrivateKey() (*ecdsa.PrivateKey, error) {

	return crypto.GenerateKey()
}

func convertPrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {

	privateKeyBytes := crypto.FromECDSA(privateKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

func privateKeyFromHex(hex string) (*ecdsa.PrivateKey, error) {

	return crypto.HexToECDSA(hex)
}

// public key

func getPublicKeyFromPrivateKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error in getting public key")
	}

	return publicKeyECDSA, nil
}

func convertPublicKeyToHex(publicKey *ecdsa.PublicKey) string {

	privateKeyBytes := crypto.FromECDSAPub(publicKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

// address

func getAddressFromPublicKey(publicKey *ecdsa.PublicKey) string {

	return crypto.PubkeyToAddress(*publicKey).Hex()
}

// balance

func (ew *EthereumWallet) Balance() (int64, error) {

	c, err := geth.GetGETHClient(ew.Node.Http)
	if err != nil {
		return 0, err
	}

	balance, err := c.BalanceAt(context.Background(), common.HexToAddress(ew.Address), nil)
	if err != nil {
		return 0, err
	}

	return balance.Int64(), nil
}

func (ew *EthereumWallet) BalanceERC20(token *Token) (int64, error) {

	balance, err := token.GetBalance(ew.Node, ew.Address)
	if err != nil {
		return 0, err
	}

	return balance.Int64(), nil
}

// transaction

func (ew *EthereumWallet) Transfer(toAddressHex string, amountInWei *big.Int) (string, error) {

	privateRCDSA, err := ew.PrivateKeyRCDSA()
	if err != nil {
		return "", fmt.Errorf("RCDSA private key error: %v", err)
	}

	tx, err := createTransactionInput(ew.Node, ew.Address, toAddressHex, amountInWei)
	if err != nil {
		return "", fmt.Errorf("creating tx pb error: %v", err)
	}

	tx, err = signTransaction(ew.Node, tx, privateRCDSA)
	if err != nil {
		return "", fmt.Errorf("signing transaction error: %v", err)
	}

	txId, err := broadcastTransaction(ew.Node, tx)
	if err != nil {
		return "", fmt.Errorf("broadcast transaction error: %v", err)
	}

	return txId, nil
}

func (ew *EthereumWallet) EstimateTransferFee(toAddressHex string) (int64, error) {

	return estimateEthTransactionFee(ew.Node, toAddressHex)
}

func (ew *EthereumWallet) TransferERC20(t *Token, toAddressHex string, amountInTokenSubAmount *big.Int) (string, error) {

	c, err := geth.GetGETHClient(ew.Node.Http)
	if err != nil {
		return "", err
	}

	contractAddress := common.HexToAddress(t.ContractAddress.String())
	toAddress := common.HexToAddress(toAddressHex)
	tokenInstance, err := token.NewToken(contractAddress, c)
	if err != nil {
		return "", err
	}

	txInput, err := createERC20Transaction(ew.Node, toAddressHex, c, ew)
	if err != nil {
		return "", err
	}

	tx, err := tokenInstance.Transfer(txInput, toAddress, amountInTokenSubAmount)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (ew *EthereumWallet) EstimateTransferERC20Fee(t *Token) (int64, error) {

	return estimateErc20TransactionFee(ew.Node)
}
