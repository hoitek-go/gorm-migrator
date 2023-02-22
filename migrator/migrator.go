package migrator

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/hoitek-go/gorm-migrator/ports"
	"github.com/hoitek-go/gorm-migrator/share"
	"gorm.io/gorm"
)

var AllMigrations = []ports.MigrationType{}

// Append MigrationType{} interface to AllMigrations variable.
func Add(m ...ports.MigrationType) {
	AllMigrations = append(AllMigrations, m...)
}

// Make AllMigrations nil
func DeleteAllMigrations() {
	AllMigrations = nil
}

// Return TypeOf().Name() of given MigrationType{} interface.
func ModelName(m ports.MigrationType) string {
	return reflect.TypeOf(m).Name()
}

// Return a migration by given name if type is migration.
func GetByName(name string) *share.Migration {

	var m share.Migration
	result := share.DB.Where("name = ? and type = ?", name, share.TYPE_MIGRATION).First(&m)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &m
}

// Return all migration if type is migration.
func All() []share.Migration {
	var migrations []share.Migration
	result := share.DB.Model(&share.Migration{}).Order("id desc").Select("id, name, type").Where("type = ?", share.TYPE_MIGRATION).Find(&migrations)
	if result.Error != nil || share.IsTestMode {
		return []share.Migration{}
	}
	return migrations
}

// Return desired number of migrations if type is migration.
func GetByLimit(number int) []share.Migration {
	var migrations []share.Migration
	share.DB.Model(&share.Migration{}).Order("id desc").Limit(number).Select("id, name, type").Where("type = ?", share.TYPE_MIGRATION).Find(&migrations)
	return migrations
}

// Return all number of migration in int format if type is migration.
func Count() int {
	var count int64
	share.DB.Model(&share.Migration{}).Where("type = ?", share.TYPE_MIGRATION).Count(&count)
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
			MigrateUpAll()
		} else {
			MigrateUp(arg2)
		}
	case "down":
		if arg2 == "" {
			MigrateDownAll()
		} else {
			limit, err := strconv.Atoi(arg2)
			if err != nil {
				return err
			}
			MigrateDown(limit)
		}
	}

	return nil
}
