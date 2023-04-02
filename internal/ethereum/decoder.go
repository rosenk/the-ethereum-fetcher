package ethereum

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/sumup-oss/go-pkgs/errors"
)

func DecodeRLPEncodedHashes(encodedHashesHex string) ([]common.Hash, error) {
	encodedHashes, err := hex.DecodeString(encodedHashesHex)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode hex encoded hashes")
	}

	var decodedHashBytes [][]byte
	if err := rlp.DecodeBytes(encodedHashes, &decodedHashBytes); err != nil {
		return nil, errors.Wrap(err, "failed to decode RLP encoded hashes")
	}

	decodedHashes := make([]common.Hash, len(decodedHashBytes))
	for i, hashBytes := range decodedHashBytes {
		decodedHashes[i] = common.BytesToHash(hashBytes)
	}

	return decodedHashes, nil
}
