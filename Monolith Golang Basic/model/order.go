package model

type Orders struct {
	Id int64 `db:"id"`
	GoodsName string `db:"goodsName"`
}
