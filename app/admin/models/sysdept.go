package models

import (
	"fly/common/models"
)

/* jason-comment
部门应该有那些信息？
	部门id，主键？
	部门名称
	上级部门
	部门领导
	领导电话
	领导email
	部门信息创建人
	部门信息修改人
	子部门列表

*/
type SysDept struct {
	DeptId   int       `json:"deptId" gorm:"primary_key; auto_increment"`
	ParentId int       `json:"parentId" gorm:"-"`
	DeptName string    `json:"depatName" gorm:"size: 128"`
	Leader   string    `json:"leader" gorm:"size: 128"`
	Phone    string    `json:"phone" gorm:"size:11"`
	Email    string    `json:"email" gorm:"size:128"`
	CreateBy string    `json:"createBy" gorm:"size: 128"`
	UpdateBy string    `json:"updateBy" gorm:"size:128"`
	Children []SysDept `json:"children" gorm:"-"`

	DeptPath string `json:"deptPath" gorm:"size:255;"` //
	Status   string `json:"status" gorm:"size:4;"`     //状态

	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params" gorm:"-"`
	models.BaseModel
}
