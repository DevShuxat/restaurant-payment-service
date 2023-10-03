package service

import (
	"context"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/merchant/models"
)

type MerchantSetting interface {
	CreateMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
	UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error
	GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
}