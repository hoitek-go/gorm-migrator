package migrator

import (
	"reflect"
	"testing"

	"github.com/hoitek-go/gorm-migrator/share"
)

func TestMigrateDownAll(t *testing.T) {
	t.Run("Test MigrateDownAll()", func(t *testing.T) {
		db := share.ConnectToDBTest()
		share.SetGorm(db)
		//Clean up database
		defer share.CleanDBAfterTest()
		Add(share.User{})
		defer DeleteAllMigrations()
		MigrateUpAll()
		MigrateDownAll()
		user := GetByName("User")
		if reflect.ValueOf(user).IsZero() {
			t.Logf("Test MigrateDownAll() has been PASSED!")
		} else {
			t.Errorf("Test MigrateDownAll() has been FAILED! Should down all migrations")
		}
	})
}
