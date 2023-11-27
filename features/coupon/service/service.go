package service

import (
	"errors"
	"strings"
	"user_barang/features/coupon"
	"user_barang/helper/jwt"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type CouponService struct {
	m coupon.Repo
}

func New(model coupon.Repo) coupon.Service {
	return &CouponService{
		m: model,
	}
}

func (cs *CouponService) TambahCoupon(token *golangjwt.Token, newCoupon coupon.Coupon) (coupon.Coupon, error) {
	userId, err := jwt.ExtractToken(token)
	if err != nil {
		return coupon.Coupon{}, err
	}
	result, err := cs.m.InsertCoupon(userId, newCoupon)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return coupon.Coupon{}, errors.New("dobel input")
		}
		return coupon.Coupon{}, errors.New("terjadi kesalahan")
	}
	return result, nil
}

func (cs *CouponService) DapatCoupon(token *golangjwt.Token, page, limit int) ([]coupon.Coupon, error) {
	userId, err := jwt.ExtractToken(token)
	if err != nil {
		return nil, err
	}
	result, err := cs.m.GetCoupon(userId, page, limit)
	if err != nil {
		return nil, errors.New("failed to retrieve inserted coupon")
	}
	return result, nil
}

func (cs *CouponService) SemuaCoupon(page, limit int) ([]coupon.Coupon, error) {
	result, err := cs.m.GetCouponTanpa(page, limit)
	if err != nil {
		return nil, errors.New("failed to retrieve inserted coupon")
	}
	return result, nil
}
