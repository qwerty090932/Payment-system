package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"payment_system/internal/config"
	"payment_system/internal/handlers"
	"payment_system/internal/repository"
	"payment_system/internal/usecase"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	connStr := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	walletRepo := repository.NewWalletRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	walletUseCase := usecase.NewWalletUseCase(walletRepo)
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepo, walletRepo)

	walletHandler := handlers.NewWalletHandler(walletUseCase)
	transactionHandler := handlers.NewTransactionHandler(transactionUseCase)

	r := gin.Default()

	r.GET("/api/wallet/:address/balance", walletHandler.GetBalance)
	r.POST("/api/send", transactionHandler.Send)
	r.GET("/api/transactions", transactionHandler.GetLastTransactions)

	// Initialize wallets with random addresses and 100.0 balance
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		address := generateRandomAddress()
		if err := walletUseCase.CreateWallet(address, 100.0); err != nil {
			log.Fatalf("Could not create wallet: %v", err)
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func generateRandomAddress() string {
	const charset = "abcdef0123456789"
	const length = 64
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
