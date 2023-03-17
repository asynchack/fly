package version

import (
	"fly/app/admin/models"
	"fly/app/admin/models/system"
	"fly/cmd/migrate/migration"
	common "fly/common/models"
	"runtime"
	"time"

	"gorm.io/gorm"
)

func init() {

	_, fileName, _, _ := runtime.Caller(0)
	migration.MigrateInstance.SetVersions(migration.ConvertFileName2Int(fileName), _1599190683670Migrate)
}

func _1599190683670Migrate(db *gorm.DB, version string) error {
	/* jason-comment
	拿到db实例后， 开启一个事务，在事务中进行表插入 ，确保数据一致性

	如果插入成功，则记录一条数据到sys_migration表中
	*/

	return db.Transaction(func(tx *gorm.DB) error {
		// 应该时：事务中，会调用Transaction中传入的函数，根据函数的返回的error判断，如果为nil，则commit，不是则rollback，所以函数内部，处理时，每一步都要判断err是否为nil，为nil，才能进行下一步，知道最后返回nil，说明事务成功可以commit
		// 否则，就rollback
		var listSysDept = []models.SysDept{
			{DeptId: 1, ParentId: 0, DeptPath: "0/1", DeptName: "王哥科技", Sort: 0, Leader: "aituo", Phone: "13782218188", Email: "goolesre@gmail.com", Status: "0", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},

			{DeptId: 7, ParentId: 1, DeptPath: "/0/1/7", DeptName: "研发部", Sort: 1, Leader: "aituo", Phone: "13782218188", Email: "fdevops@fdevops.com", Status: "0", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DeptId: 8, ParentId: 1, DeptPath: "/0/1/8", DeptName: "运维部", Sort: 0, Leader: "aituo", Phone: "13782218188", Email: "fdevops@fdevops.com", Status: "0", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DeptId: 9, ParentId: 1, DeptPath: "/0/1/9", DeptName: "客服部", Sort: 0, Leader: "aituo", Phone: "13782218188", Email: "fdevops@fdevops.com", Status: "0", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DeptId: 10, ParentId: 1, DeptPath: "/0/1/10", DeptName: "人力资源", Sort: 3, Leader: "aituo", Phone: "13782218188", Email: "fdevops@fdevops.com", Status: "1", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}
		err := tx.Create(&listSysDept).Error
		if err != nil {
			return err
		}
		var listSysConfig = []system.SysConfig{
			{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}}, ControlBy: common.ControlBy{CreateBy: 1, UpdateBy: 1}, ConfigName: "主页框架-默认皮肤样式名称", ConfigKey: "sys_index_skinName", ConfigValue: "skin-blue", ConfigType: "Y", Remark: "蓝色 skin blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow"},
			{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}}, ControlBy: common.ControlBy{CreateBy: 1, UpdateBy: 1}, ConfigName: "用户管理-账号初始密码", ConfigKey: "sys.user.initPassword", ConfigValue: "123456", ConfigType: "Y", Remark: "初始化密码 123456"},
			{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}}, ControlBy: common.ControlBy{CreateBy: 1, UpdateBy: 1}, ConfigName: "主框架页-侧边栏主题", ConfigKey: "sys_index_sideTheme", ConfigValue: "theme-dark", ConfigType: "Y", Remark: "深色主题theme-dark，浅色主题theme-light"},
		}

		err = tx.Create(listSysConfig).Error
		if err != nil {
			return err
		}

		var listPost = []models.Post{
			{PostId: 1, PostName: "首席执行官", PostCode: "CEO", Sort: 0, Status: "0", Remark: "首席执行官", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{PostId: 2, PostName: "首席技术执行官", PostCode: "CTO", Sort: 2, Status: "0", Remark: "首席技术执行官", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{PostId: 3, PostName: "首席运营官", PostCode: "COO", Sort: 3, Status: "0", Remark: "测试工程师", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}
		err = tx.Create(listPost).Error
		if err != nil {
			return err
		}

		var listSysRole = []models.SysRole{
			{1, "系统管理员", "0", "admin", 1, "", "1", "", "", true, "", common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, "", []int{}, []int{}},
		}
		err = tx.Create(listSysRole).Error
		if err != nil {
			return err
		}

		var listDictType = []models.DictType{
			{DictId: 1, DictName: "系统开关", DictType: "sys_normal_disable", Status: "0", CreateBy: "1", UpdateBy: "1", Remark: "系统开关列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictId: 2, DictName: "用户性别", DictType: "sys_user_sex", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "用户性别列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictId: 3, DictName: "菜单状态", DictType: "sys_show_hide", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "菜单状态列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictId: 4, DictName: "系统是否", DictType: "sys_yes_no", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "系统是否列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictId: 5, DictName: "通知类型", DictType: "sys_notice_type", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "通知类型列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictId: 6, DictName: "系统状态", DictType: "sys_common_status", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "登录状态列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictId: 7, DictName: "操作类型", DictType: "sys_oper_type", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "操作类型列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictId: 8, DictName: "通知状态", DictType: "sys_notice_status", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "通知状态列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}
		err = tx.Create(listDictType).Error
		if err != nil {
			return err
		}

		var listSysUser = []models.SysUser{
			{models.SysUserId{1}, models.LoginM{models.UserName{"admin"}, models.PassWord{"$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2"}}, models.SysUserB{"lanyulei", "13818888888", 1, "", "", "0", "1@qq.com", 1, 1, "1", "1", "", "0", common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, "", ""}},
		}
		err = tx.Create(listSysUser).Error
		if err != nil {
			return nil
		}

		var listDictData = []models.DictData{
			{DictCode: 1, DictSort: 0, DictLabel: "正常", DictValue: "0", DictType: "sys_normal_disable", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "系统正常", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 2, DictSort: 0, DictLabel: "停用", DictValue: "1", DictType: "sys_normal_disable", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "系统停用", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 3, DictSort: 0, DictLabel: "男", DictValue: "0", DictType: "sys_user_sex", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "性别男", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 4, DictSort: 0, DictLabel: "女", DictValue: "1", DictType: "sys_user_sex", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "性别女", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 5, DictSort: 0, DictLabel: "未知", DictValue: "2", DictType: "sys_user_sex", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "性别未知", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 6, DictSort: 0, DictLabel: "显示", DictValue: "0", DictType: "sys_show_hide", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "显示菜单", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 7, DictSort: 0, DictLabel: "隐藏", DictValue: "1", DictType: "sys_show_hide", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "隐藏菜单", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 8, DictSort: 0, DictLabel: "是", DictValue: "Y", DictType: "sys_yes_no", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "系统默认是", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 9, DictSort: 0, DictLabel: "否", DictValue: "N", DictType: "sys_yes_no", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "系统默认否", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 14, DictSort: 0, DictLabel: "通知", DictValue: "1", DictType: "sys_notice_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "通知", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 15, DictSort: 0, DictLabel: "公告", DictValue: "2", DictType: "sys_notice_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "公告", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 16, DictSort: 0, DictLabel: "正常", DictValue: "0", DictType: "sys_common_status", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "正常状态", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 17, DictSort: 0, DictLabel: "关闭", DictValue: "1", DictType: "sys_common_status", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "关闭状态", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 18, DictSort: 0, DictLabel: "新增", DictValue: "1", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "新增操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 19, DictSort: 0, DictLabel: "修改", DictValue: "2", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "修改操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 20, DictSort: 0, DictLabel: "删除", DictValue: "3", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "删除操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 21, DictSort: 0, DictLabel: "授权", DictValue: "4", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "授权操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 22, DictSort: 0, DictLabel: "导出", DictValue: "5", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "导出操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 23, DictSort: 0, DictLabel: "导入", DictValue: "6", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "导入操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 24, DictSort: 0, DictLabel: "强退", DictValue: "7", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "强退操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 25, DictSort: 0, DictLabel: "生成代码", DictValue: "8", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "生成操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 26, DictSort: 0, DictLabel: "清空数据", DictValue: "9", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "清空操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 27, DictSort: 0, DictLabel: "成功", DictValue: "0", DictType: "sys_notice_status", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "成功状态", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 28, DictSort: 0, DictLabel: "失败", DictValue: "1", DictType: "sys_notice_status", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "失败状态", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 29, DictSort: 0, DictLabel: "登录", DictValue: "10", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "1", Remark: "登录操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 30, DictSort: 0, DictLabel: "退出", DictValue: "11", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "1", Remark: "", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 31, DictSort: 0, DictLabel: "获取验证码", DictValue: "12", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "1", Remark: "获取验证码", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}
		err = tx.Create(listDictData).Error
		if err != nil {
			return nil
		}

		var listSysSetting = []models.SysSetting{
			{SettingsId: 1, Name: "fiy管理系统", Logo: "", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}
		err = tx.Create(listSysSetting).Error
		if err != nil {
			return nil
		}

		return tx.Create(&common.Migration{
			Version: version,
			// ApplyTime: time.Now(), 自动生成无需指定
		}).Error
	})

}

// package version

// import (
// 	"fiy/common/global"
// 	"runtime"
// 	"time"

// 	"gorm.io/gorm"

// 	"fiy/app/admin/models"
// 	"fiy/app/admin/models/system"
// 	"fiy/cmd/migrate/migration"
// 	common "fiy/common/models"
// )

// func init() {
// 	_, fileName, _, _ := runtime.Caller(0)
// 	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1599190683670Migrate)
// }

// func _1599190683670Migrate(db *gorm.DB, version string) error {
// 	return db.Transaction(func(tx *gorm.DB) error {

// 		list3 := []models.SysDept{
// 			{DeptId: 1, ParentId: 0, DeptPath: "/0/1", DeptName: "磊哥科技", Sort: 0, Leader: "aituo", Phone: "13782218188", Email: "fdevops@fdevops.com", Status: "0", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DeptId: 7, ParentId: 1, DeptPath: "/0/1/7", DeptName: "研发部", Sort: 1, Leader: "aituo", Phone: "13782218188", Email: "fdevops@fdevops.com", Status: "0", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DeptId: 8, ParentId: 1, DeptPath: "/0/1/8", DeptName: "运维部", Sort: 0, Leader: "aituo", Phone: "13782218188", Email: "fdevops@fdevops.com", Status: "0", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DeptId: 9, ParentId: 1, DeptPath: "/0/1/9", DeptName: "客服部", Sort: 0, Leader: "aituo", Phone: "13782218188", Email: "fdevops@fdevops.com", Status: "0", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DeptId: 10, ParentId: 1, DeptPath: "/0/1/10", DeptName: "人力资源", Sort: 3, Leader: "aituo", Phone: "13782218188", Email: "fdevops@fdevops.com", Status: "1", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 		}

// 		list4 := []system.SysConfig{
// 			{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}}, ControlBy: common.ControlBy{CreateBy: 1, UpdateBy: 1}, ConfigName: "主框架页-默认皮肤样式名称", ConfigKey: "sys_index_skinName", ConfigValue: "skin-blue", ConfigType: "Y", Remark: "蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow"},
// 			{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}}, ControlBy: common.ControlBy{CreateBy: 1, UpdateBy: 1}, ConfigName: "用户管理-账号初始密码", ConfigKey: "sys.user.initPassword", ConfigValue: "123456", ConfigType: "Y", Remark: "初始化密码 123456"},
// 			{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}}, ControlBy: common.ControlBy{CreateBy: 1, UpdateBy: 1}, ConfigName: "主框架页-侧边栏主题", ConfigKey: "sys_index_sideTheme", ConfigValue: "theme-dark", ConfigType: "Y", Remark: "深色主题theme-dark，浅色主题theme-light"},
// 		}

// 		list5 := []models.Post{
// 			{PostId: 1, PostName: "首席执行官", PostCode: "CEO", Sort: 0, Status: "0", Remark: "首席执行官", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{PostId: 2, PostName: "首席技术执行官", PostCode: "CTO", Sort: 2, Status: "0", Remark: "首席技术执行官", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{PostId: 3, PostName: "首席运营官", PostCode: "COO", Sort: 3, Status: "0", Remark: "测试工程师", CreateBy: "1", UpdateBy: "1", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 		}

// 		list6 := []models.SysRole{
// 			{1, "系统管理员", "0", "admin", 1, "", "1", "", "", true, "", common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, "", []int{}, []int{}},
// 		}

// 		list7 := []models.DictType{
// 			{DictId: 1, DictName: "系统开关", DictType: "sys_normal_disable", Status: "0", CreateBy: "1", UpdateBy: "1", Remark: "系统开关列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictId: 2, DictName: "用户性别", DictType: "sys_user_sex", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "用户性别列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictId: 3, DictName: "菜单状态", DictType: "sys_show_hide", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "菜单状态列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictId: 4, DictName: "系统是否", DictType: "sys_yes_no", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "系统是否列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictId: 5, DictName: "通知类型", DictType: "sys_notice_type", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "通知类型列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictId: 6, DictName: "系统状态", DictType: "sys_common_status", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "登录状态列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictId: 7, DictName: "操作类型", DictType: "sys_oper_type", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "操作类型列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictId: 8, DictName: "通知状态", DictType: "sys_notice_status", Status: "0", CreateBy: "1", UpdateBy: "", Remark: "通知状态列表", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 		}

// 		list8 := []models.SysUser{
// 			{models.SysUserId{1}, models.LoginM{models.UserName{"admin"}, models.PassWord{"$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2"}}, models.SysUserB{"lanyulei", "13818888888", 1, "", "", "0", "1@qq.com", 1, 1, "1", "1", "", "0", common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, "", ""}},
// 		}

// 		list9 := []models.DictData{
// 			{DictCode: 1, DictSort: 0, DictLabel: "正常", DictValue: "0", DictType: "sys_normal_disable", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "系统正常", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 2, DictSort: 0, DictLabel: "停用", DictValue: "1", DictType: "sys_normal_disable", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "系统停用", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 3, DictSort: 0, DictLabel: "男", DictValue: "0", DictType: "sys_user_sex", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "性别男", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 4, DictSort: 0, DictLabel: "女", DictValue: "1", DictType: "sys_user_sex", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "性别女", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 5, DictSort: 0, DictLabel: "未知", DictValue: "2", DictType: "sys_user_sex", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "性别未知", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 6, DictSort: 0, DictLabel: "显示", DictValue: "0", DictType: "sys_show_hide", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "显示菜单", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 7, DictSort: 0, DictLabel: "隐藏", DictValue: "1", DictType: "sys_show_hide", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "隐藏菜单", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 8, DictSort: 0, DictLabel: "是", DictValue: "Y", DictType: "sys_yes_no", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "系统默认是", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 9, DictSort: 0, DictLabel: "否", DictValue: "N", DictType: "sys_yes_no", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "系统默认否", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 14, DictSort: 0, DictLabel: "通知", DictValue: "1", DictType: "sys_notice_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "通知", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 15, DictSort: 0, DictLabel: "公告", DictValue: "2", DictType: "sys_notice_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "公告", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 16, DictSort: 0, DictLabel: "正常", DictValue: "0", DictType: "sys_common_status", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "正常状态", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 17, DictSort: 0, DictLabel: "关闭", DictValue: "1", DictType: "sys_common_status", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "关闭状态", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 18, DictSort: 0, DictLabel: "新增", DictValue: "1", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "新增操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 19, DictSort: 0, DictLabel: "修改", DictValue: "2", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "修改操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 20, DictSort: 0, DictLabel: "删除", DictValue: "3", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "删除操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 21, DictSort: 0, DictLabel: "授权", DictValue: "4", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "授权操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 22, DictSort: 0, DictLabel: "导出", DictValue: "5", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "导出操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 23, DictSort: 0, DictLabel: "导入", DictValue: "6", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "导入操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 24, DictSort: 0, DictLabel: "强退", DictValue: "7", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "强退操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 25, DictSort: 0, DictLabel: "生成代码", DictValue: "8", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "生成操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 26, DictSort: 0, DictLabel: "清空数据", DictValue: "9", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "清空操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 27, DictSort: 0, DictLabel: "成功", DictValue: "0", DictType: "sys_notice_status", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "成功状态", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 28, DictSort: 0, DictLabel: "失败", DictValue: "1", DictType: "sys_notice_status", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "", Remark: "失败状态", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 29, DictSort: 0, DictLabel: "登录", DictValue: "10", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "1", Remark: "登录操作", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 30, DictSort: 0, DictLabel: "退出", DictValue: "11", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "1", Remark: "", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 			{DictCode: 31, DictSort: 0, DictLabel: "获取验证码", DictValue: "12", DictType: "sys_oper_type", CssClass: "", ListClass: "", IsDefault: "", Status: "0", Default: "", CreateBy: "1", UpdateBy: "1", Remark: "获取验证码", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 		}

// 		list10 := []models.SysSetting{
// 			{1, "fiy管理系统", "", common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
// 		}

// 		err := tx.Create(list3).Error
// 		if err != nil {
// 			return err
// 		}

// 		err = tx.Create(list4).Error
// 		if err != nil {
// 			return err
// 		}

// 		err = tx.Create(list5).Error
// 		if err != nil {
// 			return err
// 		}

// 		err = tx.Create(list6).Error
// 		if err != nil {
// 			return err
// 		}

// 		err = tx.Create(list7).Error
// 		if err != nil {
// 			return err
// 		}

// 		err = tx.Create(list8).Error
// 		if err != nil {
// 			return err
// 		}

// 		err = tx.Create(list9).Error
// 		if err != nil {
// 			return err
// 		}

// 		err = tx.Create(list10).Error
// 		if err != nil {
// 			return err
// 		}

// 		if err = models.InitDb(tx, "config/sql/db.sql"); err != nil {
// 			global.Logger.Errorf("同步菜单数据失败, %v", err)
// 		}

// 		return tx.Create(&common.Migration{
// 			Version: version,
// 		}).Error
// 	})
// }
