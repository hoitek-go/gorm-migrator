package share

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type connectToDBTestCases struct {
	testName string
}

var tc = []connectToDBTestCases{
	{
		testName: "Test Connect To Test DB",
	},
	{
		testName: "Test Connect To Test DB When Error",
	},
}

func TestConnectToDB(t *testing.T) {

	for i, v := range tc {

		t.Run(v.testName, func(t *testing.T) {

			if i == 1 {
				IsTestMode = true
				db := ConnectToDBTest()
				IsTestMode = false
				if assert.Nil(t, db) {
					t.Logf("db instance only can be nil when there is an error")
				}
			} else {
				db := ConnectToDBTest()
				if assert.NotNil(t, db) {
					t.Logf("Connected to DB successfully")
				} else {
					t.Error("There is an error when connect to database")
				}
			}

		})
	}
}

func TestSetGorm(t *testing.T) {
	t.Run("Test SetGorm()", func(t *testing.T) {
		db := ConnectToDBTest()

		if db == nil {
			t.Errorf("TestSetGorm Failed while trying to connect to DB")
		}

		instance := SetGorm(db)

		assert.Equal(t, instance, DB)
	})

}

func TestPrepareDBForTest(t *testing.T) {
	t.Run("Test PrepareDBForTest()", func(t *testing.T) {
		migrations, err := PrepareDBForTest()
		//Clean up database
		defer CleanDBAfterTest()
		if err != nil {
			t.Errorf("Test PrepareDBForTest() has been FAILED! Should not return error")
		}

		assert.Equal(t, 3, len(migrations))
	})
}

func TestCleanDBAfterTest(t *testing.T) {
	t.Run("TestCleanDBAfterTest() should drop Migration Table", func(t *testing.T) {
		PrepareDBForTest()
		CleanDBAfterTest()
		if DB.Migrator().HasTable(&Migration{}) {
			t.Errorf("TestCleanDBAfterTest() has been FAILED! Should drop Migration Table")
		} else {
			t.Logf("TestCleanDBAfterTest() has been PASSED!")
		}
	})

}
