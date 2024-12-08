package migrations

import (
	"log"

	"github.com/tapeds/fp-pbkk-golang/entity"
	"gorm.io/gorm"
)

func dropAllTables(db *gorm.DB) error {
	if err := db.Migrator().DropTable(
		&entity.Maskapai{},
		&entity.User{},
		&entity.Penerbangan{},
		&entity.BandaraPenerbangan{},
		&entity.Tiket{},
		&entity.Penumpang{},
		&entity.Bandara{},
	); err != nil {
		return err
	}

	log.Println("All tables dropped successfully.")
	return nil
}

func Fresh(db *gorm.DB) error {
	if err := dropAllTables(db); err != nil {
		log.Printf("Error dropping tables: %v", err)
		return err
	}

	if err := Migrate(db); err != nil {
		log.Printf("Error during migration: %v", err)
		return err
	}

	return nil
}
