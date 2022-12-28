package migrator

import "testing"

func TestMigrateDownAll(t *testing.T) {
	t.Run("Test Migrate Down All", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetMigrations(User{})
		MigrateUpAll()
		MigrateDownAll()
		count := CountMigrations()
		if count > 0 {
			t.Error("Should delete all migrations")
		}
	})
}
