package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`         // Уникальный идентификатор транзакции
	From      string    `json:"from"`       // Адрес кошелька, с которого отправлены средства
	To        string    `json:"to"`         // Адрес кошелька, на который отправлены средства
	Amount    float64   `json:"amount"`     // Сумма перевода
	CreatedAt time.Time `json:"created_at"` // Время создания транзакции
}
