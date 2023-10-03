package merchant

import (
	"context"

	"github.com/DevShuxat/restaurant-payment-service/src/domains/merchant/models"
	"github.com/DevShuxat/restaurant-payment-service/src/domains/merchant/repositories"
	"gorm.io/gorm"
)

const (
	tableMerchant = "restaurant_payment.merchant_settings"
)

type merchantSvcImpl struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) repository.MerchantRepository {
	return &merchantSvcImpl{
		db: db,
	}
}

func (r *merchantSvcImpl) SaveMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error {
	err := r.db.WithContext(ctx).Table(tableMerchant).Create(&setting).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *merchantSvcImpl) UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error {
	err := r.db.WithContext(ctx).Table(tableMerchant).Save(&setting).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *merchantSvcImpl) GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {
	var merchant *models.MerchantSetting
	result := r.db.WithContext(ctx).Table(tableMerchant).First(&merchant, entityID)
	if result.Error != nil {
		return nil, result.Error
	}
	return merchant, nil
}