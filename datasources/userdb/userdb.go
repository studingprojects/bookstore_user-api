package userdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUserdbUsername = "MYSQL_USERDB_USER"
	mysqlUserdbPassword = "MYSQL_USERDB_PASS"
	mysqlUserdbHost     = "MYSQL_USERDB_HOST"
	mysqlUserdbPort     = "MYSQL_USERDB_PORT"
	mysqlUserdbSchema   = "MYSQL_USERDB_DB"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUserdbUsername)
	password = os.Getenv(mysqlUserdbPassword)
	host     = fmt.Sprintf("%s:%s", os.Getenv(mysqlUserdbHost), os.Getenv(mysqlUserdbPort))
	schema   = os.Getenv(mysqlUserdbSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		// panic(err)
	}
	if err = Client.Ping(); err != nil {
		// panic(err)
	}
	log.Print("database configured!")
}
