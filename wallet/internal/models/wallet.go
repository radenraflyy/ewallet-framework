package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Wallet struct {
	ID        int       `json:"id" gorm:"primaryKey" validate:"required"`
	UserID    int       `json:"user_id" gorm:"not null;unique" validate:"required"`
	Balance   float64   `json:"balance" gorm:"not null;type:decimal(15,2)" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (*Wallet) TableName() string {
	return "wallets"
}

func (l Wallet) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type WalletTrasaction struct {
	ID                   int       `json:"id" gorm:"primaryKey"`
	WalletID             int       `json:"wallet_id" gorm:"not null" validate:"required"`
	Amount               float64   `json:"balance" gorm:"not null;type:decimal(15,2)" validate:"required"`
	WalletTrasactionType string    `json:"wallet_transaction_type" gorm:"not null;type:ENUM('CREDIT', 'DEBIT')" validate:"required"`
	Reference            string    `json:"reference" gorm:"not null;type:varchar(255)" validate:"required"`
	CreatedAt            time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt            time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (*WalletTrasaction) TableName() string {
	return "Wallet_trasaction"
}

func (l WalletTrasaction) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
