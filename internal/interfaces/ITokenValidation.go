package interfaces

import (
	"context"
	pb "ewallet-ums/cmd/proto"
	"ewallet-ums/helpers"
)

type ITokenValidationHandler interface {
	TokenValidationHandler(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error)
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
