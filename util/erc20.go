package util

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

const Erc20TransferMethodSignature = "a9059cbb"

type Erc20TokenTransfer struct {
	To    string
	Value big.Int
}

func ParsDataErc20TokenTransfer(data []byte) (Erc20TokenTransfer, bool) {
	dataHex := hex.EncodeToString(data)

	isDataErc20TokenTransfer := isDataErc20TokenTransfer(dataHex)

	if !isDataErc20TokenTransfer {
		return Erc20TokenTransfer{}, false
	}

	addressNonPaddedHex := dataHex[len(Erc20TransferMethodSignature) : 64+len(Erc20TransferMethodSignature)]
	address, err := gainAddressFromPaddedHex(addressNonPaddedHex)
	if err != nil {
		return Erc20TokenTransfer{}, false
	}

	valueStr := dataHex[64+len(Erc20TransferMethodSignature):]
	value := new(big.Int)
	value.SetString(valueStr, 16)

	tokenData := Erc20TokenTransfer{
		To:    address,
		Value: *value,
	}

	return tokenData, true
}

func isDataErc20TokenTransfer(dataHex string) bool {
	if len(dataHex) == 136 && strings.HasPrefix(dataHex, Erc20TransferMethodSignature) {
		return true
	}
	return false
}

func gainAddressFromPaddedHex(s string) (string, error) {
	var t big.Int
	var ok bool
	if has0xPrefix(s) {
		_, ok = t.SetString(s[2:], 16)
	} else {
		_, ok = t.SetString(s, 16)
	}
	if !ok {
		return "", errors.New("data is not a number")
	}
	a := common.BigToAddress(&t)
	return a.String(), nil
}

func has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1]|32) == 'x'
}
