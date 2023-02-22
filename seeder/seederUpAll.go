package seeder

import "github.com/hoitek-go/gorm-migrator/share"

// Make all seeds Up.
func SeedUpAll() {
	share.DB.Migrator().AutoMigrate(&share.Migration{})
	for _, mStruct := range AllSeeds {
		name := ModelName(mStruct)
		lastMigration := GetByName(name)
		if lastMigration == nil {
			mStruct.Up()
			share.DB.Create(&share.Migration{
				Name: name,
				Type: share.TYPE_SEED,
			})
		}
	}
}
