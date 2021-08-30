package auth

import (
	"context"

	"github.com/VolodimirKorpan/go_kobi/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.DBUser) error
	GetUser(ctx context.Context, username, password string) (*models.DBUser, error)
}
