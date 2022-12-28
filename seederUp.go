package migrator

func SeedUp() {
	DB.Migrator().AutoMigrate(&Migration{})
	for _, mStruct := range AllSeeds {
		name := GetSeedName(mStruct)
		lastMigration := FindSeedByName(name)
		if lastMigration == nil {
			mStruct.Up()
			DB.Create(&Migration{
				Name: name,
				Type: TYPE_SEED,
			})
		}
	}
}
