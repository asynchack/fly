package version

import (
	"fly/app/admin/models"
	"fly/app/admin/models/system"
	"fly/app/admin/models/tools"
	"fly/cmd/migrate/migration"
	commonModels "fly/common/models"
	"runtime"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.MigrateInstance.SetVersions(migration.ConvertFileName2Int(fileName), _1599190683659Tables)
}

func _1599190683659Tables(db *gorm.DB, version string) error {
	// 这里是定义了，实际表迁移的操作
	// 1、先迁移
	// 2、迁移后，写入一条对应记录到sys_migration表中

	err := db.Debug().Migrator().AutoMigrate(
		new(models.CasbinRule),
		new(models.SysDept),
		new(system.SysConfig),
		new(tools.SysTables),
		new(tools.SysColumns),
		new(models.Menu),
		new(system.SysLoginLog),
		new(system.SysOperaLog),
		new(models.RoleMenu),
		new(models.SysRoleDept),
		new(models.SysUser),
		new(models.SysRole),
		new(models.Post),
		new(models.DictData),
		new(models.DictType),
		new(models.SysSetting),
	)
	if err != nil {
		return err
	}
	return db.Create(&commonModels.Migration{
		Version: version,
		// ApplyTime: time.Now(), 时间自动创建的
	}).Error
}

// --------------------------------------------
// package version

// import (
// 	"runtime"

// 	"fiy/app/admin/models/system"

// 	"gorm.io/gorm"

// 	"fiy/app/admin/models"
// 	"fiy/app/admin/models/tools"
// 	"fiy/cmd/migrate/migration"
// 	common "fiy/common/models"
// )

// func init() {
// 	_, fileName, _, _ := runtime.Caller(0)
// 	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1599190683659Tables)
// }

// func _1599190683659Tables(db *gorm.DB, version string) error {
// 	err := db.Debug().Migrator().AutoMigrate(
// 		new(models.CasbinRule),
// 		new(models.SysDept),
// 		new(system.SysConfig),
// 		new(tools.SysTables),
// 		new(tools.SysColumns),
// 		new(models.Menu),
// 		new(system.SysLoginLog),
// 		new(system.SysOperaLog),
// 		new(models.RoleMenu),
// 		new(models.SysRoleDept),
// 		new(models.SysUser),
// 		new(models.SysRole),
// 		new(models.Post),
// 		new(models.DictData),
// 		new(models.DictType),
// 		new(models.SysSetting),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	return db.Create(&common.Migration{
// 		Version: version,
// 	}).Error
// }
