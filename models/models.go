package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init()  {
	var (
		err error
		dbType, dbName ,user, password, host, tablePrefix string
	)
	dbType = viper.GetString("TYPE")
	dbName = viper.GetString("NAME")
	user   = viper.GetString("USER")
	password = viper.GetString("PASSWORD")
	host =  viper.GetString("HOST")
	tablePrefix = viper.GetString("TABLE_PREFIX")

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}
	// 默认的表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxOpenConns(viper.GetInt("MAX_OPEN_CONNS"))
	db.DB().SetMaxIdleConns(viper.GetInt("MAX_IDLE_CONNS"))
}

// 关闭数据库连接
func CloseDB()  {
	defer db.Close()
}

