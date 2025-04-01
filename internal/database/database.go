package database

import (
	"gorm.io/gorm"
	"test/internal/models"
)

type Database struct {
	connection *gorm.DB
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{connection: conn}
}

func (d *Database) DeleteUser(id string) error {
	err := d.connection.Table("users_t").Where("id", id).UpdateColumn("active", false).Error
	if err != nil {
		return err
	}
	return nil
}
func (d *Database) GetUsers() ([]models.UsersT, error) {
	var users []models.UsersT
	d.connection.Table("users_t").Find(&users)
	if d.connection.Error != nil {
		return nil, d.connection.Error
	}
	return users, nil
}
