package migrator

import "testing"

func TestGetSeeds(t *testing.T) {
	t.Run("Test GetSeeds", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetSeeds(UserSeed{})
		SeedUpAll()
		seeds := GetSeeds()
		SeedDownAll()
		if len(seeds) != 1 {
			t.Error("Should return seeds properly")
		}
	})
}

func TestGetSeedsWhenError(t *testing.T) {
	t.Run("Test GetSeeds When Error", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetSeeds(UserSeed{})
		SeedUpAll()
		IsTestMode = true
		seeds := GetSeeds()
		IsTestMode = false
		SeedDownAll()
		if len(seeds) != 0 {
			t.Error("Should return no seeds when there is an error")
		}
	})
}

func TestSeederSetArgsWhenTwoArgsAreReady(t *testing.T) {
	t.Run("Test SetArgs When Two Args Are Ready", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"seed",
			"up",
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestSetArgsForSeedUpBasedOnSpecifiecStruct(t *testing.T) {
	t.Run("Test Set Args For Seed Up Based On Specifiec Struct", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"seed",
			"up",
			"UserSeed",
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestSeedUpBasedOnSpecifiecStruct(t *testing.T) {
	t.Run("Test Seed Up Based On Specifiec Struct", func(t *testing.T) {
		SeedDownAll()
		MigrateDownAll()
		SetMigrations(User{})
		SetSeeds(UserSeed{})
		MigrateUp("User")
		SeedUp("UserSeed")
		err := SetArgs([]string{
			"",
			"seed",
			"up",
			"UserSeed",
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestSeederSetArgsForDownAll(t *testing.T) {
	t.Run("Test Set Args For Down All", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"seed",
			"down",
		})
		if err != nil {
			t.Error(err)
		}
	})
}
func TestSeederSetArgsForDownBasedOnNumber(t *testing.T) {
	t.Run("Test Set Args For Down Based On Number", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"seed",
			"down",
			"1",
		})
		if err != nil {
			t.Error(err)
		}
	})
}
func TestSeederSetArgsForDownWhenLimitIsNotNumber(t *testing.T) {
	t.Run("Test SetArgs For Down When Limit Is Not Number", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"seed",
			"down",
			"test",
		})
		if err == nil {
			t.Error("Should return error when limit is not number")
		}
	})
}
