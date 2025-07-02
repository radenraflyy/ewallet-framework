package interfaces

import (
	"context"
	"ums/external"
)

type IWallet interface {
	CreateWallet(ctx context.Context, userID int) (*external.Wallet, error)
}
