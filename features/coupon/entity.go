package coupon

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Coupon struct {
	ID          uint
	NamaProgram string
	Link        string
	Gambar      string
	UserID      uint
}

type Handler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetAllTanpa() echo.HandlerFunc
}

type Service interface {
	TambahCoupon(token *jwt.Token, newCoupon Coupon) (Coupon, error)
	DapatCoupon(token *jwt.Token, page, limit int) ([]Coupon, error)
	SemuaCoupon(page, limit int) ([]Coupon, error)
}

type Repo interface {
	InsertCoupon(userId uint, newCoupon Coupon) (Coupon, error)
	GetCoupon(userId uint, page, limit int) ([]Coupon, error)
	GetCouponTanpa(page, limit int) ([]Coupon, error)
}
