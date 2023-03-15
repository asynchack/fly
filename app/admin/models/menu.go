package models

import (
	"fly/common/models"
)

/* jason-comment
菜单id
	名称
	url的路径
	icon
	菜单类型
	菜单权限
	菜单关联的role的id
	是否可见visible
	是否是父级菜单框架
	breadcrumb面包屑
	是否被选中
	对应的vue组件名称

*/

type Menu struct {
	MenuId     int    `json:"menuId" gorm:"primary_key; auto_increment"`
	ParentId   int    `json:"parentdId" gorm:"-"`
	MenuName   string `json:"menuName" gorm:"size:128"`
	Title      string `json:"title" gorm:"size:128"`
	Icon       string `json:"icon" gorm:"size:128;"`
	Path       string `json:"path" gorm:"size:128;"`
	Paths      string `json:"paths" gorm:"size:128;"`
	MenuType   string `json:"menuType" gorm:"size:1;"`
	Action     string `json:"action" gorm:"size:16;"`
	Permission string `json:"permission" gorm:"size:255;"`
	NoCache    bool   `json:"noCache" gorm:"size:8;"`
	Breadcrumb string `json:"breadcrumb" gorm:"size:255;"`
	Component  string `json:"component" gorm:"size:255;"`
	Sort       int    `json:"sort" gorm:"size:4;"`
	Visible    string `json:"visible" gorm:"size:1;"`
	CreateBy   string `json:"createBy" gorm:"size:128;"`
	UpdateBy   string `json:"updateBy" gorm:"size:128;"`
	IsFrame    bool   `json:"isFrame" gorm:"size:1; default:0"`
	DataScope  string `json:"dataScope" gorm:"-"`
	Params     string `json:"params" gorm:"-"`
	RoleId     int    `gorm:"-"`
	Children   []Menu `json:"children" gorm:"-"`
	IsSelect   bool   `json:"is_select" gorm:"-"`
	models.BaseModel
}

func (Menu) TableName() string {
	return "sys_menu"
}
