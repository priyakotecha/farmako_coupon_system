package router

import (
	"coupon_system/pkg/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl *controller.CouponController) *gin.Engine {
	r := gin.Default()

	r.POST("/coupons/applicable", ctrl.GetApplicable)
	r.POST("/coupons/validate", ctrl.Validate)

	return r
}
