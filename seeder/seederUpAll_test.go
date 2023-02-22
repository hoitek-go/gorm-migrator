package seeder

import (
	"testing"

	"github.com/hoitek-go/gorm-migrator/share"
	"github.com/stretchr/testify/assert"
)

func TestMigrateUpAll(t *testing.T) {
	t.Run("Test SeedUpAll()", func(t *testing.T) {
		db := share.ConnectToDBTest()
		share.SetGorm(db)
		//Clean up database
		defer share.CleanDBAfterTest()
		Add(share.User{}, share.UserSeed{})
		defer DeleteAllSeeds()
		SeedDownAll()
		SeedUpAll()
		userMigration := GetByName("User")
		userSeed := GetByName("UserSeed")
		if userMigration == nil || userSeed == nil {
			t.Errorf("Test SeedUpAll() has been FAILED! Cannot get migration by name.")
		}

		assert.Equal(t, userMigration.Name, "User")
		assert.Equal(t, userSeed.Name, "UserSeed")

	})

}
