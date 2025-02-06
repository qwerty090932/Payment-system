package usecase

import (
	"payment_system/internal/models"
	"payment_system/internal/repository"
)

type TransactionUseCase struct {
	transactionRepo *repository.TransactionRepository
	walletRepo      *repository.WalletRepository
}

func NewTransactionUseCase(transactionRepo *repository.TransactionRepository, walletRepo *repository.WalletRepository) *TransactionUseCase {
	return &TransactionUseCase{transactionRepo: transactionRepo, walletRepo: walletRepo}
}

func (uc *TransactionUseCase) Send(from, to string, amount float64) error {
	balance, err := uc.walletRepo.GetBalance(from)
	if err != nil {
		return err
	}

	if balance < amount {
		return models.ErrInsufficientFunds
	}

	if err := uc.walletRepo.UpdateBalance(from, balance-amount); err != nil {
		return err
	}

	toBalance, err := uc.walletRepo.GetBalance(to)
	if err != nil {
		return err
	}

	if err := uc.walletRepo.UpdateBalance(to, toBalance+amount); err != nil {
		return err
	}

	return uc.transactionRepo.CreateTransaction(from, to, amount)
}

func (uc *TransactionUseCase) GetLastTransactions(count int) ([]models.Transaction, error) {
	return uc.transactionRepo.GetLastTransactions(count)
}
