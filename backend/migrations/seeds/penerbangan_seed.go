package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/tapeds/fp-pbkk-golang/entity"
	"gorm.io/gorm"
)

func ListPenerbanganSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/penerbangan.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listPenerbangan []entity.Penerbangan
	if err := json.Unmarshal(jsonData, &listPenerbangan); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Penerbangan{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Penerbangan{}); err != nil {
			return err
		}
	}

	for _, data := range listPenerbangan {
		var penerbangan entity.Penerbangan
		err := db.Where(&entity.Penerbangan{NoPenerbangan: data.NoPenerbangan}).First(&penerbangan).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&penerbangan, "no_penerbangan = ?", data.NoPenerbangan).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
