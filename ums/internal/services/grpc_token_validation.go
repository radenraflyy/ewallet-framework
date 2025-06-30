package services

import (
	"context"
	"ums/helpers"
	"ums/internal/interfaces"

	"github.com/pkg/errors"
)

type TokenValidationService struct {
	UserRepository interfaces.IUserRepository
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	var (
		claimToken *helpers.ClaimToken
	)
	claimToken, err := helpers.ValidateToken(ctx, token)
	if err != nil {
		return claimToken, errors.Wrap(err, "failed to validate token")
	}

	_, err = s.UserRepository.GetUserSessionByToken(ctx, token)
	if err != nil {
		return claimToken, errors.Wrap(err, "failed to get user session by token")
	}

	return claimToken, nil
}
