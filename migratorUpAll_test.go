package migrator

import "testing"

func TestMigrateUpAll(t *testing.T) {
	db := ConnectToDBTest()
	SetGorm(db)
	SetMigrations(User{})
	MigrateDownAll()
	MigrateUpAll()
	count := CountMigrations()
	if count != 1 {
		t.Error("Should create name")
	}
	MigrateDownAll()
}
