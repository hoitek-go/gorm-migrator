package migrator

import (
	"testing"
)

func TestSeedDownAll(t *testing.T) {
	t.Run("Test Seed Down All", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetSeeds(&UserSeed{})
		SeedUp()
		SeedDownAll()
		count := CountSeeds()
		if count > 0 {
			t.Error("Should delete all seeds")
		}
	})
}
