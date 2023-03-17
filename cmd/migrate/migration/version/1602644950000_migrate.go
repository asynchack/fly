package version

import (
	"runtime"

	"gorm.io/gorm"

	"fly/app/admin/models"
	"fly/app/admin/models/system"
	"fly/cmd/migrate/migration"
	common "fly/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.MigrateInstance.SetVersions(migration.ConvertFileName2Int(fileName), _1602644950000Migrate)
}

func _1602644950000Migrate(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if tx.Migrator().HasColumn(&system.SysConfig{}, "config_id") {
			err := tx.Migrator().RenameColumn(&system.SysConfig{}, "config_id", "id")
			if err != nil {
				return err
			}
		}
		list2 := []models.CasbinRule{
			{PType: "p", V0: "admin", V1: "/api/v1/config", V2: "GET"},
		}
		err := tx.Create(list2).Error
		if err != nil {
			return err
		}

		menu := models.Menu{MenuId: 86, Path: "/api/v1/config"}
		err = tx.Model(&menu).Where("menu_id = ?", 86).Update("Path", "/api/v1/config").Error
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
