package migrator

import "testing"

func TestSetArgsWhenArgsIsEmpty(t *testing.T) {
	t.Run("Test SetArgs When Args Is Empty", func(t *testing.T) {
		err := SetArgs([]string{})
		if err == nil {
			t.Error("Should return error when args is empty")
		}
	})
}

func TestSetArgsWhenTwoArgsAreReady(t *testing.T) {
	t.Run("Test SetArgs When Two Args Are Ready", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"migrate",
			"up",
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestSetArgsForUpBasedOnSpecifiecStruct(t *testing.T) {
	t.Run("Test Set Args For Up Based On Specifiec Struct", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"migrate",
			"up",
			"User",
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestSetArgsWhenArgsLengthMoreThan3(t *testing.T) {
	t.Run("Test SetArgs When Args Length More Than 3", func(t *testing.T) {
		err := SetArgs([]string{
			"arg1",
			"arg2",
			"arg3",
			"arg4",
		})
		if err != nil {
			t.Error("Should not return error when args is more than 3")
		}
	})
}

func TestSetArgsForDownAll(t *testing.T) {
	t.Run("Test Set Args For Down All", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"migrate",
			"down",
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestSetArgsForDownBasedOnNumber(t *testing.T) {
	t.Run("Test Set Args For Down Based On Number", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"migrate",
			"down",
			"1",
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestSetArgsForDownWhenLimitIsNotNumber(t *testing.T) {
	t.Run("Test SetArgs For Down When Limit Is Not Number", func(t *testing.T) {
		err := SetArgs([]string{
			"",
			"migrate",
			"down",
			"test",
		})
		if err == nil {
			t.Error("Should return error when limit is not number")
		}
	})
}

func TestGetMigrations(t *testing.T) {
	t.Run("Test GetMigrations", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetMigrations(User{})
		MigrateUpAll()
		migrations := getMigrations()
		MigrateDownAll()
		if len(migrations) != 1 {
			t.Error("Should return migrations properly")
		}
	})
}

func TestGetMigrationsWhenError(t *testing.T) {
	t.Run("Test GetMigrations When Error", func(t *testing.T) {
		db := ConnectToDBTest()
		SetGorm(db)
		SetMigrations(User{})
		MigrateUpAll()
		IsTestMode = true
		migrations := getMigrations()
		IsTestMode = false
		MigrateDownAll()
		if len(migrations) != 0 {
			t.Error("Should return no migrations when there is an error")
		}
	})
}

func TestConnectToTestDB(t *testing.T) {
	t.Run("Test Connect To Test DB", func(t *testing.T) {
		db := ConnectToDBTest()
		if db == nil {
			t.Error("There is an error when connect to database")
		}
	})
}

func TestConnectToTestDBWhenError(t *testing.T) {
	t.Run("Test Connect To Test DB When Error", func(t *testing.T) {
		IsTestMode = true
		db := ConnectToDBTest()
		IsTestMode = false
		if db != nil {
			t.Error("db instance only can be nil when there is an error")
		}
	})
}
