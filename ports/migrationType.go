package ports

type MigrationType interface {
	Up()
	Down()
}
