package service

import (
	"context"
	"errors"
	"time"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/card/models"
	repository "github.com/DevShuxat/restaurant-payment-service/src/domains/card/repositories"
)

type PaymentCardService interface {
	SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}

type paymetCardSvcImpl struct {
	paymentRepo repository.PaymentCardRepository
}

func NewPaymentCardRepository(paymentRepo repository.PaymentCardRepository) PaymentCardService {
	return &paymetCardSvcImpl{
		paymentRepo: paymentRepo,
	}
}

func (s *paymetCardSvcImpl) SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error) {
	payment := &models.PaymentCard{
		ID: cardID,
		CardToken: cardToken,
		CreatedAt: time.Now(),
	}
	err := s.paymentRepo.SaveCard(ctx, payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *paymetCardSvcImpl) DeleteCard(ctx context.Context, cardID string) error {
	exitingCard, err := s.paymentRepo.GetCard(ctx, cardID)
	if err != nil{
		return err
	}
	if exitingCard == nil {
		return errors.New("card not found")
	}
	return s.paymentRepo.DeleteCard(ctx, exitingCard.ID)
}
func (s *paymetCardSvcImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {
	return s.paymentRepo.GetCard(ctx, cardID)
}
