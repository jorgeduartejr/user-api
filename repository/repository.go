package repository

import (
	"context"
	"user-api/models"
)

type Repository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	GetWallet(ctx context.Context, id string) (*models.Wallet, error)
	UpdateWallet(ctx context.Context, wallet *models.Wallet) error
	ListWallets(ctx context.Context) ([]*models.Wallet, error)
}
