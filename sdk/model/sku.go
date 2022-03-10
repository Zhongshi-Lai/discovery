package model

type SKU struct {
	BaseModel
	HDGoodsBarcode string `json:"hd_goods_barcode"` // 海鼎商品条码
	HDGoodsCode    string `json:"hd_goods_code"`    // 海鼎商品代码(确定商品)
	Price          int64  `json:"price"`            // sku售价 (单位:分)
	OriginalPrice  int64  `json:"original_price"`   // sku原价 (单位:分)
	Name           string `json:"name"`             // 商品sku名称
	MainImg        string `json:"main_img"`         // sku主图
	Status         int64  `json:"status"`           // sku状态
	GoodsType      int64  `json:"goods_type"`       // 商品类型
}

// TableName 设定表名
func (SKU) TableName() string {
	return "sku"
}
