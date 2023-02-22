package migrator

import (
	"errors"

	"github.com/hoitek-go/gorm-migrator/share"
)

// Make down desired number of migrations otherwise return error.
func MigrateDown(number int) error {
	share.DB.Migrator().AutoMigrate(&share.Migration{})
	if number > Count() {
		return errors.New("Your number of migrations is less than the entered number")
	}
	migrations := GetByLimit(number)
	ids := []uint{}
	migrationNames := map[string]string{}
	for _, row := range migrations {
		ids = append(ids, row.ID)
		migrationNames[row.Name] = row.Name
	}
	// I have problem with this if statement. when it returns false?
	if len(ids) > 0 {
		share.DB.Unscoped().Where("id in ? and type = ?", ids, share.TYPE_MIGRATION).Delete(&share.Migration{})
		for _, mStruct := range AllMigrations {
			name := ModelName(mStruct)
			_, ok := migrationNames[name]
			if ok {
				mStruct.Down()
			}
		}
	}
	return nil
}
