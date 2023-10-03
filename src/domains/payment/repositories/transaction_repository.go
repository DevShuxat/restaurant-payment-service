package repository

import (
	"context"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/payment/models"
)

type TransactionRepository interface {
	SaveTx(ctx context.Context, tx *models.Transaction) error
	UpdateTx(ctx context.Context, tx *models.Transaction) error
	GetTx(ctx context.Context, txID string) (*models.Transaction, error)
}