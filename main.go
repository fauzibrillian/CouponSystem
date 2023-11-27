package main

import (
	"user_barang/config"
	ch "user_barang/features/coupon/handler"
	cr "user_barang/features/coupon/repository"
	cs "user_barang/features/coupon/service"
	uh "user_barang/features/user/handler"
	ur "user_barang/features/user/repository"
	us "user_barang/features/user/service"
	ek "user_barang/helper/enkrip"
	"user_barang/routes"
	"user_barang/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	if cfg == nil {
		e.Logger.Fatal("tidak bisa start server kesalahan database")
	}

	db, err := database.InitMySql(*cfg)
	if err != nil {
		e.Logger.Fatal("tidak bisa start bro", err.Error())
	}

	db.AutoMigrate(&ur.UserModel{}, cr.CouponModel{})

	ekrip := ek.New()
	userRepo := ur.New(db)
	userService := us.New(userRepo, ekrip)
	userHandler := uh.New(userService)

	couponRepo := cr.New(db)
	couponService := cs.New(couponRepo)
	couponHandler := ch.New(couponService)

	routes.InitRoute(e, userHandler, couponHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
