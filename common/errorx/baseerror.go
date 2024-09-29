package errorx

import (
	"go_zero_pgsql/common/globalkey"

	errors "github.com/zeromicro/x/errors"
)

// CodeMsg is a struct that contains a code and a message.
// It implements the error interface.
//type CodeMsg struct {
//	Code int
//	Msg  string
//}
//
//func (c *CodeMsg) Error() string {
//	return fmt.Sprintf("code: %d, msg: %s", c.Code, c.Msg)
//}
//
//// New creates a new CodeMsg.
//func New(code int, msg string) error {
//	return &CodeMsg{Code: code, Msg: msg}
//}
//

//
//type CodeError struct {
//	Code int    `json:"code"`
//	Msg  string `json:"msg"`
//}
//
//type CodeErrorResponse struct {
//	Code int    `json:"code"`
//	Msg  string `json:"msg"`
//}

func NewCodeError(code int, msg string) error {
	return &errors.CodeMsg{Code: code, Msg: msg}
}

func NewDefaultError(code int) error {
	return errors.New(code, MapErrMsg(code))
}

func NewHandlerError(code int, msg string) error {
	return errors.New(code, msg)
}

func NewSystemError(code int, msg string) error {
	if globalkey.SysShowSystemError {
		return errors.New(code, msg)
	} else {
		return errors.New(code, MapErrMsg(code))
	}
}

//func (e *CodeError) Error() string {
//	return e.Msg
//}

//func (e *CodeError) Data() *CodeErrorResponse {
//	return &CodeErrorResponse{
//		Code: e.Code,
//		Msg:  e.Msg,
//	}
//}
