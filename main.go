package main

import (
	"Day25/config"
	ch "Day25/features/coupon/handler"
	cr "Day25/features/coupon/repository"
	cs "Day25/features/coupon/service"
	uh "Day25/features/user/handler"
	ur "Day25/features/user/repository"
	us "Day25/features/user/service"
	ek "Day25/helper/enkrip"
	"Day25/routes"
	"Day25/utils/database"

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
