package seeder

import (
	"testing"

	"github.com/hoitek-go/gorm-migrator/share"
	"github.com/stretchr/testify/assert"
)

func TestMigrateUp(t *testing.T) {
	t.Run("SeedUp() should make a seeder for given model.", func(t *testing.T) {
		share.PrepareDBForTest()
		// Clean up database
		defer share.CleanDBAfterTest()
		Add(share.UserSeed{})
		defer DeleteAllSeeds()
		SeedDownAll()
		SeedUp("UserSeed")
		userMigration := GetByName("UserSeed")
		if userMigration == nil {
			t.Errorf("Test SeedUp() has been FAILED! Cannot get migration by name.")
		}
		assert.Equal(t, userMigration.Name, "UserSeed")

	})
}
