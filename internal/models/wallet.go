package models

type Wallet struct {
	Address string  `json:"address"` // Уникальный адрес кошелька
	Balance float64 `json:"balance"` // Баланс кошелька
}
