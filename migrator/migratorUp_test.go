package migrator

import (
	"testing"

	"github.com/hoitek-go/gorm-migrator/share"
	"github.com/stretchr/testify/assert"
)

func TestMigrateUp(t *testing.T) {
	t.Run("MigrateUp() should make a migration for given model.", func(t *testing.T) {
		share.PrepareDBForTest()
		// Clean up database
		defer share.CleanDBAfterTest()
		Add(share.User{})
		defer DeleteAllMigrations()
		MigrateDownAll()
		MigrateUp("User")
		userMigration := GetByName("User")
		if userMigration == nil {
			t.Errorf("Test MigrationUp() has been FAILED! Cannot get migration by name.")
		}
		assert.Equal(t, userMigration.Name, "User")

	})
}
