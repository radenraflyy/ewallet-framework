package services

import (
	"context"
	"ewallet-ums/internal/interfaces"
	"strings"
)

type LogoutService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LogoutService) Logout(ctx context.Context, token string, userID uint) error {
	parts := strings.Split(token, "Bearer ")
	token = parts[1]
	err := s.UserRepo.DeleteUserSession(ctx, token, userID)
	if err != nil {
		return err
	}

	return nil
}
