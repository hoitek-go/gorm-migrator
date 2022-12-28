package migrator

func MigrateDownAll() {
	DB.Migrator().AutoMigrate(&Migration{})
	for _, mStruct := range AllMigrations {
		name := GetMigrationName(mStruct)
		lastMigration := FindMigrationByName(name)
		if lastMigration != nil {
			mStruct.Down()
			DB.Unscoped().Where("name = ? and type = ?", name, TYPE_MIGRATION).Delete(&Migration{})
		}
	}
}
