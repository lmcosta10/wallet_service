package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmcosta10/wallet_service/internal/model"
)

type UserService struct {
	DB *pgxpool.Pool
}

type WalletService struct {
	DB *pgxpool.Pool
}

// ====================
// User services
// ====================

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) FetchUserByID(c *gin.Context, id int) {
	var user model.User

	user.Id = id

	err := s.DB.QueryRow(
		c.Request.Context(),
		`SELECT username FROM users WHERE id = $1`,
		id,
	).Scan(&user.Username)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// ====================
// Wallet services
// ====================

func NewWalletService(db *pgxpool.Pool) *WalletService {
	return &WalletService{DB: db}
}

func (s *WalletService) Transfer(
	ctx context.Context,
	fromWalletID int,
	toWalletID int,
	amount float64,
) error {

	tx, err := s.DB.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var fromBalance float64

	// 1) Lock sender wallet
	err = tx.QueryRow(
		ctx,
		`SELECT balance FROM wallets WHERE wallet_id = $1 FOR UPDATE`,
		fromWalletID,
	).Scan(&fromBalance)

	if err != nil {
		return err
	}

	if fromBalance < amount {
		return errors.New("insufficient funds")
	}

	// 2) Lock receiver wallet
	_, err = tx.Exec(
		ctx,
		`SELECT 1 FROM wallets WHERE wallet_id = $1 FOR UPDATE`,
		toWalletID,
	)
	if err != nil {
		return err
	}

	// 3) Debit sender
	_, err = tx.Exec(
		ctx,
		`UPDATE wallets SET balance = balance - $1 WHERE wallet_id = $2`,
		amount,
		fromWalletID,
	)
	if err != nil {
		return err
	}

	// 4) Credit receiver
	_, err = tx.Exec(
		ctx,
		`UPDATE wallets SET balance = balance + $1 WHERE wallet_id = $2`,
		amount,
		toWalletID,
	)
	if err != nil {
		return err
	}

	// 5) Commit
	return tx.Commit(ctx)
}
