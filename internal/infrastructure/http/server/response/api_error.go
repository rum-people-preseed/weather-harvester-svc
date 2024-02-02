package response

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces/errors"
)

var (
	ApiErrorUnprocessableEntity  = errors.New(http.StatusUnprocessableEntity, GeneralCode_UnprocessableEntity, "Unprocessable Entity")
	ApiErrorResourceNotFound     = errors.New(http.StatusNotFound, GeneralCode_ResourceNotFound, "Resource not found")
	ApiErrorMethodNotAllowed     = errors.New(http.StatusMethodNotAllowed, GeneralCode_MethodNotAllowed, "Method Not Allowed")
	ApiErrorNotAcceptable        = errors.New(http.StatusNotAcceptable, GeneralCode_NotAcceptable, "Not Acceptable")
	ApiErrorGatewayTimeout       = errors.New(http.StatusGatewayTimeout, GeneralCode_NotHTTPError, "Gateway Timeout")
	ApiErrorUnspecified          = errors.New(http.StatusNotAcceptable, GeneralCode_Unspecified, "Unspecified")
	ApiErrorNotHTTPError         = errors.New(http.StatusNotFound, GeneralCode_NotHTTPError, "Not Http Error")
	ApiErrorInternalServerError  = errors.New(http.StatusInternalServerError, GeneralCode_InternalHTTPError, "Internal Server Error")
	ApiErrorValidationBodyFailed = errors.New(http.StatusBadRequest, GeneralCode_ValidationFailedError, "Validation Failed")
	ApiErrorBodyBindingFailed    = errors.New(http.StatusUnprocessableEntity, GeneralCode_BodyBindingFailed, "Body Binding Failed")
)

func ApiErrorMapping(echoError *echo.HTTPError) errors.CodeMsg {
	switch echoError.Code {
	case http.StatusNotFound:
		// ApiErrorResourceNotFound from errors.CodeMsg
		return errors.CodeMsg{
			Code:    http.StatusNotFound,
			Message: "Resource not found",
		}
	case http.StatusMethodNotAllowed:
		return errors.CodeMsg{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
		}
	case http.StatusNotAcceptable:
		return errors.CodeMsg{
			Code:    http.StatusNotAcceptable,
			Message: "Not acceptable",
		}
	case http.StatusGatewayTimeout:
		return errors.CodeMsg{
			Code:    http.StatusGatewayTimeout,
			Message: "Gateway timeout",
		}
	default:
		return errors.CodeMsg{
			Code:    http.StatusInternalServerError,
			Message: "Unspecified error",
		}
	}
}
