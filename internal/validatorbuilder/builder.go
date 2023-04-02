package validatorbuilder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

func Build() (*validator.Validate, error) {
	validate := validator.New()

	// Non-standard validators
	if err := validate.RegisterValidation("notblank", validators.NotBlank); err != nil {
		return nil, fmt.Errorf("failed to register non standard validators: %w", err)
	}

	// Custom validators
	for tag, fn := range bakedInValidators {
		if err := validate.RegisterValidation(tag, fn); err != nil {
			return nil, fmt.Errorf("failed to register custom validator: %w", err)
		}
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		nameParam := strings.SplitN(fld.Tag.Get("param"), ",", 2)[0] //nolint:gomnd
		if nameParam == "-" {
			return ""
		}
		if nameParam != "" {
			return nameParam
		}

		queryParam := strings.SplitN(fld.Tag.Get("query"), ",", 2)[0] //nolint:gomnd
		if queryParam == "-" {
			return ""
		}
		if queryParam != "" {
			return queryParam
		}

		jsonParam := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0] //nolint:gomnd
		if jsonParam == "-" {
			return ""
		}

		return jsonParam
	})

	return validate, nil
}
