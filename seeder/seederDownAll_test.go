package seeder

import (
	"testing"

	"github.com/hoitek-go/gorm-migrator/migrator"
	"github.com/hoitek-go/gorm-migrator/share"
	"github.com/stretchr/testify/assert"
)

func TestSeedDownAll(t *testing.T) {
	t.Run("Test SeedDownAll()", func(t *testing.T) {
		db := share.ConnectToDBTest()
		share.SetGorm(db)
		//Clean up database
		defer migrator.MigrateDown(4)
		Add(share.UserSeed{})
		defer DeleteAllSeeds()
		SeedUpAll()
		SeedDownAll()
		user := GetByName("UserSeed")
		if assert.Nil(t, user) {
			t.Logf("Test SeedDownAll() has been PASSED!")
		} else {
			t.Errorf("Test SeedDownAll() has been FAILED! Should down all migrations")
		}
	})
}
