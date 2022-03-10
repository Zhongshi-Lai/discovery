# err
1. 不要在代码中使用err.Error() 根据字符串判断,做一些操作


## dao层的error
代码进行分层的含义,各层之前要进行足够的解耦
所以对dao层进行调用的时候,上层不需要关注是 mysql出现问题/redis出现问题/http/grpc 出现问题
所以上层只需要知道
1. 是否有error
2. error 是否是 找不到数据/其他error (找不到数据毕竟比较普通)
3. 具体是mysql的什么error(根因) 需要被记录下来,方便日志数据

基于以上几点,dao层具体的error(mysql/redis error)需要被吞掉,返回一个dao error

## pkg 的error
请使用
import 	"github.com/pkg/errors"
import 	"github.com/pkg/errors"
import 	"github.com/pkg/errors"
重要的事情说三遍

因为pkg的函数有可能被抽出来,放到依赖中
所以统一使用 
1. errors.New("xxx")

ps:大部分第三方库, 如gorm ,会使用标准库的errors.New() 这个是不带堆栈信息的

## service error
我们自己业务生成的异常

### go-paqu使用简述

#### kratos及grpc的要求
api层,使用krators的话, 整个api层是使用.proto文件生成的
api(pb生成) 直接调用的service func 返回给api的error,需要使用proto的&Status{} 结构
如果不这么返回的话,统一被认为是500,并且返回给前端的msg 很难看,非常不友好
详细在这段代码中  middleware logger
```go
c.Next()
err := c.Error
cerr := ecode.Cause(err)
```

```go
// Cause cause from error to ecode.
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}
```

但是如果初始化一个proto的&Status{} ,就会丢失dao层的error/调用其他service的error
比如go-paqu的写法
```go
if dErr := s.acDao.CheckIn(ctx, req.ActivityId, userId, req.Source); dErr != nil {
    return xec.ErrorHasFirstSign // ecode.Error
    
}
```

```go
func (d *Dao) CheckIn(ctx context.Context, activityId, userId int64, source string) error {
	req := &v1pb.CheckInReq{
		UserId:  userId,
		BizType: v1pb.BizType_ActivityType,
		BizId:   activityId,
		Source:  source,
	}
	_, err := d.ActivityCli.ActivityClient.CheckIn(ctx, req)
	if err != nil {
		log.Errorc(ctx, "[CheckIn]ActivityCli.ActivityClient.CheckIn call fail err:%v, req:%v", err, req)
	}

	return err
}
```

坏处
1. dErr 丢失了,返回的是 xec.ErrorHasFirstSign
2. 需要在 s.acDao.CheckIn 打印日志,造成每个dao都需要打印err

### swan项目的实现

