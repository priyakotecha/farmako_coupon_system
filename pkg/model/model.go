package model

import "time"

type UsageType string

const (
	OneTime   UsageType = "one_time"
	MultiUse  UsageType = "multi_use"
	TimeBased UsageType = "time_based"
)

type DiscountType string

const (
	Inventory DiscountType = "inventory"
	Charges   DiscountType = "charges"
)

type Coupon struct {
	Code                  string `gorm:"primaryKey"`
	ExpiryDate            time.Time
	UsageType             UsageType
	ApplicableMedicineIDs []string
	ApplicableCategories  []string
	MinOrderValue         float64
	ValidFrom             time.Time
	ValidUntil            time.Time
	TermsAndConditions    string
	DiscountType          DiscountType
	DiscountValue         float64
	MaxUsagePerUser       int
}

type CartItem struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

type ApplicableCouponsRequest struct {
	CartItems  []CartItem `json:"cart_items"`
	OrderTotal float64    `json:"order_total"`
	Timestamp  string     `json:"timestamp"`
}

type ApplicableCoupon struct {
	CouponCode    string  `json:"coupon_code"`
	DiscountValue float64 `json:"discount_value"`
}

type ApplicableCouponsResponse struct {
	ApplicableCoupons []ApplicableCoupon `json:"applicable_coupons"`
}

type ValidateCouponRequest struct {
	CouponCode string     `json:"coupon_code"`
	CartItems  []CartItem `json:"cart_items"`
	OrderTotal float64    `json:"order_total"`
	Timestamp  string     `json:"timestamp"`
}

type ValidateCouponResponse struct {
	IsValid  bool   `json:"is_valid"`
	Message  string `json:"message"`
	Discount struct {
		ItemsDiscount   float64 `json:"items_discount"`
		ChargesDiscount float64 `json:"charges_discount"`
	} `json:"discount,omitempty"`
	Reason string `json:"reason,omitempty"`
}
