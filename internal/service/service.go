package service

import (
	"test/internal/database"
	"test/internal/models"
)

type Service struct {
	database *database.Database
}

func NewService(db *database.Database) *Service {
	return &Service{database: db}
}

func (s *Service) DeleteUser(id string) error {
	err := s.database.DeleteUser(id)
	return err
}

func (s *Service) GetUsers() ([]models.UsersT, error) {
	users, err := s.database.GetUsers()
	return users, err
}
