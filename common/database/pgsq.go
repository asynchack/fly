package database

import (
	"fly/common/global"
	"fly/tools/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pgsql struct{}

func (p *Pgsql) Setup() {
	// 本质就是借助gorm进行数据库连接实例的创建
	db, err := p.Open(p.GetSource())
	if err != nil {
		fmt.Println(err.Error())
		panic("get global.DB instance failed!")

	}
	global.DB = db

}

func (p *Pgsql) Open(dsn string) (*gorm.DB, error) {
	// 创建数据库实例，赋值给global.DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err

}

func (p *Pgsql) GetDriver() string {
	// 返回Driver字符串 比如 “mysql”
	return config.DatabaseConfig.Driver

}

func (p *Pgsql) GetSource() string {
	// 返回连接字符串， 比如 fly@fly@tcp(127.0.0.1:3306)/fly?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
	// fmt.Println(config.DatabaseConfig.Source)

	return config.DatabaseConfig.Source
}
