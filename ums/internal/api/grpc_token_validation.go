package api

import (
	"context"
	"fmt"
	tokenvalidation "ums/cmd/proto"
	"ums/constant"
	"ums/helpers"
	"ums/internal/interfaces"
)

type TokenValidationHandler struct {
	TokenValidationService interfaces.ITokenValidationService
	tokenvalidation.UnimplementedTokenValidationServiceServer
}

func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error) {
	var (
		log   = helpers.Logger
		token = req.GetToken()
	)

	if token == "" {
		err := fmt.Errorf("token is empty")
		log.Error(err)
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	log.Println("token: ", token)
	session, err := s.TokenValidationService.TokenValidation(ctx, token)
	log.Println("session.Username: ", session.Username)
	if err != nil {
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &tokenvalidation.TokenResponse{
		Message: constant.SuccessMessage,
		Data: &tokenvalidation.UserData{
			Username: session.Username,
			UserId:   int64(session.UserID),
		},
	}, nil
}
