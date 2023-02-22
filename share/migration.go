package share

import "gorm.io/gorm"

const (
	TYPE_MIGRATION = "migration"
	TYPE_SEED      = "seed"
)

type Migration struct {
	gorm.Model
	Name string
	Type string
}
