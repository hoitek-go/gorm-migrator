package seeder

import "github.com/hoitek-go/gorm-migrator/share"

// Make down all seeds.
func SeedDownAll() {
	share.DB.Migrator().AutoMigrate(&share.Migration{})
	for _, mStruct := range AllSeeds {
		name := ModelName(mStruct)
		lastMigration := GetByName(name)
		if lastMigration != nil {
			mStruct.Down()
			share.DB.Unscoped().Where("name = ? and type = ?", name, share.TYPE_SEED).Delete(&share.Migration{})
		}
	}
}
