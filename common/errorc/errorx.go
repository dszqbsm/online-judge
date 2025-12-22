package errorc

import "fmt"

const (
	ErrorCodeInvalidParameter = "INVALID_PARAMETER"

	ErrorCodeInvalidAuthorization = "INVALID_AUTHORIZATION"

	ErrorCodeInternalSystemError = "INTERNAL_SERVER_ERROR"
)

type ErrorResponse struct {
	StatusCode int `json:"-"`

	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (c *ErrorResponse) Error() string {
	return fmt.Sprintf("error code: %s, error msg: %s", c.ErrorCode, c.ErrorMsg)
}

func New(statusCode int, errorCode string, errorMsg string) *ErrorResponse {
	return &ErrorResponse{statusCode, errorCode, errorMsg}
}
