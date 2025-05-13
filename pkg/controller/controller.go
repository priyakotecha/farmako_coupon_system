package controller

import (
	"coupon_system/pkg/model"
	"coupon_system/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CouponController struct {
	service service.CouponService
}

func NewCouponController(s service.CouponService) *CouponController {
	return &CouponController{service: s}
}

func (c *CouponController) GetApplicable(ctx *gin.Context) {
	var req model.ApplicableCouponsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, _ := c.service.GetApplicableCoupons(req)
	ctx.JSON(http.StatusOK, model.ApplicableCouponsResponse{ApplicableCoupons: result})
}

func (c *CouponController) Validate(ctx *gin.Context) {
	var req model.ValidateCouponRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, _ := c.service.ValidateCoupon(req)
	ctx.JSON(http.StatusOK, result)
}
