package migrator

import (
	"testing"
)

func TestSeedUp(t *testing.T) {
	db := ConnectToDBTest()
	SetGorm(db)
	SetSeeds(UserSeed{})
	SeedDownAll()
	SeedUpAll()
	count := CountSeeds()
	if count != 1 {
		t.Error("Should seed up")
	}
	SeedDownAll()
}
