package models

import "errors"

var (
	ErrInsufficientFunds = errors.New("insufficient funds") // Ошибка, если на кошельке недостаточно средств
	ErrWalletNotFound    = errors.New("wallet not found")   // Ошибка, если кошелек не найден
)
