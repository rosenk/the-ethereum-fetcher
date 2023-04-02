package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/ju-popov/the-ethereum-fetcher/internal/validatorbuilder"
	"github.com/sumup-oss/go-pkgs/errors"
)

func Validator() (*validator.Validate, error) {
	validate, err := validatorbuilder.Build()
	if err != nil {
		return nil, errors.Wrap(err, "create validator")
	}

	return validate, nil
}
