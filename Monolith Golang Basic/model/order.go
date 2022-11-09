package model

type Orders struct {
	Id              int64  `db:"id"`
	GoodsName       string `db:"goods_name"`
	ReceiverName    string `db:"receiver_name"`
	ReceiverAddress string `db:"receiver_address"`
	ShipperID       int32  `db:"shipper_id"`
}
