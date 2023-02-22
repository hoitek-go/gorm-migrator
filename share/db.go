package share

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var IsTestMode = false

func SetGorm(instance *gorm.DB) *gorm.DB {
	DB = instance
	return DB
}

// Connect to postgres database.
func ConnectToDBTest() *gorm.DB {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=111111 dbname=gorm_migrator_test port=5432 sslmode=disable TimeZone=Asia/Tehran"))

	if err != nil || IsTestMode {
		return nil
	}

	log.Println("Connected to database")
	return db
}

// Make connection to DB and make a migration table with 3 records.
// Prepare functions for testing purpose.
func PrepareDBForTest() ([]*Migration, error) {
	db := ConnectToDBTest()
	SetGorm(db)
	if DB.Migrator().HasTable(&Migration{}) {
		err := DB.Migrator().DropTable(Migration{})
		if err != nil {
			return nil, err
		}

	}
	err := DB.Migrator().AutoMigrate(&Migration{})
	if err != nil {
		return nil, err
	}
	m := []Migration{
		{
			Name: "Test1",
			Type: TYPE_MIGRATION,
		},
		{
			Name: "Test2",
			Type: TYPE_MIGRATION,
		},
		{
			Name: "Test3",
			Type: TYPE_SEED,
		},
	}

	result := DB.Create(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	var queryResult []*Migration

	tx := DB.Find(&queryResult)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return queryResult, nil
}

func CleanDBAfterTest() {
	if DB.Migrator().HasTable(&Migration{}) {
		DB.Migrator().DropTable(Migration{})
	}

}
