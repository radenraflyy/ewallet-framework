package external

import (
	"context"
	"fmt"
	"wallet/constant"
	tokenvalidation "wallet/external/proto"
	"wallet/internal/models"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func ValidateToken(ctx context.Context, token string) (models.TokenData, error) {
	var (
		resp models.TokenData
	)

	conn, err := grpc.Dial("localhost:7000", grpc.WithInsecure())
	if err != nil {
		return resp, errors.Wrap(err, "failed to dial ums grpc")
	}
	defer conn.Close()

	client := tokenvalidation.NewTokenValidationServiceClient(conn)

	req := &tokenvalidation.TokenRequest{
		Token: token,
	}
	tokenResp, err := client.ValidateToken(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "failed to validate token")
	}

	if tokenResp.Message != constant.SuccessMessage {
		return resp, fmt.Errorf("got response error from ums: %s", tokenResp.Message)
	}

	resp.UserID = tokenResp.Data.UserId
	resp.Username = tokenResp.Data.Username

	return resp, nil
}
