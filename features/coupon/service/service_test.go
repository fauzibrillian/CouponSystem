package service_test

import (
	"Day25/features/coupon"
	"Day25/features/coupon/mocks"
	"Day25/features/coupon/service"
	"Day25/helper/jwt"
	"errors"
	"testing"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var userID = uint(1)
var str, _ = jwt.GenerateJWT(userID)
var token, _ = gojwt.Parse(str, func(t *gojwt.Token) (interface{}, error) {
	return []byte("$!1gnK3yyy!!!"), nil
})

func TestTambahCoupon(t *testing.T) {
	repo := mocks.NewRepo(t)
	m := service.New(repo)

	var repoData = coupon.Coupon{NamaProgram: "hoax", Link: "wwww.facebook.com", Gambar: "hoax.png", UserID: userID}
	var falseData = coupon.Coupon{}

	t.Run("Success Case", func(t *testing.T) {
		repo.On("InsertCoupon", userID, repoData).Return(repoData, nil).Once()
		res, err := m.TambahCoupon(token, repoData)

		assert.Nil(t, err)
		assert.Equal(t, userID, res.UserID)
		assert.Equal(t, "hoax", res.NamaProgram)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		repo.On("InsertCoupon", userID, falseData).Return(falseData, errors.New("ERROR db")).Once()
		res, err := m.TambahCoupon(token, falseData)
		assert.Error(t, err)
		assert.Equal(t, uint(0), res.UserID)
		assert.Equal(t, "", res.NamaProgram)
		repo.AssertExpectations(t)
	})
}

func TestDapatCoupon(t *testing.T) {
	repo := mocks.NewRepo(t)
	couponService := service.New(repo)

	t.Run("Success Case", func(t *testing.T) {
		expectedResult := []coupon.Coupon{
			{ID: 1, NamaProgram: "Coupon1", Link: "www.coupon1.com", UserID: userID},
			{ID: 2, NamaProgram: "Coupon2", Link: "www.coupon2.com", UserID: userID},
		}

		repo.On("GetCoupon", userID, 1, 10).Return(expectedResult, nil).Once()

		result, err := couponService.DapatCoupon(token, 1, 10)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)

		repo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		repo.On("GetCoupon", userID, 1, 10).Return(nil, errors.New("repository error")).Once()

		result, err := couponService.DapatCoupon(token, 1, 10)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "failed to retrieve inserted coupon", err.Error())

		repo.AssertExpectations(t)
	})
}

func TestSemuaCoupon(t *testing.T) {
	repo := mocks.NewRepo(t)
	couponService := service.New(repo)

	t.Run("Success Case", func(t *testing.T) {
		expectedResult := []coupon.Coupon{
			{ID: 1, NamaProgram: "Coupon1", Link: "www.coupon1.com", UserID: 1},
			{ID: 2, NamaProgram: "Coupon2", Link: "www.coupon2.com", UserID: 1},
		}

		repo.On("GetCouponTanpa", 1, 10).Return(expectedResult, nil).Once()

		result, err := couponService.SemuaCoupon(1, 10)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)

		repo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		repo.On("GetCouponTanpa", 1, 10).Return(nil, errors.New("repository error")).Once()

		result, err := couponService.SemuaCoupon(1, 10)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "failed to retrieve inserted coupon", err.Error())

		repo.AssertExpectations(t)
	})
}
