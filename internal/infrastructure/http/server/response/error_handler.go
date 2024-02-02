package response

import (
	stderrors "errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces/errors"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces/httpx"
)

// HttpError basic http error
func HttpError(httpStatus int) error {
	return HttpWithServiceError(httpStatus, *ApiErrorUnspecified.CodeMsg())
}

// HttpWithServiceError Make an echo compatible error out of a service error
func HttpWithServiceError(httpStatus int, code errors.CodeMsg) error {
	return HttpErrorWithInternal(httpStatus, code, nil)
}

// HttpErrorValidation Validation errror with internal error
func HttpErrorValidation(internalErr error) error {
	he := echo.NewHTTPError(http.StatusUnprocessableEntity, ApiErrorUnprocessableEntity)
	if internalErr != nil {
		he.Internal = internalErr
	}
	return he
}

func HttpErrorBodyValidation(internalErr error) error {
	he := echo.NewHTTPError(http.StatusBadRequest, ApiErrorUnprocessableEntity)
	if internalErr != nil {
		he.Internal = internalErr
	}
	return he
}

// HttpErrorValidationWithCode Validation error with internal error
func HttpErrorValidationWithCode(code errors.CodeMsg, internalErr error) error {
	he := echo.NewHTTPError(http.StatusUnprocessableEntity, code)
	if internalErr != nil {
		he.Internal = internalErr
	}
	return he
}

// HttpErrorWithInternal echo.HttpError has a message interface{} field, which can be used to include the ServiceCode
func HttpErrorWithInternal(httpStatus int, code interface{}, internalErr error) error {
	he := echo.NewHTTPError(httpStatus, code)
	if internalErr != nil {
		he.Internal = internalErr
	}
	return he
}

func doResponse(httpStatus int, apiError errors.CodeMsg, c echo.Context, internal error) {
	if c.Response().Committed {
		return
	}
	if c.Request().Method == http.MethodHead {
		_ = c.NoContent(httpStatus)
		return
	}

	var (
		internalErrorMsg string
		details          []httpx.InnerError
		validationErrors validator.ValidationErrors
	)
	if stderrors.As(internal, &validationErrors) {
		details = make([]httpx.InnerError, len(validationErrors))
		for k, err := range validationErrors {
			fieldErr := ParseFieldError(err)

			details[k] = httpx.InnerError{
				Field:  fieldErr.Field,
				Code:   fieldErr.Code,
				Params: fieldErr.Params,
			}
		}
	}

	if internal != nil {
		internalErrorMsg = internal.Error()
		c.Logger().Errorj(map[string]interface{}{
			"id":  c.Response().Header().Get(echo.HeaderXRequestID),
			"msg": internal,
		})
	}

	_ = c.JSON(httpStatus, httpx.GenericResponse{
		Error: &httpx.BaseError{
			Code:         apiError.Code,
			Message:      apiError.Message,
			ErrorMessage: internalErrorMsg,
			Details:      details,
		},
	})
}

func HttpErrorHandler(err error, c echo.Context) {
	var (
		httpErr *echo.HTTPError
		codeMsg *errors.CodeMsg
	)
	if ok := stderrors.As(err, &httpErr); ok {
		if codeMsg, ok = httpErr.Message.(*errors.CodeMsg); !ok && !stderrors.As(httpErr.Internal, &codeMsg) {
			v := ApiErrorMapping(httpErr)
			codeMsg = &v
		}
		doResponse(httpErr.Code, *codeMsg, c, httpErr.Internal)
		return
	}

	if stderrors.As(err, &codeMsg) {
		doResponse(codeMsg.HttpStatus, *codeMsg, c, nil)
		return
	}
	doResponse(http.StatusInternalServerError, *ApiErrorInternalServerError.CodeMsg(), c, err)
}
