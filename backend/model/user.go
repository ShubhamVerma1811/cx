package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `json:"id" gorm:"primaryKey;default:0;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func getUser(db *gorm.DB, id uint64) (User, error) {
	var user User
	tx := db.Where("id = ?", id).First(&user)
	return user, tx.Error
}

func createUser(db *gorm.DB, user User) (User, error) {
	tx := db.Create(&user)
	return user, tx.Error
}

func deleteUser(db *gorm.DB, user User) error {
	tx := db.Unscoped().Delete(&user)
	return tx.Error
}

func updateUser(db *gorm.DB, user User) (User, error) {
	tx := db.Save(&user)
	return user, tx.Error
}

func getUsers(db *gorm.DB) ([]User, error) {
	var users []User
	tx := db.Find(&users)
	return users, tx.Error
}
