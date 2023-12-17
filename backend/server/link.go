package server

import (
	"cx/model"

	"gorm.io/gorm"
)

func GetLinks(db *gorm.DB) ([]model.Link, error) {
	var links []model.Link
	tx := db.Find(&links)
	return links, tx.Error
}

func GetLinkByLinkID(db *gorm.DB, id string) (model.Link, error) {
	var link model.Link
	tx := db.Where("id = ? AND user_id = ?", id).First(&link)
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

func DeleteLink(db *gorm.DB, id string) error {
	tx := db.Unscoped().Delete(&model.Link{
		ShortURL: id,
	})
	return tx.Error
}

func UpdateLink(db *gorm.DB, id string) (model.Link, error) {
	tx := db.Save(id)

	if tx.Error != nil {
		return model.Link{}, tx.Error
	}

	link, err := GetLinkByLinkID(db, id)

	return link, err
}

func GetLinksByUserID(db *gorm.DB, userID uint64) ([]model.Link, error) {
	var links []model.Link
	tx := db.Where("user_id = ?", userID).Find(&links)
	return links, tx.Error
}
