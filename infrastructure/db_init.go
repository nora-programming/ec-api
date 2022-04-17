package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbConf struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *gorm.DB {
	return initDB(&dbConf{
		Host:     "localhost",
		Username: "root",
		Password: "password",
		DBName:   "ec_api",
	})
}

func initDB(c *dbConf) *gorm.DB {
	dsn := c.Username + ":" + c.Password + "@tcp(" + c.Host + ")/" + c.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
