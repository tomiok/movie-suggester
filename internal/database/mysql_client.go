package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tomiok/movies-suggester/internal/logs"
)

type MySqlClient struct {
	*sql.DB
}

func NewMySQLClient() *MySqlClient {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/movies")

	if err != nil {
		logs.Error("cannot create mysql client")
		panic(err)
	}

	err = db.Ping()

	if err != nil {

	}

	return &MySqlClient{db}
}
