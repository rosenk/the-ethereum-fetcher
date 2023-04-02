package validatorbuilder

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-playground/validator/v10"
)

//nolint:gochecknoglobals
var bakedInValidators = map[string]validator.Func{
	"min_hash_size": minHashSize,
	"max_hash_size": maxHashSize,
}

func minHashSize(fl validator.FieldLevel) bool {
	param := fl.Param()

	minSize, err := strconv.Atoi(param)
	if err != nil {
		return false
	}

	fmt.Println("minSize", minSize)
	fmt.Println("fl.Field().Interface()", fl.Field().Interface())

	hash, ok := fl.Field().Interface().(common.Hash)
	fmt.Println("hash", hash)
	fmt.Println("ok", ok)
	if !ok {
		return false
	}

	fmt.Println("hash.Big().BitLen()", hash.Big().BitLen())

	return hash.Big().BitLen() >= minSize
}

func maxHashSize(fl validator.FieldLevel) bool {
	param := fl.Param()

	maxSize, err := strconv.Atoi(param)
	if err != nil {
		return false
	}

	fmt.Println("minSize", maxSize)
	fmt.Println("fl.Field().Interface()", fl.Field().Interface())

	hash, ok := fl.Field().Interface().(common.Hash)
	fmt.Println("hash", hash)
	fmt.Println("ok", ok)
	if !ok {
		return false
	}

	fmt.Println("hash.Big().BitLen()", hash.Big().BitLen())

	return hash.Big().BitLen() <= maxSize
}
