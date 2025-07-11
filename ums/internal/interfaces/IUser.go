package interfaces

import (
	"context"
	"ums/internal/models"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	DeleteUserSession(ctx context.Context, token string, userID uint) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
}
