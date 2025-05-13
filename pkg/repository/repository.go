package repository

import (
	"coupon_system/pkg/model"

	"gorm.io/gorm"
)

type CouponRepository interface {
	GetAllCoupons() ([]model.Coupon, error)
	GetCouponByCode(code string) (*model.Coupon, error)
}

type couponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) CouponRepository {
	return &couponRepository{db: db}
}

func (r *couponRepository) GetAllCoupons() ([]model.Coupon, error) {
	var coupons []model.Coupon
	err := r.db.Find(&coupons).Error
	return coupons, err
}

func (r *couponRepository) GetCouponByCode(code string) (*model.Coupon, error) {
	var coupon model.Coupon
	err := r.db.First(&coupon, "code = ?", code).Error
	return &coupon, err
}
