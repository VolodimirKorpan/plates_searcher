package mysql

import (
	"context"
	"errors"

	"github.com/VolodimirKorpan/go_kobi/models"
)

type UserRepository struct {
	db *MySQL
}

func NewUserRepository(db *MySQL) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *models.DBUser)  error {
	if user == nil {
		return errors.New("No user provided")
	}

	err := repo.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUser(ctx context.Context, username, password string) (*models.DBUser, error) {
	var user *models.DBUser

	err := repo.db.Where(&models.User{Username: username, Password: password}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}
