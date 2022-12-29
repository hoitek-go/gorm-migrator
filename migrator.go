package migrator

import (
	"errors"
	"log"
	"reflect"
	"strconv"

	"github.com/hoitek-go/gorm-migrator/ports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var AllMigrations = []ports.MigrationType{}
var DB *gorm.DB
var IsTestMode = false

type Migration struct {
	gorm.Model
	Name string
	Type string
}

func SetGorm(instance *gorm.DB) {
	DB = instance
}

func GetMigrationName(m ports.MigrationType) string {
	return reflect.TypeOf(m).Name()
}

func FindMigrationByName(name string) *Migration {
	var m Migration
	result := DB.Where("name = ? and type = ?", name, TYPE_MIGRATION).First(&m)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &m
}

func SetMigrations(m ...ports.MigrationType) {
	AllMigrations = append(AllMigrations, m...)
}

func SetArgs(args []string) error {
	if len(args) < 3 {
		return errors.New("command not found!")
	}
	command := args[1]
	arg1 := args[2]
	arg2 := ""
	if len(args) > 3 {
		arg2 = args[3]
	}
	switch command {
	case "migrate":
		return runMigrateViaCommand(arg1, arg2)
	case "seed":
		return runSeedViaCommand(arg1, arg2)
	}
	return nil
}

func runMigrateViaCommand(arg1, arg2 string) error {
	if arg1 == "up" {
		if arg2 == "" {
			MigrateUpAll()
		} else {
			MigrateUp(arg2)
		}
	}
	if arg1 == "down" {
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

func getMigrations() []Migration {
	var migrations []Migration
	result := DB.Model(&Migration{}).Order("id desc").Select("id, name, type").Where("type = ?", TYPE_MIGRATION).Find(&migrations)
	if result.Error != nil || IsTestMode {
		return []Migration{}
	}
	return migrations
}

func getMigrationsByLimit(number int) []Migration {
	var migrations []Migration
	DB.Model(&Migration{}).Order("id desc").Limit(number).Select("id, name, type").Where("type = ?", TYPE_MIGRATION).Find(&migrations)
	return migrations
}

func CountMigrations() int {
	var count int64
	DB.Model(&Migration{}).Where("type = ?", TYPE_MIGRATION).Count(&count)
	return int(count)
}

func ConnectToDBTest() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=111111 dbname=gorm_migrator_test port=5432 sslmode=disable TimeZone=Asia/Tehran"))
	if err != nil || IsTestMode {
		return nil
	}
	log.Println("Connected to database")
	return db
}
