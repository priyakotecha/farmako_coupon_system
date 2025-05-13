package service

import (
	"coupon_system/pkg/model"
	"coupon_system/pkg/repository"
	"time"
)

type CouponService interface {
	GetApplicableCoupons(model.ApplicableCouponsRequest) ([]model.ApplicableCoupon, error)
	ValidateCoupon(model.ValidateCouponRequest) (*model.ValidateCouponResponse, error)
}

type couponService struct {
	repo repository.CouponRepository
}

func NewCouponService(r repository.CouponRepository) CouponService {
	return &couponService{repo: r}
}

func (s *couponService) GetApplicableCoupons(req model.ApplicableCouponsRequest) ([]model.ApplicableCoupon, error) {
	coupons, err := s.repo.GetAllCoupons()
	if err != nil {
		return nil, err
	}

	parsedTime, _ := time.Parse(time.RFC3339, req.Timestamp)
	var result []model.ApplicableCoupon

	for _, c := range coupons {
		if c.ExpiryDate.Before(parsedTime) || req.OrderTotal < c.MinOrderValue {
			continue
		}

		for _, item := range req.CartItems {
			if contains(c.ApplicableMedicineIDs, item.ID) || contains(c.ApplicableCategories, item.Category) {
				result = append(result, model.ApplicableCoupon{
					CouponCode:    c.Code,
					DiscountValue: c.DiscountValue,
				})
				break
			}
		}
	}

	return result, nil
}

func (s *couponService) ValidateCoupon(req model.ValidateCouponRequest) (*model.ValidateCouponResponse, error) {
	coupon, err := s.repo.GetCouponByCode(req.CouponCode)
	if err != nil {
		return &model.ValidateCouponResponse{
			IsValid: false,
			Reason:  "Coupon not found",
		}, nil
	}

	parsedTime, _ := time.Parse(time.RFC3339, req.Timestamp)
	if coupon.ExpiryDate.Before(parsedTime) || parsedTime.Before(coupon.ValidFrom) || parsedTime.After(coupon.ValidUntil) {
		return &model.ValidateCouponResponse{
			IsValid: false,
			Reason:  "Coupon expired or not in valid time window",
		}, nil
	}

	if req.OrderTotal < coupon.MinOrderValue {
		return &model.ValidateCouponResponse{
			IsValid: false,
			Reason:  "Order value too low",
		}, nil
	}

	resp := &model.ValidateCouponResponse{
		IsValid: true,
		Message: "coupon applied successfully",
	}

	if coupon.DiscountType == model.Inventory {
		resp.Discount.ItemsDiscount = coupon.DiscountValue
	} else if coupon.DiscountType == model.Charges {
		resp.Discount.ChargesDiscount = coupon.DiscountValue
	}

	return resp, nil
}

func contains(list []string, target string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}
