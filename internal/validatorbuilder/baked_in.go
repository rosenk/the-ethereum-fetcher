package validatorbuilder

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-playground/validator/v10"
)

//nolint:gochecknoglobals
var bakedInValidators = map[string]validator.Func{
	"min_hash_size": minHashSize,
	"max_hash_size": maxHashSize,
}

func minHashSize(fieldLevel validator.FieldLevel) bool {
	param := fieldLevel.Param()

	minSize, err := strconv.Atoi(param)
	if err != nil {
		return false
	}

	hash, ok := fieldLevel.Field().Interface().(common.Hash)
	if !ok {
		return false
	}

	return hash.Big().BitLen() >= minSize
}

func maxHashSize(fieldLevel validator.FieldLevel) bool {
	param := fieldLevel.Param()

	maxSize, err := strconv.Atoi(param)
	if err != nil {
		return false
	}

	hash, ok := fieldLevel.Field().Interface().(common.Hash)
	if !ok {
		return false
	}

	return hash.Big().BitLen() <= maxSize
}
