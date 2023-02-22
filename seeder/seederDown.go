package seeder

import (
	"errors"

	"github.com/hoitek-go/gorm-migrator/share"
)

// Make down desired number of seeds otherwise return error.
func SeedDown(number int) error {
	share.DB.Migrator().AutoMigrate(&share.Migration{})
	if number > Count() {
		return errors.New("Your number of seeds is less than the entered number")
	}
	seeds := GetByLimit(number)
	ids := []uint{}
	seedNames := map[string]string{}
	for _, row := range seeds {
		ids = append(ids, row.ID)
		seedNames[row.Name] = row.Name
	}
	if len(ids) > 0 {
		share.DB.Unscoped().Where("id in ? and type = ?", ids, share.TYPE_SEED).Delete(&share.Migration{})
		for _, mStruct := range AllSeeds {
			name := ModelName(mStruct)
			_, ok := seedNames[name]
			if ok {
				mStruct.Down()
			}
		}
	}
	return nil
}
