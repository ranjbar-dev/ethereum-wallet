package ethereumWallet

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ranjbar-dev/ethereum-wallet/enums"
	"github.com/ranjbar-dev/ethereum-wallet/geth"
	"github.com/ranjbar-dev/ethereum-wallet/util"
	"math/big"
	"strings"
	"sync"
	"time"
)

type Crawler struct {
	Node      Node
	Addresses []string
}

type CrawlResult struct {
	Address      string
	Transactions []CrawlTransaction
}

type CrawlTransaction struct {
	TxId          string
	Confirmations int64
	FromAddress   string
	ToAddress     string
	Amount        uint64
	Symbol        string
}

func (c *Crawler) ScanBlocks(count int) ([]CrawlResult, error) {

	var wg sync.WaitGroup

	var allTransactions [][]CrawlTransaction

	client, err := geth.GetGETHClient(c.Node.Http)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// check block for transaction
	allTransactions = append(allTransactions, c.extractOurTransactionsFromBlock(block, block))
	if err != nil {
		return nil, err
	}

	blockNumber := block.Number().Int64()

	for i := count; i > 0; i-- {
		wg.Add(1)
		blockNumber = blockNumber - 1
		// sleep to avoid 503 error
		time.Sleep(100 * time.Millisecond)
		go c.getBlockData(&wg, client, block, &allTransactions, blockNumber)
	}

	wg.Wait()

	return c.prepareCrawlResultFromTransactions(allTransactions), nil
}

func (c *Crawler) ScanBlocksFromTo(from int, to int) ([]CrawlResult, error) {

	if to-from < 1 {
		return nil, errors.New("to number should be more than from number")
	}

	var wg sync.WaitGroup

	var allTransactions [][]CrawlTransaction

	client, err := geth.GetGETHClient(c.Node.Http)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	for i := to; i > from; i-- {
		wg.Add(1)
		// sleep to avoid 503 error
		time.Sleep(100 * time.Millisecond)
		go c.getBlockData(&wg, client, block, &allTransactions, int64(i))
	}

	wg.Wait()

	return c.prepareCrawlResultFromTransactions(allTransactions), nil
}

func (c *Crawler) getBlockData(wg *sync.WaitGroup, client *ethclient.Client, currentBlock *types.Block, allTransactions *[][]CrawlTransaction, num int64) {

	defer wg.Done()

	block, err := client.BlockByNumber(context.Background(), big.NewInt(num))
	if err != nil {
		fmt.Println(err)
		return
	}

	// check block for transaction
	*allTransactions = append(*allTransactions, c.extractOurTransactionsFromBlock(currentBlock, block))
}

func (c *Crawler) extractOurTransactionsFromBlock(currentBlock *types.Block, block *types.Block) []CrawlTransaction {

	chainConfig := params.MainnetChainConfig
	if strings.Contains(c.Node.Http, "goerli") {
		chainConfig = params.GoerliChainConfig
	}
	if strings.Contains(c.Node.Http, "sepolia") {
		chainConfig = params.SepoliaChainConfig
	}

	blockSigner := types.MakeSigner(chainConfig, block.Number())

	var txs []CrawlTransaction

	for _, transaction := range block.Transactions() {

		symbol := "ETH"

		txMsg, errMessage := transaction.AsMessage(blockSigner, nil)
		if errMessage != nil {
			continue
		}

		fromAddress := txMsg.From().Hex()

		if txMsg.To() == nil {
			continue
		}
		toAddress := txMsg.To().Hex()

		if txMsg.Value() == nil {
			continue
		}

		amount := txMsg.Value().Int64()

		// is ERC20 token transfer
		if len(transaction.Data()) > 0 {

			tokenData, exist := util.ParsDataErc20TokenTransfer(transaction.Data())
			if !exist {
				continue
			}

			// toAddress -> contract address

			var err error
			token := Token{ContractAddress: enums.CreateContractAddress(toAddress)}
			symbol, err = token.GetSymbol(c.Node)
			if err != nil {
				continue
			}

			amount = tokenData.Value.Int64()
			toAddress = tokenData.To
		}

		txId := transaction.Hash().Hex()
		confirmations := int64(currentBlock.NumberU64() - block.NumberU64())

		for _, ourAddress := range c.Addresses {
			if ourAddress == toAddress || ourAddress == fromAddress {
				txs = append(txs, CrawlTransaction{
					TxId:          txId,
					FromAddress:   fromAddress,
					ToAddress:     toAddress,
					Amount:        uint64(amount),
					Confirmations: confirmations,
					Symbol:        symbol,
				})
			}
		}
	}

	return txs
}

func (c *Crawler) prepareCrawlResultFromTransactions(transactions [][]CrawlTransaction) []CrawlResult {

	var result []CrawlResult

	for _, transaction := range transactions {
		for _, tx := range transaction {

			if c.addressExistInResult(result, tx.ToAddress) {
				id, res := c.getAddressCrawlInResultList(result, tx.ToAddress)
				res.Transactions = append(res.Transactions, tx)
				result[id] = res

			} else {
				result = append(result, CrawlResult{
					Address:      tx.ToAddress,
					Transactions: []CrawlTransaction{tx},
				})
			}
		}
	}

	return result
}

func (c *Crawler) addressExistInResult(result []CrawlResult, address string) bool {
	for _, res := range result {
		if res.Address == address {
			return true
		}
	}
	return false
}

func (c *Crawler) getAddressCrawlInResultList(result []CrawlResult, address string) (int, CrawlResult) {
	for id, res := range result {
		if res.Address == address {
			return id, res
		}
	}
	panic("crawl result not found")
}
