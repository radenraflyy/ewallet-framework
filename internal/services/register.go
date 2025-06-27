package services

import (
	"context"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	RegisterRepository interfaces.IRegisterRepository
}

func (s *RegisterService) RegisterUser(ctx context.Context, req *models.User) (interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPassword)

	if err := s.RegisterRepository.InsertNewUser(ctx, req); err != nil {
		return nil, err
	}

	resp := req
	resp.Password = "" // Clear password from response for security reasons
	return resp, nil
}
