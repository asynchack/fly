package models

import "fly/common/models"

type Post struct {
	PostId   int    `gorm:"primary_key;AUTO_INCREMENT" json:"postId"` //岗位编号
	PostName string `gorm:"size:128;" json:"postName"`                //岗位名称
	PostCode string `gorm:"size:128;" json:"postCode"`                //岗位代码
	Sort     int    `gorm:"" json:"sort"`                             //岗位排序
	Status   string `gorm:"size:4;" json:"status"`                    //状态
	Remark   string `gorm:"size:255;" json:"remark"`                  //描述
	CreateBy string `gorm:"size:128;" json:"createBy"`
	UpdateBy string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel

	DataScope string `gorm:"-" json:"dataScope"`
	Params    string `gorm:"-" json:"params"`
}

func (Post) TableName() string {
	return "sys_post"
}
