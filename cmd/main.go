package main

import (
	"coupon_system/pkg/controller"
	"coupon_system/pkg/model"
	"coupon_system/pkg/repository"
	"coupon_system/pkg/router"
	"coupon_system/pkg/service"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=db user=user password=password dbname=coupon_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&model.Coupon{})

	repo := repository.NewCouponRepository(db)
	svc := service.NewCouponService(repo)
	ctrl := controller.NewCouponController(svc)

	r := router.SetupRouter(ctrl)
	r.Run(":8080")
}
