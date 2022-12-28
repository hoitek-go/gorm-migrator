package migrator

func SeedUp(seedName string) {
	DB.Migrator().AutoMigrate(&Migration{})
	for _, mStruct := range AllSeeds {
		name := GetSeedName(mStruct)
		if name == seedName {
			lastMigration := FindSeedByName(name)
			if lastMigration == nil {
				mStruct.Up()
				DB.Create(&Migration{
					Name: name,
					Type: TYPE_SEED,
				})
			}
			break
		}
	}
}
