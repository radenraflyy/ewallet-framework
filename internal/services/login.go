package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(c context.Context, req models.UserLogin) (models.UserLoginResponse, error) {
	var (
		resp models.UserLoginResponse
		now  = time.Now()
	)
	userDetial, err := s.UserRepo.GetUserByUsername(c, req.Username)
	if err != nil {
		return resp, errors.Wrap(err, "failed to get user by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetial.Password), []byte(req.Password)); err != nil {
		return resp, errors.Wrap(err, "failed to compare password")
	}

	token, err := helpers.GenerateToken(c, userDetial.ID, userDetial.Username, userDetial.FullName, "token", now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate token")
	}

	refresh_token, err := helpers.GenerateToken(c, userDetial.ID, userDetial.Username, userDetial.FullName, "refresh_token", now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate token")
	}

	userSession := models.UserSession{
		UserID:             userDetial.ID,
		Token:              token,
		RefreshToken:       refresh_token,
		TokenExpiry:        now.Add(helpers.MapTokenType["token"]),
		RefreshTokenExpiry: now.Add(helpers.MapTokenType["refresh_token"]),
	}
	if err := s.UserRepo.InsertNewUserSession(c, &userSession); err != nil {
		return resp, errors.Wrap(err, "failed to insert user session")
	}

	resp.UserID = userDetial.ID
	resp.Username = userDetial.Username
	resp.FullName = userDetial.FullName
	resp.Email = userDetial.Email
	resp.Token = token
	resp.RefreshToken = refresh_token
	return resp, nil
}
