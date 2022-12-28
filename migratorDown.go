package migrator

import (
	"errors"
)

func MigrateDown(number int) error {
	DB.Migrator().AutoMigrate(&Migration{})
	if number > CountMigrations() {
		return errors.New("Your number of migrations is less than the entered number")
	}
	migrations := getMigrationsByLimit(number)
	ids := []uint{}
	migrationNames := map[string]string{}
	for _, row := range migrations {
		ids = append(ids, row.ID)
		migrationNames[row.Name] = row.Name
	}
	if len(ids) > 0 {
		DB.Unscoped().Where("id in ? type = ?", ids, TYPE_MIGRATION).Delete(&Migration{})
		for _, mStruct := range AllMigrations {
			name := GetMigrationName(mStruct)
			_, ok := migrationNames[name]
			if ok {
				mStruct.Down()
			}
		}
	}
	return nil
}
