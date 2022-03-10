package errbuilder

type BaseCode struct {
	BaseErrCode    int
	BaseErrMessage string
	Err            error
}

func (c *BaseCode) Error() string          { return c.Err.Error() }
func (c *BaseCode) Code() int              { return c.BaseErrCode }
func (c *BaseCode) Message() string        { return c.BaseErrMessage }
func (c *BaseCode) Details() []interface{} { return nil }
func (c *BaseCode) DetailErr() error       { return c.Err }

func (c *BaseCode) Cause() error {
	if c == nil {
		return nil
	}
	return c.DetailErr()
}
