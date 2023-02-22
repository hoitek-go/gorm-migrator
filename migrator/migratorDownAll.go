package migrator

import "github.com/hoitek-go/gorm-migrator/share"

// Make down all migrations.
func MigrateDownAll() {
	share.DB.Migrator().AutoMigrate(&share.Migration{})
	for _, mStruct := range AllMigrations {
		name := ModelName(mStruct)
		lastMigration := GetByName(name)
		if lastMigration != nil {
			mStruct.Down()
			share.DB.Unscoped().Where("name = ? and type = ?", name, share.TYPE_MIGRATION).Delete(&share.Migration{})
		}
		return
	}
}
