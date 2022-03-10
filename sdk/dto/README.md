# dto
放置数据传输对象
即req/res 结构
允许引用sdk.model 的内容

## 数据格式规范
1. 不要使用uint/int,请使用int64/uint64 (为了配合其他语言,proto需要标明)
2. 