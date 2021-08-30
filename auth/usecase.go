package auth

import (
	"context"

	"github.com/VolodimirKorpan/go_kobi/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
