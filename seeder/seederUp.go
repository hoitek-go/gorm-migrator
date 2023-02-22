package seeder

import "github.com/hoitek-go/gorm-migrator/share"

// Make up a desired seeds based on the given name.
func SeedUp(seedName string) {
	share.DB.Migrator().AutoMigrate(&share.Migration{})
	for _, mStruct := range AllSeeds {
		name := ModelName(mStruct)
		if name == seedName {
			lastMigration := GetByName(name)
			if lastMigration == nil {
				mStruct.Up()
				share.DB.Create(&share.Migration{
					Name: name,
					Type: share.TYPE_SEED,
				})
			}
			break
		}
	}
}
