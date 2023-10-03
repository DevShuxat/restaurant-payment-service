package repository

import (
	"context"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/card/models"
)

type PaymentCardRepository interface {
	SaveCard(ctx context.Context, card *models.PaymentCard) error
	DeleteCard(ctx context.Context, cardID string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}