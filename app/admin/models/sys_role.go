package models

import "fly/common/models"

/* jason-comment
id、name、关联的部门列表、关联的菜单列表
*/

type SysRole struct {
	RoleId    int    `json:"roleId" gorm:"primary_key;AUTO_INCREMENT"` // 角色编码
	RoleName  string `json:"roleName" gorm:"size:128;"`                // 角色名称
	Status    string `json:"status" gorm:"size:4;"`                    //
	RoleKey   string `json:"roleKey" gorm:"size:128;"`                 //角色代码
	RoleSort  int    `json:"roleSort" gorm:""`                         //角色排序
	Flag      string `json:"flag" gorm:"size:128;"`                    //
	CreateBy  string `json:"createBy" gorm:"size:128;"`                //
	UpdateBy  string `json:"updateBy" gorm:"size:128;"`                //
	Remark    string `json:"remark" gorm:"size:255;"`                  //备注
	Admin     bool   `json:"admin" gorm:"size:4;"`
	DataScope string `json:"dataScope" gorm:"size:128;"`
	models.BaseModel

	Params  string `json:"params" gorm:"-"`
	MenuIds []int  `json:"menuIds" gorm:"-"`
	DeptIds []int  `json:"deptIds" gorm:"-"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
