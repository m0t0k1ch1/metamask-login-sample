package response

import "github.com/m0t0k1ch1/metamask-login-sample/domain"

const (
	ResponseStateSuccess = "success"
	ResponseStateError   = "error"
)

type Response struct {
	State  string      `json:"state"`
	Result interface{} `json:"result"`
}

func newSuccessResponse(result interface{}) *Response {
	return &Response{
		State:  ResponseStateSuccess,
		Result: result,
	}
}

func newErrorResponse(err *domain.Error) *Response {
	return &Response{
		State:  ResponseStateError,
		Result: err,
	}
}
