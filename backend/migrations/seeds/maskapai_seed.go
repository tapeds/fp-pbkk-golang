package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/tapeds/fp-pbkk-golang/entity"
	"gorm.io/gorm"
)

func ListMaskapaiSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/maskapai.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listMaskapai []entity.Maskapai
	if err := json.Unmarshal(jsonData, &listMaskapai); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Maskapai{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Maskapai{}); err != nil {
			return err
		}
	}

	for _, data := range listMaskapai {
		var maskapai entity.Maskapai
		err := db.Where(&entity.Maskapai{Name: data.Name}).First(&maskapai).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&maskapai, "name = ?", data.Name).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
