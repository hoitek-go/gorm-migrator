package share

type UserSeed struct {
	Name string
}

func (u UserSeed) Up() {
	DB.Migrator().CreateTable(&UserSeed{})
	DB.Create([]UserSeed{
		{
			Name: "test",
		},
	})
}

func (u UserSeed) Down() {
	DB.Unscoped().Where("name in ?", []string{"test"}).Delete(&UserSeed{})
}

type User struct {
	Name string
}

func (u User) Up() {
	DB.Migrator().CreateTable(&User{})
}

func (u User) Down() {
	DB.Migrator().DropTable(&User{})
}
