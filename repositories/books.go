package repositories

import (
	"demo-restful-api/models"
	"errors"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*").Group("books.id")
	if err := query.Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

func GetBookByID(id string, db *gorm.DB) (models.Book, bool, error) {
	b := models.Book{}

	query := db.Select("books.*")
	query = db.Group("books.id")
	err := query.Where("books.id = ?", id).First(&b).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return b, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return b, false, nil
	}
	return b, true, nil
}

func DeleteBook(id string, db *gorm.DB) error {
	var b models.Book
	if err := db.Where("id = ?", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, b *models.Book) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}
