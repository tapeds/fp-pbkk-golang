package migrations

import (
	"fmt"

	"github.com/tapeds/fp-pbkk-golang/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	queries := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
		`CREATE TYPE arah_enum AS ENUM ('BERANGKAT', 'DATANG');`,
	}

	for _, query := range queries {
		result := db.Exec(query)
		if result.Error != nil {
			fmt.Println("Error executing query:", result.Error)
		} else {
			fmt.Println("Executed query successfully:", query)
		}
	}

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Maskapai{},
		&entity.Bandara{},
		&entity.Penerbangan{},
		&entity.BandaraPenerbangan{},
		&entity.Tiket{},
		&entity.Penumpang{},
	); err != nil {
		return err
	}

	return nil
}
