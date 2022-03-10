package errbuilder

import (
	errcode "discovery/sdk/err_code"

	"github.com/pkg/errors"
)

type BizCode struct {
	BizErrCode    int
	BizErrMessage string
	Err           error
	BizErrContent string
}

func (c *BizCode) Error() string          { return c.Err.Error() }
func (c *BizCode) Code() int              { return c.BizErrCode }
func (c *BizCode) Message() string        { return c.BizErrMessage }
func (c *BizCode) Details() []interface{} { return nil }
func (c *BizCode) DetailErr() error       { return c.Err }
func (c *BizCode) BizContent() string     { return c.BizErrContent }
func (c *BizCode) Cause() error {
	if c == nil {
		return nil
	}
	return c.DetailErr()
}

func NewBizErr(err error, code errcode.ErrCode, message string, content string) error {

	// err != nil 意味着您需要将一个 !! 已经含有堆栈信息 !! 的err包装成业务错误
	// err == nil 意味着这是您的业务判断产生的异常

	if err == nil {
		err = errors.New(message)
	}

	return &BizCode{
		BizErrCode:    int(code),
		BizErrMessage: message,
		Err:           err,
		BizErrContent: content,
	}
}
