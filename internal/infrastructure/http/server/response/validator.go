package response

import (
	"github.com/go-playground/validator/v10"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces/httpx"
)

var (
	// Tag to field resolving map
	tagToErrorCode = map[string]func(validator.FieldError) int{
		"required":         func(validator.FieldError) int { return ValidationCode_FieldRequired },
		"iso3166_1_alpha2": func(validator.FieldError) int { return ValidationCode_FieldTypeISO3166_Alpha2 },
		"numeric":          func(validator.FieldError) int { return ValidationCode_FieldTypeNumeric },
		"oneof":            func(validator.FieldError) int { return ValidationCode_FieldOneOf },
		"min":              fieldMinCodeResolver,
		"max":              fieldMaxCodeResolver,
		"iso4217":          func(validator.FieldError) int { return ValidationCode_FieldTypeISO4217 },
		"decimalsMin":      fieldMinDecimalsCodeResolver,
		"decimalsMax":      fieldMaxDecimalsCodeResolver,
	}
	fieldUnknownTagFunc = func(validator.FieldError) int { return ValidationCode_FieldUnknownTagError }

	// Tag to params map resolving map
	tagToParamsFunc = map[string]func(string) map[string]string{
		"oneof":       oneOfParams,
		"min":         minParams,
		"max":         maxParams,
		"decimalsMin": minParams,
		"decimalsMax": maxParams,
	}

	// helper for empty params, when none found in tagToParamsFunc
	emptyParamsFunc = func(string) map[string]string {
		return map[string]string{}
	}

	// common response for non validator.ValidationErrors typed errors
	unknownErrorFieldErrors = []httpx.InnerError{
		{
			Field:  "-",
			Code:   ValidationCode_FieldUnknownError,
			Params: map[string]string{},
		},
	}
)

func ParseFieldError(validationErr validator.FieldError) httpx.InnerError {
	tag := validationErr.Tag()
	tagCodeFunc, ok := tagToErrorCode[tag]
	if !ok {
		tagCodeFunc = fieldUnknownTagFunc
	}
	paramsFunc, ok := tagToParamsFunc[tag]
	if !ok {
		paramsFunc = emptyParamsFunc
	}
	return httpx.InnerError{
		Field:  validationErr.Field(),
		Code:   tagCodeFunc(validationErr),
		Params: paramsFunc(validationErr.Param()),
	}
}

// ==================== params map parsing functions ====================//

func oneOfParams(params string) map[string]string {
	return map[string]string{
		ParamsTagOneOf: params,
	}
}

func minParams(params string) map[string]string {
	return map[string]string{
		ParamsTagMin: params,
	}
}

func maxParams(params string) map[string]string {
	return map[string]string{
		ParamsTagMax: params,
	}
}

// =============== field resolution custom functions ================/

func fieldMinCodeResolver(fe validator.FieldError) int {
	if fe.Param() == "1" {
		return ValidationCode_FieldMinSign
	}
	return ValidationCode_FieldMinSigns
}

func fieldMaxCodeResolver(fe validator.FieldError) int {
	if fe.Param() == "1" {
		return ValidationCode_FieldMinSign
	}
	return ValidationCode_FieldMinSigns
}

func fieldMinDecimalsCodeResolver(fe validator.FieldError) int {
	if fe.Param() == "1" {
		return ValidationCode_FieldMinDecimal
	}
	return ValidationCode_FieldMinDecimals
}

func fieldMaxDecimalsCodeResolver(fe validator.FieldError) int {
	if fe.Param() == "1" {
		return ValidationCode_FieldMaxDecimal
	}
	return ValidationCode_FieldMaxDecimals
}
