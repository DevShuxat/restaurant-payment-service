package card

import (
	"context"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/card/models"
	repository "github.com/DevShuxat/restaurant-payment-service/src/domains/card/repositories"
	"gorm.io/gorm"
)

const (
	tableCard = "restaurant_payment.cards"
)

type cardSvcImpl struct {
	db *gorm.DB
}

func NewPaymentCardRepository(db *gorm.DB) repository.PaymentCardRepository {
	return &cardSvcImpl{
		db: db,
	}
}

func (r *cardSvcImpl) 	SaveCard(ctx context.Context, card *models.PaymentCard) error {
	err := r.db.WithContext(ctx).Table(tableCard).Create(&card).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *cardSvcImpl) DeleteCard(ctx context.Context, cardID string) error {
	var card *models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableCard).Delete(&card, "id = ?", cardID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *cardSvcImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {
	var card *models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableCard).First(&card, cardID)
	if result.Error != nil {
		return nil, result.Error
	}
	return card, nil
}