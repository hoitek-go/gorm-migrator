package migrator

import "testing"

func TestMigrateDownWhenThereIsNoMigrations(t *testing.T) {
	t.Run("Test Migrate Down When There is No Migrations", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetMigrations(User{})
		err := MigrateDown(5)
		if err == nil {
			t.Error("Should return error when number is greater than the count of the migrations")
		}
	})
}

func TestMigrateDownWhenThereIsMigrations(t *testing.T) {
	t.Run("Test Migrate Down When There is Migrations", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetMigrations(User{})
		MigrateUpAll()
		err := MigrateDown(1)
		if err != nil {
			t.Error(err)
		}
	})
}
