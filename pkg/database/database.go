package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Create() (*sql.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/chatify?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		str := fmt.Errorf("failed to open database: %v", err)
		return nil, str
	}

	if err = db.Ping(); err != nil {
		err = fmt.Errorf("failed to ping database: %v", err)
		return nil, err
	}

	return db, nil
}
