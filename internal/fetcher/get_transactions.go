package fetcher

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
)

type getTransactionRoutineResult struct {
	transaction *ethereum.TransactionFull
	err         error
}

func (c *Client) getTransactionRoutine(
	ctx context.Context,
	txHash common.Hash,
	results *[]getTransactionRoutineResult,
	mutex *sync.Mutex,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	transaction, err := c.GetTransaction(ctx, txHash)
	if err != nil {
		mutex.Lock()
		*results = append(*results, getTransactionRoutineResult{err: err})
		mutex.Unlock()

		return
	}

	if transaction == nil {
		return
	}

	mutex.Lock()
	*results = append(*results, getTransactionRoutineResult{transaction: transaction})
	mutex.Unlock()
}

func (c *Client) GetTransactionsSimultaneously(
	ctx context.Context,
	txHashes []common.Hash,
) ([]ethereum.TransactionFull, error) {
	var (
		mutex     sync.Mutex
		waitGroup sync.WaitGroup
	)

	waitGroup.Add(len(txHashes))

	results := make([]getTransactionRoutineResult, 0, len(txHashes))

	for _, txHash := range txHashes {
		go c.getTransactionRoutine(ctx, txHash, &results, &mutex, &waitGroup)
	}

	waitGroup.Wait()

	transactions := make([]ethereum.TransactionFull, 0, len(results))

	for _, result := range results {
		if result.err != nil {
			return nil, result.err
		}

		transactions = append(transactions, *result.transaction)
	}

	return transactions, nil
}
