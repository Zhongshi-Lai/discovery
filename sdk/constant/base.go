package constant

// DaoErrKind dao层的异常类型枚举
type DaoErrKind int

const (
	DaoErrKindDefault  DaoErrKind = 0
	DaoErrKindNotFind  DaoErrKind = 1 // 数据不存在
	DaoErrKindOtherErr DaoErrKind = 2 // 其他异常
)

// DataGetMethod dao层数据获取方式
type DataGetMethod int

const (
	DataGetMethodGORM   DataGetMethod = 1
	DataGetMethodXRedis DataGetMethod = 2
	DataGetMethodGRPC   DataGetMethod = 3
)

// BaseError 本项目基础error
type BaseError interface {
	// 兼容grpc
	Error() string
	Code() int
	Message() string
	Details() []interface{}

	Cause() error     // 追溯根因使用
	DetailErr() error // 具体的错误,根因放置于此
}

// BizError 本项目service层错误
type BizError interface {
	BaseError

	BizContent() string // 业务的错误的内容
}

// DaoError dao层error的
type DaoError interface {
	BaseError

	DataErrKind() DaoErrKind      // 错误类型
	DataGetMethod() DataGetMethod // 获取数据的方式
}
