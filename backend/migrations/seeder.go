package migrations

import (
	"github.com/tapeds/fp-pbkk-golang/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}

	return nil
}
