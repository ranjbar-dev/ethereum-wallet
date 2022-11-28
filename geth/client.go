package geth

import "github.com/ethereum/go-ethereum/ethclient"

func GetGETHClient(httpNode string) (*ethclient.Client, error) {

	client, err := ethclient.Dial(httpNode)

	if err != nil {
		return nil, err
	}

	return client, nil
}
