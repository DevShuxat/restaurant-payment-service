package payment

import (
	"context"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/payment/models"
	"github.com/DevShuxat/restaurant-payment-service/src/domains/payment/repositories"
	"gorm.io/gorm"
)

const (
	tableReceipt = "restaurant_payment.receipts"
)

type receiptSvcImpl struct {
	db *gorm.DB
}

func NewRecieptRepository(db *gorm.DB) repository.ReceiptRepository {
	return &receiptSvcImpl{
		db: db,
	}
}

func (r receiptSvcImpl)	SaveReceipt(ctx context.Context, receipt *models.Receipt) error {
	err := r.db.WithContext(ctx).Table(tableReceipt).Create(receipt).Error
	if err != nil {
		return err
	}
	return nil
}
func (r receiptSvcImpl)	GetReceipt(ctx context.Context, receiptID string) (*models.Receipt, error) {
	var receipt *models.Receipt
	result := r.db.WithContext(ctx).Table(tableReceipt).First(&receipt, receiptID)
	if result.Error != nil {
		return nil, result.Error
	}
	return receipt, nil
}
func (r receiptSvcImpl)	GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error) {
	var receipt *models.Receipt
	result := r.db.WithContext(ctx).Table(tableReceipt).Where(&receipt, "id = ?", orderID).Find(&receipt)
	if result.Error != nil {
		return nil, result.Error
	}
	return receipt, nil
}
