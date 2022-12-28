package migrator

import "testing"

func TestMigrateUp(t *testing.T) {
	db := ConnectToDBTest()
	SetGorm(db)
	SetMigrations(&User{})
	MigrateDownAll()
	MigrateUp()
	count := CountMigrations()
	if count != 1 {
		t.Error("Should create name")
	}
	MigrateDownAll()
}
