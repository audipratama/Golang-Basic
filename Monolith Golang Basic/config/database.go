package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

func ConnectDB(conf *DBConfig) (*sqlx.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp("+conf.DB["golang-basic"].Host+")/golang_basic_sql")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	db.SetMaxIdleConns(conf.DB["golang-basic"].MaxIdleConn)
	db.SetMaxOpenConns(conf.DB["golang-basic"].MaxOpenConn)
	db.SetConnMaxLifetime(5 * time.Minute)
	fmt.Println("connection success")
	return sqlx.NewDb(db,"mysql"), nil
}
