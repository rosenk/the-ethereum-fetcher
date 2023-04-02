package fetcher

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
)

type getTransactionResult struct {
	transaction *ethereum.TransactionOverview
	err         error
}

func (c *Client) getTransactionRoutine(ctx context.Context, txHash common.Hash, results *[]getTransactionResult, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	transaction, err := c.GetTransaction(ctx, txHash)
	if err != nil {
		mutex.Lock()
		*results = append(*results, getTransactionResult{err: err})
		mutex.Unlock()

		return
	}

	if transaction == nil {
		return
	}

	mutex.Lock()
	*results = append(*results, getTransactionResult{transaction: transaction})
	mutex.Unlock()
}

func (c *Client) GetTransactions(ctx context.Context, txHashes []common.Hash) ([]ethereum.TransactionOverview, error) {
	var (
		mutex sync.Mutex
		wg    sync.WaitGroup
	)

	wg.Add(len(txHashes))

	results := make([]getTransactionResult, 0, len(txHashes))

	for _, txHash := range txHashes {
		go c.getTransactionRoutine(ctx, txHash, &results, &mutex, &wg)
	}

	wg.Wait()

	var transactions []ethereum.TransactionOverview

	for _, result := range results {
		if result.err != nil {
			return nil, result.err
		}

		transactions = append(transactions, *result.transaction)
	}

	return transactions, nil
}
