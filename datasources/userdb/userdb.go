package userdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/federicoleon/bookstore_utils-go/rest_errors"
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
	UserDb userDbClientInterface = &userDbClient{}

	username = os.Getenv(mysqlUserdbUsername)
	password = os.Getenv(mysqlUserdbPassword)
	host     = fmt.Sprintf("%s:%s", os.Getenv(mysqlUserdbHost), os.Getenv(mysqlUserdbPort))
	schema   = os.Getenv(mysqlUserdbSchema)
)

type userDbClientInterface interface {
	setMySQLsClient(*sql.DB)
	Prepare(string) (*sql.Stmt, rest_errors.RestErr)
}

type userDbClient struct {
	client *sql.DB
}

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(); err != nil {
		panic(err)
	}
	UserDb.setMySQLsClient(client)
	log.Print("database configured!")
}

func (c *userDbClient) setMySQLsClient(client *sql.DB) {
	c.client = client
}

func (c *userDbClient) Prepare(query string) (*sql.Stmt, rest_errors.RestErr) {
	stmt, err := c.client.Prepare(query)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error(), nil)
	}
	return stmt, nil
}
