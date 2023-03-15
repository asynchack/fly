package models

import "time"

/* jason-comment
定义了Migration表的结构，和表名
该表负责记录所有migrate表的名称和apply的时间
*/
type Migration struct {
	Version   string    `gorm:"primary_key"`
	ApplyTime time.Time `gorm:"autoCreateTime"`
}

func (Migration) TableName() string {
	return "sys_migration"
}
