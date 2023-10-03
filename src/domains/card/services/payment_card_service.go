package service

import (
	"context"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/card/models"
)

type PaymentCardService interface {
	SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}