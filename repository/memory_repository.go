package repository

import (
	"context"
	"errors"
	"sync"
	"user-api/models"
)

type MemoryRepository struct {
	wallets map[string]*models.Wallet
	mutex   sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		wallets: make(map[string]*models.Wallet),
	}
}

func (r *MemoryRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.wallets[wallet.ID.Hex()] = wallet
	return nil
}

func (r *MemoryRepository) GetWallet(ctx context.Context, id string) (*models.Wallet, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	wallet, exists := r.wallets[id]
	if !exists {
		return nil, errors.New("wallet not found")
	}
	return wallet, nil
}

func (r *MemoryRepository) UpdateWallet(ctx context.Context, wallet *models.Wallet) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.wallets[wallet.ID.Hex()]; !exists {
		return errors.New("wallet not found")
	}

	r.wallets[wallet.ID.Hex()] = wallet
	return nil
}

func (r *MemoryRepository) ListWallets(ctx context.Context) ([]*models.Wallet, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	wallets := make([]*models.Wallet, 0, len(r.wallets))
	for _, wallet := range r.wallets {
		wallets = append(wallets, wallet)
	}
	return wallets, nil
}
