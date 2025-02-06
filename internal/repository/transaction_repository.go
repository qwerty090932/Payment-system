package repository

import (
	"database/sql"
	"payment_system/internal/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(from, to string, amount float64) error {
	_, err := r.db.Exec("INSERT INTO transactions (from_address, to_address, amount) VALUES ($1, $2, $3)", from, to, amount)
	return err
}

func (r *TransactionRepository) GetLastTransactions(count int) ([]models.Transaction, error) {
	rows, err := r.db.Query("SELECT id, from_address, to_address, amount, created_at FROM transactions ORDER BY created_at DESC LIMIT $1", count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.From, &t.To, &t.Amount, &t.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
