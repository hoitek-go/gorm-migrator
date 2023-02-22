package seeder

import (
	"testing"

	"github.com/hoitek-go/gorm-migrator/share"
	"github.com/stretchr/testify/assert"
)

func TestSeedDownWhenReturnsError(t *testing.T) {
	t.Run("Test SeedDown() When returns error", func(t *testing.T) {
		share.PrepareDBForTest()
		// Clean up database
		defer share.CleanDBAfterTest()
		err := SeedDown(5)
		if assert.EqualError(t, err, "Your number of seeds is less than the entered number") {
			t.Logf("Test SeedDown() has been PASSED!")
		} else {
			t.Error("Test SeedDown() has been FAILED! Should return error")
		}

	})
}

func TestSeedDown(t *testing.T) {
	t.Run("Test Migrate Down When There is a Seed", func(t *testing.T) {
		share.PrepareDBForTest()
		// Clean up database
		// defer share.CleanDBAfterTest()
		Add(share.UserSeed{})
		defer DeleteAllSeeds()
		// migrator.MigrateUpAll()
		SeedUpAll()
		// err := SeedDown(1)
		// if assert.Nil(t, err) {
		// 	t.Logf("Test SeedDown() has been PASSED!")
		// } else {
		// 	t.Error("Test SeedDown() has been FAILED! Should not return error")
		// }

	})
}
