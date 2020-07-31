package connection

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlConfig struct {
	Host          string
	Port          int
	User          string
	Pass          string
	DbName        string
	MaxConnection int
}

func (c MysqlConfig) getConnectionString() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User,
		c.Pass,
		c.Host,
		c.DbName,
	)
}

func InitMysqlDB(conf MysqlConfig) *gorm.DB {
	db, err := gorm.Open("mysql", conf.getConnectionString())
	defer db.Close()
	if err != nil {
		panic(err)
	}
	// db.DB().SetMaxOpenConns(100)
	return db
}
