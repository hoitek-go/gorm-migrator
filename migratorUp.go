package migrator

func MigrateUp(migrationName string) {
	DB.Migrator().AutoMigrate(&Migration{})
	for _, mStruct := range AllMigrations {
		name := GetMigrationName(mStruct)
		if name == migrationName {
			lastMigration := FindMigrationByName(name)
			if lastMigration == nil {
				mStruct.Up()
				DB.Create(&Migration{
					Name: name,
					Type: TYPE_MIGRATION,
				})
			}
			break
		}
	}
}
