package migrator

const (
	TYPE_MIGRATION = "migration"
	TYPE_SEED      = "seed"
)

type UserSeed struct {
	Name string
}

func (u UserSeed) Up() {
	db := ConnectToDBTest()
	db.Migrator().CreateTable(&UserSeed{})
	db.Create([]UserSeed{
		{
			Name: "test",
		},
	})
}

func (u UserSeed) Down() {
	db := ConnectToDBTest()
	db.Unscoped().Where("name in ?", []string{"test"}).Delete(&UserSeed{})
}

type User struct {
	Name string
}

func (u User) Up() {
	db := ConnectToDBTest()
	db.Migrator().CreateTable(&User{})
}

func (u User) Down() {
	db := ConnectToDBTest()
	db.Migrator().DropTable(&User{})
}
