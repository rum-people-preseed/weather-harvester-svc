package errors

import "fmt"

type CodeMsg struct {
	HttpStatus int    `json:"-"`
	Code       int    `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (c *CodeMsg) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", c.Code, c.Message)
}

func (c *CodeMsg) CodeMsg() *CodeMsg {
	return c
}

// New creates a new CodeMsg.
func New(httpStatus int, code int, msg string) interface {
	Error() string
	CodeMsg() *CodeMsg
} {
	return &CodeMsg{HttpStatus: httpStatus, Code: code, Message: msg}
}
