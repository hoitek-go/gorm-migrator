package migrator

import "github.com/hoitek-go/gorm-migrator/share"

// Make all migrations Up.
func MigrateUpAll() {
	share.DB.Migrator().AutoMigrate(&share.Migration{})
	for _, mStruct := range AllMigrations {
		name := ModelName(mStruct)
		lastMigration := GetByName(name)
		if lastMigration == nil {
			mStruct.Up()
			share.DB.Create(&share.Migration{
				Name: name,
				Type: share.TYPE_MIGRATION,
			})
		}
	}
}
