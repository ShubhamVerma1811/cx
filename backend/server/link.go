package server

import (
	"cx/model"
	"errors"

	"gorm.io/gorm"
)

func GetLinks(db *gorm.DB) ([]model.Link, error) {
	var links []model.Link
	tx := db.Find(&links)
	return links, tx.Error
}

func GetLinkByLinkID(db *gorm.DB, id string) (model.Link, error) {
	var link model.Link
	tx := db.Where("id = ?", id).First(&link)
	return link, tx.Error
}

func GetLinkByShortURL(db *gorm.DB, shortURL string) (model.Link, error) {
	var link model.Link
	tx := db.Where("short_url = ?", shortURL).First(&link)
	return link, tx.Error
}

func CreateLink(db *gorm.DB, link *model.Link) error {
	tx := db.Create(link)
	return tx.Error
}

func DeleteLink(db *gorm.DB, id string) (model.Link, error) {

	link := model.Link{}

	tx := db.Where("id = ?", id).Unscoped().Delete(&link)

	if tx.Error != nil {
		return model.Link{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return model.Link{}, errors.New("Link not found")

	}

	return link, tx.Error
}

func UpdateLink(db *gorm.DB, link *model.Link) (model.Link, error) {
	tx := db.Where("id = ?", *&link.ID).Updates(&link)

	if tx.Error != nil {
		return model.Link{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return model.Link{}, errors.New("Link not found")
	}

	return *link, nil
}

func GetLinksByUserID(db *gorm.DB, userID uint64) ([]model.Link, error) {
	var links []model.Link
	tx := db.Where("user_id = ?", userID).Find(&links)
	return links, tx.Error
}
