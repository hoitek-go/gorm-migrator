package migrator

import (
	"github.com/hoitek-go/gorm-migrator/share"
)

// Make up a desired migration based on the given name.
func MigrateUp(migrationName string) {
	share.DB.Migrator().AutoMigrate(&share.Migration{})
	for _, mStruct := range AllMigrations {
		name := ModelName(mStruct)
		if name == migrationName {
			lastMigration := GetByName(name)
			if lastMigration == nil {
				mStruct.Up()
				share.DB.Create(&share.Migration{
					Name: name,
					Type: share.TYPE_MIGRATION,
				})
			}
			break
		}
	}
}
