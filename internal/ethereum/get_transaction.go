package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sumup-oss/go-pkgs/errors"
)

type Transaction struct {
	Hash   common.Hash
	Sender common.Address
	To     *common.Address
	Data   []byte
	Value  *big.Int
}

func (c *Client) GetTransaction(ctx context.Context, txHash common.Hash) (*Transaction, error) {
	hash := common.HexToHash(txHash.String())

	transaction, _, err := c.client.TransactionByHash(ctx, hash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to get transaction %v", txHash)
	}

	signer := types.LatestSignerForChainID(transaction.ChainId())

	sender, err := types.Sender(signer, transaction)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get 'from' address")
	}

	result := &Transaction{
		Hash:   transaction.Hash(),
		Sender: sender,
		To:     transaction.To(),
		Data:   transaction.Data(),
		Value:  transaction.Value(),
	}

	return result, nil
}
