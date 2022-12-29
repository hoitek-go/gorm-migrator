package migrator

import (
	"reflect"
	"strconv"

	"github.com/hoitek-go/gorm-migrator/ports"
	"gorm.io/gorm"
)

var AllSeeds = []ports.SeederType{}

func GetSeedName(m ports.SeederType) string {
	return reflect.TypeOf(m).Name()
}

func FindSeedByName(name string) *Migration {
	var m Migration
	result := DB.Where("name = ? and type = ?", name, TYPE_SEED).First(&m)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &m
}

func SetSeeds(m ...ports.SeederType) {
	AllSeeds = append(AllSeeds, m...)
}

func GetSeeds() []Migration {
	var seeds []Migration
	result := DB.Model(&Migration{}).Order("id desc").Select("id, name, type").Where("type = ?", TYPE_SEED).Find(&seeds)
	if result.Error != nil || IsTestMode {
		return []Migration{}
	}
	return seeds
}

func GetSeedsByLimit(number int) []Migration {
	var seeds []Migration
	DB.Model(&Migration{}).Order("id desc").Limit(number).Select("id, name, type").Where("type = ?", TYPE_SEED).Find(&seeds)
	return seeds
}

func CountSeeds() int {
	var count int64
	DB.Model(&Migration{}).Where("type = ?", TYPE_SEED).Count(&count)
	return int(count)
}

func runSeedViaCommand(arg1, arg2 string) error {
	if arg1 == "up" {
		if arg2 == "" {
			SeedUpAll()
		} else {
			SeedUp(arg2)
		}
	}
	if arg1 == "down" {
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
