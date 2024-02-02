package httpx

import (
	"encoding/json"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces/errors"
)

// GenericResponse is an extension of the BaseResponse which contains additional requestUUID tag
type GenericResponse struct {
	Success bool        `json:"success"`         // Success - true, false
	Data    interface{} `json:"data"`            // Optional
	Error   *BaseError  `json:"error,omitempty"` // Optional
	Next    string      `json:"next,omitempty"`  // Next Paginated Url
	Prev    string      `json:"prev,omitempty"`  // Prev Paginated Url
}

func (gen *GenericResponse) UnmarshalData(data interface{}) {
	str, _ := json.Marshal(gen.Data)
	err := json.Unmarshal(str, data)
	if err != nil {
		return
	}
	gen.Data = data
}

// BaseError base error wich will be shown to the client
type BaseError struct {
	Code         int          `json:"code"`    // One of a server-defined set of error codes.
	Message      string       `json:"message"` // A human-readable representation of the error.
	ErrorMessage string       `json:"errorMessage"`
	Details      []InnerError `json:"details"` // Optional, An array of Details about specific errors that led to this reported error.
}

// InnerError is an additional error that can be used in validation
type InnerError struct {
	Field  string            `json:"field"`
	Code   int               `json:"code"`
	Params map[string]string `json:"params,omitempty"`
}

func NewSuccess() GenericResponse {
	return GenericResponse{
		Success: true,
	}
}

func NewSuccessWithData(data interface{}) GenericResponse {
	return GenericResponse{
		Success: true,
		Data:    data,
	}
}

func NewError(msg errors.CodeMsg) GenericResponse {
	return GenericResponse{
		Error: &BaseError{
			Code:    msg.Code,
			Message: msg.Message,
		},
	}
}

func NewErrorWithData(msg errors.CodeMsg, data interface{}) GenericResponse {
	return GenericResponse{
		Data: data,
		Error: &BaseError{
			Code:    msg.Code,
			Message: msg.Message,
		},
	}
}

func NewErrorWithDetails(msg errors.CodeMsg, details []InnerError) GenericResponse {
	return GenericResponse{
		Error: &BaseError{
			Code:    msg.Code,
			Message: msg.Message,
			Details: details,
		},
	}
}

type Builder struct {
	GenericResponse
}

func NewResponseBuilder() *Builder {
	return &Builder{}
}

func (rb *Builder) IsSuccess() *Builder {
	rb.Success = true
	return rb
}

func (rb *Builder) SetData(data interface{}) *Builder {
	rb.Data = data
	return rb
}

func (rb *Builder) SetError(data interface{}) *Builder {
	rb.Data = data
	return rb
}

func (rb *Builder) Build() GenericResponse {
	return rb.GenericResponse
}

func (rb *Builder) BuildPointer() *GenericResponse {
	return &rb.GenericResponse
}
