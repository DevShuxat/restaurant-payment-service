package service

import (
	"context"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/payment/models"
)

type PaymentService interface {
	CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error)
	PayReceipt(ctx context.Context, receiptID string) error
	GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error)
	SaveTransaction(ctx context.Context, tx *models.Transaction) error
	UpdateTransaction(ctx context.Context, tx *models.Transaction) error
	GetTransaction(ctx context.Context, txID string) (*models.Transaction, error)
}