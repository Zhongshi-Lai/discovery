package points

type ChangeReq struct {
	OrderID string `json:"order_id"`
	Occur   int64  `json:"occur"`
}
