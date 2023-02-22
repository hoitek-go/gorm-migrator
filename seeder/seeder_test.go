package seeder

import (
	"reflect"
	"testing"

	"github.com/hoitek-go/gorm-migrator/share"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("Add() should add a model to AllSeeds", func(t *testing.T) {
		Add(share.User{}, share.UserSeed{})
		if assert.Equal(t, len(AllSeeds), 2) {
			t.Logf("Test Add() has been PASSED!")
		} else {
			t.Errorf("Test Add() has been FAILED!")
		}
	})

}

func TestDeleteAllSeeds(t *testing.T) {
	t.Run("TestDeleteAllSeeds() should empty the AllSeeds slice", func(t *testing.T) {
		Add(share.User{}, share.UserSeed{})
		DeleteAllSeeds()

		assert.Equal(t, len(AllSeeds), 0)
	})

}
func TestModelName(t *testing.T) {
	tc := share.UserSeed{Name: "test"}
	tcName := reflect.TypeOf(tc).Name()
	t.Run("ModelName() should return struct name", func(t *testing.T) {
		name := ModelName(tc)
		if assert.Equal(t, name, tcName) {
			t.Logf("Testing ModelName() has been PASSED!")
		} else {
			t.Errorf("Testing ModelName() has been FAILED!")
		}
	})

}

type seederTestCase struct {
	testName      string
	migrationName string
	expected      *share.Migration
}

func TestGetByName(t *testing.T) {
	migrations, err := share.PrepareDBForTest()
	// Clean up database
	defer share.CleanDBAfterTest()
	if err != nil {
		t.Errorf(err.Error())
	}
	mtc := []seederTestCase{
		{
			testName:      "GetByName() should returns a seeds by its name",
			migrationName: "Test3",
			expected:      migrations[2],
		},
		{
			testName:      "GetByName() should return nil",
			migrationName: "notExist",
			expected:      nil,
		},
	}

	for _, v := range mtc {
		t.Run(v.testName, func(t *testing.T) {
			res := GetByName(v.migrationName)
			assert.Equal(t, v.expected, res)
		})
	}
}

func TestAll(t *testing.T) {
	t.Run("All(), returns all seeds", func(t *testing.T) {
		_, err := share.PrepareDBForTest()
		// Clean up database
		defer share.CleanDBAfterTest()
		if err != nil {
			t.Errorf(err.Error())
		}
		seeds := All()

		assert.Equal(t, len(seeds), 1)
	})
}

func TestAllReturnEmptySlice(t *testing.T) {
	t.Run("All() should returns empty slice", func(t *testing.T) {
		_, err := share.PrepareDBForTest()
		// Clean up database
		defer share.CleanDBAfterTest()
		if err != nil {
			t.Errorf(err.Error())
		}
		share.IsTestMode = true
		seeds := All()
		share.IsTestMode = false

		if assert.Empty(t, seeds) {
			t.Logf("Test All() has been PASSED! returned empty slice in Test mode.")
		} else {
			t.Errorf("Test All() FAILED! Should return empty slice.")
		}
	})
}
func TestGetByLimit(t *testing.T) {
	t.Run("GetByLimit() should return seeds base on the given number.", func(t *testing.T) {
		_, err := share.PrepareDBForTest()
		// Clean up database
		defer share.CleanDBAfterTest()
		if err != nil {
			t.Errorf(err.Error())
		}

		seeds := GetByLimit(1)
		assert.Equal(t, len(seeds), 1)
	})
}

func TestCount(t *testing.T) {
	t.Run("Count() should return number of seeds in db.", func(t *testing.T) {
		_, err := share.PrepareDBForTest()
		// Clean up database
		defer share.CleanDBAfterTest()
		if err != nil {
			t.Errorf(err.Error())
		}

		numOfSeeds := Count()

		assert.Equal(t, numOfSeeds, 1)
	})

}

type runViaCommandTestCases struct {
	name     string
	arg1     string
	arg2     string
	expected string
}

func TestRunViaCommandFirstArgs(t *testing.T) {
	tc := []runViaCommandTestCases{
		{
			name:     "Test RunViaCommand() when args is empty",
			arg1:     "",
			arg2:     "",
			expected: "First argument could not be empty",
		},
		{
			name:     "Test RunViaCommand() when first args is not `up` or `down`",
			arg1:     "upgrade",
			arg2:     "",
			expected: "First argument just accepts up and down",
		},
	}
	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {
			err := RunViaCommand(v.arg1, v.arg2)
			assert.EqualError(t, err, v.expected)
		})
	}
}
func TestRunViaCommandWhenFirstArgsIsUp(t *testing.T) {
	tc := []runViaCommandTestCases{
		{
			name: "Test RunViaCommand() when second args is User",
			arg1: "up",
			arg2: "UserSeed",
		},
		{
			name: "Test RunViaCommand() when second args is empty",
			arg1: "up",
			arg2: "",
		},
	}
	share.PrepareDBForTest()
	// Clean up database
	defer share.CleanDBAfterTest()
	defer DeleteAllSeeds()
	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {

			Add(share.UserSeed{})
			err := RunViaCommand(v.arg1, v.arg2)
			if err != nil {
				t.Errorf("Test RunViaCommand() has been FAILED! Should not return Error.")
			}
			user := GetByName("UserSeed")

			assert.Equal(t, user.Name, "UserSeed")
		})
	}

}

func TestRunViaCommandWhenFirstArgsIsDown(t *testing.T) {
	tc := []runViaCommandTestCases{
		{
			name: "Test RunViaCommand() when second args is User",
			arg1: "down",
			arg2: "",
		},
		{
			name: "Test RunViaCommand() when second args is empty",
			arg1: "down",
			arg2: "1",
		},
	}
	share.PrepareDBForTest()
	// Clean up database
	defer share.CleanDBAfterTest()

	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {
			Add(share.UserSeed{})
			DeleteAllSeeds()
			SeedUp("UserSeed")
			err := RunViaCommand(v.arg1, v.arg2)
			if err != nil {
				t.Errorf("Test RunViaCommand() has been FAILED! Should not return Error.")
			}
			user := GetByName("UserSeed")

			if reflect.ValueOf(user).IsZero() {
				t.Logf("Test RunViaCommand() has been PASSED!")
			}

		})
	}

}
