package models

//sys_role_dept
type SysRoleDept struct {
	RoleId int `gorm:""`
	DeptId int `gorm:""`
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}
