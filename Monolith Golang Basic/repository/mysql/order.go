package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang-basic/model"
	_interface "golang-basic/repository/interface"
)

type dbRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(dbConfig *sqlx.DB) _interface.OrderRepository{
	db := &dbRepo{
		db: dbConfig,
	}
	return db
}

func NewMock(dbConfig *sql.DB) *dbRepo {
	db := &dbRepo{
		db: sqlx.NewDb(dbConfig,"sqlmock"),
	}
	return db
}

func (repository *dbRepo) Insert(ctx context.Context, order model.Orders) (int64, error) {
	tx, err := repository.db.Begin()
	if err != nil {
		panic(err)
	}

	sql := `INSERT INTO orders(goods_name,receiver_name,receiver_address,shipper_id) VALUES(?,?,?,?)`
	result, err := tx.ExecContext(ctx, sql,
		order.GoodsName,
		order.ReceiverName,
		order.ReceiverAddress,
		order.ShipperID,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	fmt.Println("Last InsertId:", insertId)
	return insertId, nil
}

func (repository *dbRepo) UpdateOrderByID(ctx context.Context, order model.Orders) (model.Orders, error){
	tx, err := repository.db.Begin()
	if err != nil {
		panic(err)
	}

	sql := `UPDATE orders SET goods_name= ? , receiver_name = ? , receiver_address = ?, shipper_id = ? WHERE id= ?`
	_ , err = tx.ExecContext(ctx, sql,
		order.GoodsName,
		order.ReceiverName,
		order.ReceiverAddress,
		order.ShipperID,
		order.Id,
	)
	if err != nil {
		tx.Rollback()
		return model.Orders{}, err
	}

	tx.Commit()
	return order, nil
}

func (repository *dbRepo) GetOrdersByIDs(ctx context.Context, id []int64) ([]model.Orders, error) {
	sql := "SELECT id, goods_name,receiver_name,receiver_address,shipper_id FROM orders WHERE id IN (?)"
	query, args, err := sqlx.In(sql, id)
	if err != nil {
		return nil, err
	}

	var orders []model.Orders
	query = repository.db.Rebind(query)
	err = repository.db.SelectContext(ctx, &orders, query, args...)

	return orders, err
}
