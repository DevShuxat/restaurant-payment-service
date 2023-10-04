package service

import (
	"context"
	"time"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/payment/models"
	repository "github.com/DevShuxat/restaurant-payment-service/src/domains/payment/repositories"
)

type PaymentService interface {
	CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error)
	PayReceipt(ctx context.Context, receiptID string) (*models.Receipt, error)
	GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error)
}

type paymentSvcImpl struct {
	paymentRepo repository.ReceiptRepository
}

func NewPaymentService(
	paymentRepo repository.ReceiptRepository,
) PaymentService {
	return &paymentSvcImpl{
		paymentRepo: paymentRepo,
	}
}

func (s *paymentSvcImpl) CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error) {
	receipt := &models.Receipt{
		OrderID:   orderID,
		CardID:    cardID,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	err := s.paymentRepo.SaveReceipt(ctx, receipt)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (s *paymentSvcImpl) PayReceipt(ctx context.Context, receiptID string) (*models.Receipt, error) {
	receipt, err := s.paymentRepo.GetReceipt(ctx, receiptID)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (s *paymentSvcImpl) GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error) {
	reciept, err :=  s.paymentRepo.GetReceiptByOrder(ctx,orderID)
	if err != nil {
		return nil, err
	}
	return reciept, nil
}