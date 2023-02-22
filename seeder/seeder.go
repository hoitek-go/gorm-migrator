package seeder

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/hoitek-go/gorm-migrator/ports"
	"github.com/hoitek-go/gorm-migrator/share"
	"gorm.io/gorm"
)

var AllSeeds = []ports.SeederType{}

// Append MigrationType{} interface to AllMigrations variable.
func Add(m ...ports.SeederType) {
	AllSeeds = append(AllSeeds, m...)
}

// Make AllSeeds nil
func DeleteAllSeeds() {
	AllSeeds = nil
}

// Return TypeOf().Name() of given MigrationType{} interface.
func ModelName(m ports.SeederType) string {
	return reflect.TypeOf(m).Name()
}

// Return a migration by given name if type is seed.
func GetByName(name string) *share.Migration {
	var m share.Migration
	result := share.DB.Where("name = ? and type = ?", name, share.TYPE_SEED).First(&m)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &m
}

// Return all migration if type is seed.
func All() []share.Migration {
	var seeds []share.Migration
	result := share.DB.Model(&share.Migration{}).Order("id desc").Select("id, name, type").Where("type = ?", share.TYPE_SEED).Find(&seeds)
	if result.Error != nil || share.IsTestMode {
		return []share.Migration{}
	}
	return seeds
}

// Return desired number of migrations if type is seed.
func GetByLimit(number int) []share.Migration {
	var seeds []share.Migration
	share.DB.Model(&share.Migration{}).Order("id desc").Limit(number).Select("id, name, type").Where("type = ?", share.TYPE_SEED).Find(&seeds)
	return seeds
}

// Return all number of migration in int format if type is seed.
func Count() int {
	var count int64
	share.DB.Model(&share.Migration{}).Where("type = ?", share.TYPE_SEED).Count(&count)
	return int(count)
}

// Run migration by command.
func RunViaCommand(arg1, arg2 string) error {
	if arg1 == "" {
		return errors.New("First argument could not be empty")
	}
	if arg1 != "up" && arg1 != "down" {
		return errors.New("First argument just accepts up and down")
	}
	switch arg1 {
	case "up":
		if arg2 == "" {
			SeedUpAll()
		} else {
			SeedUp(arg2)
		}
	case "down":
		if arg2 == "" {
			SeedDownAll()
		} else {
			limit, err := strconv.Atoi(arg2)
			if err != nil {
				return err
			}
			SeedDown(limit)
		}
	}
	return nil
}
