package enum

type PointsExchangeType int64

const (
	_ PointsExchangeType = iota
	// PointsChangePopvip 积分商城
	PointsChangePopvip
	// PointsChangeHd 海鼎
	PointsChangeHd
)
