package version

import (
	"gorm.io/gorm"

	"fly/app/admin/models/system"
	"fly/cmd/migrate/migration"
	common "fly/common/models"

	"runtime"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.MigrateInstance.SetVersions(migration.ConvertFileName2Int(fileName), _1603516925109Migrate)
}

func _1603516925109Migrate(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if tx.Migrator().HasColumn(&system.SysLoginLog{}, "info_id") {
			err := tx.Migrator().RenameColumn(&system.SysLoginLog{}, "info_id", "id")
			if err != nil {
				return err
			}
		}

		if tx.Migrator().HasColumn(&system.SysOperaLog{}, "oper_id") {
			err := tx.Migrator().RenameColumn(&system.SysOperaLog{}, "oper_id", "id")
			if err != nil {
				return err
			}
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
