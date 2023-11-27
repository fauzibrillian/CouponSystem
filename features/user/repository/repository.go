package repository

import (
	"Day25/features/coupon/repository"
	"Day25/features/user"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Nama         string
	Password     string
	CouponModels []repository.CouponModel `gorm:"foreignKey:UserID"`
}

type UserQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &UserQuery{
		db: db,
	}
}

func (uq *UserQuery) Insert(newUser user.User) (user.User, error) {
	var inputDB = new(UserModel)
	inputDB.Nama = newUser.Nama
	inputDB.Password = newUser.Password

	if err := uq.db.Create(&inputDB).Error; err != nil {
		return user.User{}, err
	}

	newUser.ID = inputDB.ID

	return newUser, nil
}

func (uq *UserQuery) Login(nama string) (user.User, error) {
	var userData = new(UserModel)

	if err := uq.db.Where("nama = ?", nama).First(userData).Error; err != nil {
		return user.User{}, err
	}

	var result = new(user.User)
	result.ID = userData.ID
	result.Nama = userData.Nama
	result.Password = userData.Password

	return *result, nil
}
