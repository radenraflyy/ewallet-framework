package repository

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var (
		user models.User
		err  error
	)
	err = r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.Wrap(err, "user not found")
	}
	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.Create(session).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string, userID uint) error {
	err := r.DB.Exec("DELETE FROM users_sessions WHERE user_id = ? AND token = ?", userID, token).Error
	if err != nil {
		return errors.Wrap(err, "failed to delete user session")
	}
	if r.DB.RowsAffected == 0 {
		return errors.New("no session found for the given user ID and token")
	}
	return nil
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error) {
	var (
		session models.UserSession
		err     error
	)
	if token == "" {
		return session, errors.Wrap(err, "token is empty")
	}

	err = r.DB.Where("token = ?", token).First(&session).Error
	if err != nil {
		return session, err
	}
	if session.ID == 0 {
		return session, errors.Wrap(err, "session not found")
	}
	return session, nil
}
