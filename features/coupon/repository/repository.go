package repository

import (
	"Day25/features/coupon"

	"gorm.io/gorm"
)

type CouponModel struct {
	gorm.Model
	NamaProgram string
	Link        string
	Gambar      string
	UserID      uint
}

type CouponQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) coupon.Repo {
	return &CouponQuery{
		db: db,
	}
}

func (cq *CouponQuery) InsertCoupon(userID uint, newCoupon coupon.Coupon) (coupon.Coupon, error) {
	var inputDB = new(CouponModel)
	inputDB.UserID = userID
	inputDB.NamaProgram = newCoupon.NamaProgram
	inputDB.Link = newCoupon.Link
	inputDB.Gambar = newCoupon.Gambar

	if err := cq.db.Create(&inputDB).Error; err != nil {
		return coupon.Coupon{}, err
	}

	newCoupon.ID = inputDB.ID
	newCoupon.UserID = inputDB.UserID

	return newCoupon, nil
}

func (cq *CouponQuery) GetCoupon(userID uint, page, limit int) ([]coupon.Coupon, error) {
	var coupons []CouponModel
	offset := (page - 1) * limit
	if err := cq.db.Offset(offset).Limit(limit).Where("user_id = ?", userID).Find(&coupons).Error; err != nil {
		return nil, err
	}
	var result []coupon.Coupon
	for _, s := range coupons {
		result = append(result, coupon.Coupon{
			ID:          s.ID,
			NamaProgram: s.NamaProgram,
			Link:        s.Link,
			Gambar:      s.Gambar,
			UserID:      s.UserID,
		})
	}
	return result, nil
}

func (cq *CouponQuery) GetCouponTanpa(page, limit int) ([]coupon.Coupon, error) {
	var coupons []CouponModel
	offset := (page - 1) * limit
	if err := cq.db.Offset(offset).Limit(limit).Find(&coupons).Error; err != nil {
		return nil, err
	}
	var result []coupon.Coupon
	for _, s := range coupons {
		result = append(result, coupon.Coupon{
			ID:          s.ID,
			NamaProgram: s.NamaProgram,
			Link:        s.Link,
			Gambar:      s.Gambar,
			UserID:      s.UserID,
		})
	}
	return result, nil
}
