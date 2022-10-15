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

func (repository *dbRepo) Insert(ctx context.Context, order model.Orders) (int64, error) {
	tx, err := repository.db.Begin()
	if err != nil {
		panic(err)
	}

	sql := "INSERT INTO orders(goodsName) VALUES(?)"
	result, err := tx.ExecContext(ctx, sql,
		order.GoodsName)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
	fmt.Println("Last InsertId:", insertId)
	return insertId, nil
}

func (repository *dbRepo) GetOrdersByIDs(ctx context.Context, id []int64) ([]model.Orders, error) {
	sql := "SELECT id, goodsName FROM orders WHERE id IN (?)"
	query, args, err := sqlx.In(sql, id)
	if err != nil {
		return nil, err
	}

	var orders []model.Orders
	query = repository.db.Rebind(query)
	err = repository.db.SelectContext(ctx, &orders, query, args...)

	return orders, err
}

func GetOrderById(db *sql.DB, id int) {
	ctx := context.Background()

	sql := "SELECT id, goodsName FROM orders WHERE id = ?"
	rows, err := db.QueryContext(ctx, sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, goodsName string
		err = rows.Scan(&id, &goodsName)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id " + id)
		fmt.Println("GoodsName " + goodsName)
	}

	fmt.Println("Success")
}
