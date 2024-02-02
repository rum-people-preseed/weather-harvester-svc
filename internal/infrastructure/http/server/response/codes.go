package response

const (
	// numbering 100-1000
	GeneralCode_UnprocessableEntity   = 100
	GeneralCode_ResourceNotFound      = 101
	GeneralCode_MethodNotAllowed      = 102
	GeneralCode_NotAcceptable         = 103
	GeneralCode_GatewayTimeout        = 104
	GeneralCode_Unspecified           = 105
	GeneralCode_NotHTTPError          = 106
	GeneralCode_InternalHTTPError     = 107
	GeneralCode_ValidationFailedError = 108
	GeneralCode_BodyBindingFailed     = 109
)

// Field validation error codes
// Field validation error descriptors
const (
	ValidationCode_FieldRequired           = 1
	ValidationCode_FieldNotEmpty           = 2
	ValidationCode_FieldTypeISO3166_Alpha2 = 3
	ValidationCode_FieldTypeNumeric        = 4
	ValidationCode_FieldOneOf              = 5
	ValidationCode_FieldMaxSign            = 6
	ValidationCode_FieldMaxSigns           = 7
	ValidationCode_FieldMinSign            = 8
	ValidationCode_FieldMinSigns           = 9
	ValidationCode_FieldTypeISO4217        = 10
	ValidationCode_FieldMinDecimal         = 11
	ValidationCode_FieldMinDecimals        = 12
	ValidationCode_FieldMaxDecimal         = 13
	ValidationCode_FieldMaxDecimals        = 14
	// Validator uncompatible errors
	ValidationCode_FieldUnknownTagError = 9998
	ValidationCode_FieldUnknownError    = 9999
)

// Field validation param tagsf
const (
	ParamsTagOneOf = ":validOptions"
	ParamsTagMin   = ":min"
	ParamsTagMax   = ":max"
)
