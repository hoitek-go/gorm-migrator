package migrator

import (
	"reflect"
	"testing"

	"github.com/hoitek-go/gorm-migrator/share"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("Add() should add a model to AllMigrations", func(t *testing.T) {
		Add(share.User{}, share.UserSeed{})
		defer DeleteAllMigrations()
		if assert.Equal(t, len(AllMigrations), 2) {
			t.Logf("Test Add() has been PASSED!")
		} else {
			t.Errorf("Test Add() has been FAILED!")
		}

	})

}

func TestDeleteAllMigrations(t *testing.T) {
	t.Run("DeleteAllMigrations() should empty the AllMigrations slice", func(t *testing.T) {
		Add(share.User{}, share.UserSeed{})
		DeleteAllMigrations()

		assert.Equal(t, len(AllMigrations), 0)
	})
}
func TestModelName(t *testing.T) {
	tc := share.User{Name: "test"}
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

type migrationTestCase struct {
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
	mtc := []migrationTestCase{
		{
			testName:      "GetByName() should returns a migration by its name",
			migrationName: "Test1",
			expected:      migrations[0],
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
	_, err := share.PrepareDBForTest()
	// Clean up database
	defer share.CleanDBAfterTest()
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Run("All(), returns all migrations", func(t *testing.T) {
		migrations := All()
		assert.Equal(t, len(migrations), 2)
	})
}

func TestAllReturnEmptySlice(t *testing.T) {
	_, err := share.PrepareDBForTest()
	// Clean up database
	defer share.CleanDBAfterTest()
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Run("All() should returns empty slice", func(t *testing.T) {
		share.IsTestMode = true
		migrations := All()
		share.IsTestMode = false

		if assert.Empty(t, migrations) {
			t.Logf("Test All() has been PASSED! returned empty slice in Test mode.")
		} else {
			t.Errorf("Test All() FAILED! Should return empty slice.")
		}
	})
}

func TestGetByLimit(t *testing.T) {
	_, err := share.PrepareDBForTest()
	// Clean up database
	defer share.CleanDBAfterTest()
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Run("GetByLimit() should return migrations base on the given number.", func(t *testing.T) {
		migrations := GetByLimit(2)
		assert.Equal(t, len(migrations), 2)
	})

}

func TestCount(t *testing.T) {
	_, err := share.PrepareDBForTest()
	// Clean up database
	defer share.CleanDBAfterTest()
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Run("Count() should return number of migrations in db.", func(t *testing.T) {
		numOfMigrations := Count()

		assert.Equal(t, 2, numOfMigrations)
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
			arg2: "User",
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
	defer DeleteAllMigrations()
	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {
			Add(share.User{})
			err := RunViaCommand(v.arg1, v.arg2)
			if err != nil {
				t.Errorf("Test RunViaCommand() has been FAILED! Should not return Error.")
			}
			user := GetByName("User")

			assert.Equal(t, user.Name, "User")
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
	defer DeleteAllMigrations()
	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {
			Add(share.User{})
			MigrateUp("User")
			err := RunViaCommand(v.arg1, v.arg2)
			if err != nil {
				t.Errorf("Test RunViaCommand() has been FAILED! Should not return Error.")
			}
			user := GetByName("User")

			if reflect.ValueOf(user).IsZero() {
				t.Logf("Test RunViaCommand() has been PASSED!")
			}

		})
	}

}
