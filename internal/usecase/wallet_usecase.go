package usecase

import (
	"payment_system/internal/repository"
)

type WalletUseCase struct {
	walletRepo *repository.WalletRepository
}

func NewWalletUseCase(walletRepo *repository.WalletRepository) *WalletUseCase {
	return &WalletUseCase{walletRepo: walletRepo}
}

func (uc *WalletUseCase) GetBalance(address string) (float64, error) {
	return uc.walletRepo.GetBalance(address)
}

func (uc *WalletUseCase) CreateWallet(address string, balance float64) error {
	return uc.walletRepo.CreateWallet(address, balance)
}

func (uc *WalletUseCase) UpdateBalance(address string, balance float64) error {
	return uc.walletRepo.UpdateBalance(address, balance)
}
