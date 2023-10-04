package service

import (
	"context"
	"time"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/merchant/models"
	repository "github.com/DevShuxat/restaurant-payment-service/src/domains/merchant/repositories"
)

type MerchantSettingService interface {
	CreateMerchantSetting(ctx context.Context, entityID, cashboxID string, Enabled bool) (*models.MerchantSetting, error)
	UpdateMerchantSetting(ctx context.Context, CashboxID string, Enabled bool) (*models.MerchantSetting, error)
	GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
}

type merchantSvcImpl struct {
	merchantRepo repository.MerchantRepository
}

func NewMerchantSettingService(
	merchantRepo repository.MerchantRepository,
) MerchantSettingService {
	return &merchantSvcImpl{
		merchantRepo: merchantRepo,
	}
}

func (s merchantSvcImpl) CreateMerchantSetting(ctx context.Context, entityID, cashboxID string, enabled bool) (*models.MerchantSetting, error)	{
	merchant := &models.MerchantSetting{
		EntityID: entityID,
		CashboxID: cashboxID,
		Enabled: enabled,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.merchantRepo.SaveMerchantSetting(ctx, merchant)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}

func (s merchantSvcImpl) UpdateMerchantSetting(ctx context.Context, CashboxID string, Enabled bool) (*models.MerchantSetting, error) {
    	merchant := models.MerchantSetting{
			CashboxID: CashboxID,
			Enabled: Enabled,
			UpdatedAt: time.Now().UTC(),
		}
		err := s.merchantRepo.UpdateMerchantSetting(ctx, &merchant)
		if err != nil {
			return nil, err
		}
    return &merchant, err
}

	func (s merchantSvcImpl) GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error){
		return s.merchantRepo.GetMerchantSetting(ctx, entityID)
	}
