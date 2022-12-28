package migrator

import (
	"errors"
)

func SeedDown(number int) error {
	DB.Migrator().AutoMigrate(&Migration{})
	if number > CountSeeds() {
		return errors.New("Your number of seeds is less than the entered number")
	}
	seeds := GetSeedsByLimit(number)
	ids := []uint{}
	seedNames := map[string]string{}
	for _, row := range seeds {
		ids = append(ids, row.ID)
		seedNames[row.Name] = row.Name
	}
	if len(ids) > 0 {
		DB.Unscoped().Where("id in ? and type = ?", ids, TYPE_SEED).Delete(&Migration{})
		for _, mStruct := range AllSeeds {
			name := GetSeedName(mStruct)
			_, ok := seedNames[name]
			if ok {
				mStruct.Down()
			}
		}
	}
	return nil
}
