package ethereumWallet

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ranjbar-dev/ethereum-wallet/geth"
	"math/big"
)

func estimateEthTransactionFee(node Node, toAddressHex string) (int64, error) {

	toAddress := common.HexToAddress(toAddressHex)

	client, err := geth.GetGETHClient(node.Http)
	if err != nil {
		return 0, err
	}

	blockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	currentBlock, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil {
		return 0, err
	}

	baseFee := currentBlock.BaseFee()

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To: &toAddress,
	})
	if err != nil {
		return 0, err
	}

	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return 0, err
	}

	fee := new(big.Int).SetInt64(baseFee.Int64() + gasTipCap.Int64())

	temp := new(big.Int).Mul(new(big.Int).SetInt64(int64(gasLimit)), fee)

	return temp.Int64(), nil
}

func estimateErc20TransactionFee(node Node) (int64, error) {

	client, err := geth.GetGETHClient(node.Http)
	if err != nil {
		return 0, err
	}

	blockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	currentBlock, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil {
		return 0, err
	}

	baseFee := currentBlock.BaseFee()
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return 0, err
	}

	gasLimit := 70000

	fee := new(big.Int).SetInt64(baseFee.Int64() + gasTipCap.Int64())

	temp := new(big.Int).Mul(new(big.Int).SetInt64(int64(gasLimit)), fee)

	return temp.Int64(), nil
}
