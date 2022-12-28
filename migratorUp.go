package migrator

func MigrateUp() {
	DB.Migrator().AutoMigrate(&Migration{})
	for _, mStruct := range AllMigrations {
		name := GetMigrationName(mStruct)
		lastMigration := FindMigrationByName(name)
		if lastMigration == nil {
			mStruct.Up()
			DB.Create(&Migration{
				Name: name,
				Type: TYPE_MIGRATION,
			})
		}
	}
}
