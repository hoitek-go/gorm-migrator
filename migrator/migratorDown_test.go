package migrator

import (
	"testing"

	"github.com/hoitek-go/gorm-migrator/share"
	"github.com/stretchr/testify/assert"
)

func TestMigrateDownWhenReturnsError(t *testing.T) {
	t.Run("Test Migrate Down When returns error", func(t *testing.T) {
		share.PrepareDBForTest()
		// Clean up database
		defer share.CleanDBAfterTest()
		err := MigrateDown(5)
		if assert.EqualError(t, err, "Your number of migrations is less than the entered number") {
			t.Logf("Test MigrationDown() has been PASSED!")
		} else {
			t.Error("Test MigrationDown() has been FAILED! Should return error")
		}

	})
}

func TestMigrateDown(t *testing.T) {
	share.PrepareDBForTest()
	// Clean up database
	defer share.CleanDBAfterTest()
	t.Run("Test Migrate Down When There is Migrations", func(t *testing.T) {

		Add(share.User{})
		defer DeleteAllMigrations()
		MigrateUpAll()
		err := MigrateDown(1)
		if err != nil {
			t.Error("Test MigrationDown() has been FAILED! Should not return error")

		} else {
			t.Logf("Test MigrationDown() has been PASSED!")
		}

	})
}
