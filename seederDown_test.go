package migrator

import (
	"testing"
)

func TestSeedDownWhenThereIsNoSeeds(t *testing.T) {
	t.Run("Test Seed Down When There is No Seeds", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetSeeds(UserSeed{})
		err := SeedDown(5)
		if err == nil {
			t.Error("Should return error when number is greater than the count of the seeds")
		}
	})
}

func TestSeedDownWhenThereIsSeeds(t *testing.T) {
	t.Run("Test Seed Down When There is Seeds", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetSeeds(UserSeed{})
		SeedUpAll()
		err := SeedDown(1)
		if err != nil {
			t.Error(err)
		}
	})
}
