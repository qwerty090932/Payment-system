package repository

import (
	"database/sql"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) CreateWallet(address string, balance float64) error {
	_, err := r.db.Exec("INSERT INTO wallets (address, balance) VALUES ($1, $2)", address, balance)
	return err
}

func (r *WalletRepository) GetBalance(address string) (float64, error) {
	var balance float64
	err := r.db.QueryRow("SELECT balance FROM wallets WHERE address = $1", address).Scan(&balance)
	return balance, err
}

func (r *WalletRepository) UpdateBalance(address string, balance float64) error {
	_, err := r.db.Exec("UPDATE wallets SET balance = $1 WHERE address = $2", balance, address)
	return err
}
