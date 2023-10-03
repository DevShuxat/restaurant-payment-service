package services

import (
	"context"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/payment/models"
)

type PaymentApplicationService interface {
	CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error)
	
}