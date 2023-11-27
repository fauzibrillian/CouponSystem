package service

import (
	"Day25/features/user"
	"Day25/helper/enkrip"
	"errors"
	"strings"
)

type userService struct {
	repo user.Repository
	h    enkrip.HashInterface
}

func New(r user.Repository, h enkrip.HashInterface) user.Service {
	return &userService{
		repo: r,
		h:    h,
	}
}

func (us *userService) Register(newUser user.User) (user.User, error) {
	// validasi
	if newUser.Password == "" {
		return user.User{}, errors.New("password cannot be empty")
	}
	// enkripsi password
	ePassword, err := us.h.HashPassword(newUser.Password)

	if err != nil {
		return user.User{}, errors.New("terdapat masalah saat memproses data")
	}

	newUser.Password = ePassword
	result, err := us.repo.Insert(newUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return user.User{}, errors.New("data telah terdaftar pada sistem")
		}
		return user.User{}, errors.New("terjadi kesalahan pada sistem")
	}

	return result, nil
}

func (us *userService) Login(nama string, password string) (user.User, error) {
	if nama == "" || password == "" {
		return user.User{}, errors.New("username and password are required")
	}
	result, err := us.repo.Login(nama)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return user.User{}, errors.New("data tidak ditemukan")
		}
		return user.User{}, errors.New("terjadi kesalahan pada sistem")
	}

	err = us.h.Compare(result.Password, password)

	if err != nil {
		return user.User{}, errors.New("password salah")
	}

	return result, nil
}
