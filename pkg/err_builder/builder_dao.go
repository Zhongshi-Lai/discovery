package errbuilder

import (
	"discovery/sdk/constant"
	errcode "discovery/sdk/err_code"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DaoCode struct {
	Kind   constant.DaoErrKind    // 错误类型
	Err    error                  // 最底层详细错误 resty/xredis/grom
	Method constant.DataGetMethod // 获取数据的方式
}

func (c *DaoCode) Error() string                         { return c.Err.Error() }
func (c *DaoCode) Code() int                             { return int(errcode.DaoErrCode) }
func (c *DaoCode) Message() string                       { return "服务异常,请稍后重试" }
func (c *DaoCode) Details() []interface{}                { return nil }
func (c *DaoCode) DetailErr() error                      { return c.Err }
func (c *DaoCode) DataErrKind() constant.DaoErrKind      { return c.Kind }
func (c *DaoCode) DataGetMethod() constant.DataGetMethod { return c.Method }
func (c *DaoCode) Cause() error {
	if c == nil {
		return nil
	}
	return c.DetailErr()
}

func NewGORMDaoErr(err error, errKind constant.DaoErrKind) error {

	if err == nil {
		return nil
	}
	// 如果传入的errKind 以传入的为准
	var thisErrKind constant.DaoErrKind

	if errKind > 2 || errKind < 0 {
		// 数据错误
		return errors.New("传入的异常类型有误")
	} else if errKind != 0 {
		// 如果指定了错误类型,以您指定的为准
		thisErrKind = errKind
	} else if errKind == 0 && errors.Is(err, gorm.ErrRecordNotFound) {
		thisErrKind = constant.DaoErrKindNotFind
	} else {
		thisErrKind = constant.DaoErrKindOtherErr
	}

	return &DaoCode{
		Kind:   thisErrKind,
		Err:    errors.New(err.Error()),
		Method: constant.DataGetMethodGORM,
	}
}

func NewRedisDaoErr(err error, errKind constant.DaoErrKind) error {

	if err == nil {
		return nil
	}

	// 如果传入的errKind 以传入的为准
	var thisErrKind constant.DaoErrKind

	if errKind > 2 || errKind < 0 {
		// 数据错误
		return errors.New("传入的异常类型有误")
	} else if errKind != 0 {
		// 如果指定了错误类型,以您指定的为准
		thisErrKind = errKind
		// TODO(laizhongshi): 目前没有引入redis
		//} else if errKind == 0 && errors.Is(err, redis.ErrNil) {
		//	thisErrKind = constant.DaoErrKindNotFind
	} else {
		thisErrKind = constant.DaoErrKindOtherErr
	}

	return &DaoCode{
		Kind:   thisErrKind,
		Err:    errors.New(err.Error()),
		Method: constant.DataGetMethodXRedis,
	}
}

func NewGRPCDaoErr(err error, errKind constant.DaoErrKind) error {

	if err == nil {
		return nil
	}
	// 如果传入的errKind 以传入的为准
	var thisErrKind constant.DaoErrKind

	if errKind > 2 || errKind < 0 {
		// 数据错误
		return errors.New("传入的异常类型有误")
	}

	return &DaoCode{
		Kind:   thisErrKind,
		Err:    errors.New(err.Error()),
		Method: constant.DataGetMethodGRPC,
	}
}
