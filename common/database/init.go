package database

import (
	"fly/common/global"
	"fly/tools/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup(Driver string) error {

	if Driver == "mysql" {
		db := new(Mysql)
		db.Setup()
		return nil
	} else if Driver == "pgsql" {
		db := new(Pgsql)
		db.Setup()
		return nil
	} else {
		return fmt.Errorf("not support database type for : [%s]", Driver)
	}

}

type Mysql struct{}

func (m *Mysql) Setup() {
	// 本质就是借助gorm进行数据库连接实例的创建
	db, err := m.Open(m.GetSource())
	if err != nil {
		fmt.Println(err.Error())
		panic("get global.DB instance failed!")

	}
	global.DB = db

}

func (m *Mysql) Open(dsn string) (*gorm.DB, error) {
	// 创建数据库实例，赋值给global.DB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err

}

func (m *Mysql) GetDriver() string {
	// 返回Driver字符串 比如 “mysql”
	return config.DatabaseConfig.Driver

}

func (m *Mysql) GetSource() string {
	// 返回连接字符串， 比如 fly@fly@tcp(127.0.0.1:3306)/fly?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
	// fmt.Println(config.DatabaseConfig.Source)

	return config.DatabaseConfig.Source
}

type Pgsql struct{}

func (p *Pgsql) Setup() {

}
