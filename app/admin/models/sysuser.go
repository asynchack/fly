package models

import "fly/common/models"

/* jason-comment
为什么拆？
	应该是为了，方便组合，构成用户相关的，不同的数据结构
*/

type SysUserId struct {
	UserId int `json:"userId" gorm:"primary_key; auto_increment"`
}

type UserName struct {
	Username string `json:"username" gorm:"size:128"`
}

type PassWord struct {
	Password string `json:"password" gorm:"size:128"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUserB struct {
	// 这里才是用户一些基本资料信息
	/* jason-comment
	昵称 性别、邮箱、手机、关联的role的id、salt（密码加盐）、头像、部门id、职位id、备注、状态、等
	*/
	NickName string `gorm:"size:128" json:"nickName"` // 昵称
	Phone    string `gorm:"size:11" json:"phone"`     // 手机号
	RoleId   int    `gorm:"" json:"roleId"`           // 角色编码
	Salt     string `gorm:"size:255" json:"salt"`     //盐
	Avatar   string `gorm:"size:255" json:"avatar"`   //头像
	Sex      string `gorm:"size:255" json:"sex"`      //性别
	Email    string `gorm:"size:128" json:"email"`    //邮箱
	DeptId   int    `gorm:"" json:"deptId"`           //部门编码
	PostId   int    `gorm:"" json:"postId"`           //职位编码
	CreateBy string `gorm:"size:128" json:"createBy"` //
	UpdateBy string `gorm:"size:128" json:"updateBy"` //
	Remark   string `gorm:"size:255" json:"remark"`   //备注
	Status   string `gorm:"size:4;" json:"status"`
	models.BaseModel

	DataScope string `gorm:"-" json:"dataScope"`
	Params    string `gorm:"-" json:"params"`
}
type SysUser struct {
	// 是继承了3个结构体来的
	SysUserId
	LoginM
	SysUserB
}

func (SysUser) TableName() string {
	return "sys_user"
}
