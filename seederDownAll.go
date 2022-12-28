package migrator

func SeedDownAll() {
	DB.Migrator().AutoMigrate(&Migration{})
	for _, mStruct := range AllSeeds {
		name := GetSeedName(mStruct)
		lastMigration := FindSeedByName(name)
		if lastMigration != nil {
			mStruct.Down()
			DB.Unscoped().Where("name = ? and type = ?", name, TYPE_SEED).Delete(&Migration{})
		}
	}
}
