package service_test

import (
	"Day25/features/user"
	"Day25/features/user/mocks"
	"Day25/features/user/service"
	eMock "Day25/helper/enkrip/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	enkrip := eMock.NewHashInterface(t)
	s := service.New(repo, enkrip)
	var inputData = user.User{Nama: "jerry", Password: "altamantul123"}
	var repoData = user.User{Nama: "jerry", Password: "some string"}
	var successReturnData = user.User{ID: uint(1), Nama: "jerry"}
	var falseData = user.User{}
	t.Run("Success Case", func(t *testing.T) {
		enkrip.On("HashPassword", inputData.Password).Return("some string", nil).Once()
		repo.On("Insert", repoData).Return(successReturnData, nil).Once()
		res, err := s.Register(inputData)

		enkrip.AssertExpectations(t)
		repo.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), res.ID)
		assert.Equal(t, "", res.Password)

	})

	t.Run("Failed Case", func(t *testing.T) {
		s = service.New(&FalseMockRepository{}, &MockEnkrip{})
		res, err := s.Register(falseData)
		assert.Error(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", res.Nama)
	})
}

func TestLogin_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	hashMock := eMock.NewHashInterface(t)

	userService := service.New(repo, hashMock)

	inputName := "testuser"
	inputPassword := "testpassword"

	repo.On("Login", inputName).Return(user.User{
		ID:       1,
		Nama:     inputName,
		Password: "hashedpassword",
	}, nil).Once()

	hashMock.On("Compare", "hashedpassword", inputPassword).Return(nil).Once()

	result, err := userService.Login(inputName, inputPassword)

	repo.AssertExpectations(t)
	hashMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, inputName, result.Nama)
}

func TestLogin_EmptyUsernameOrPassword(t *testing.T) {
	repo := mocks.NewRepository(t)
	hashMock := eMock.NewHashInterface(t)

	userService := service.New(repo, hashMock)

	result, err := userService.Login("", "")

	assert.Error(t, err)
	assert.Equal(t, user.User{}, result)
	assert.Equal(t, "username and password are required", err.Error())
}

func TestLogin_UserNotFound(t *testing.T) {
	repo := mocks.NewRepository(t)
	hashMock := eMock.NewHashInterface(t)

	userService := service.New(repo, hashMock)

	inputName := "data tidak ditemukan"
	inputPassword := "testpassword"

	repo.On("Login", inputName).Return(user.User{}, errors.New("data tidak ditemukan")).Once()

	result, err := userService.Login(inputName, inputPassword)

	repo.AssertExpectations(t)
	hashMock.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, user.User{}, result)
	assert.Equal(t, "terjadi kesalahan pada sistem", err.Error())
}

func TestLogin_IncorrectPassword(t *testing.T) {
	repo := mocks.NewRepository(t)
	hashMock := eMock.NewHashInterface(t)

	userService := service.New(repo, hashMock)

	inputName := "testuser"
	inputPassword := "password salah"

	repo.On("Login", inputName).Return(user.User{
		ID:       1,
		Nama:     inputName,
		Password: "hashedpassword",
	}, nil).Once()

	hashMock.On("Compare", "hashedpassword", inputPassword).Return(errors.New("password salah")).Once()

	result, err := userService.Login(inputName, inputPassword)

	repo.AssertExpectations(t)
	hashMock.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, user.User{}, result)
	assert.Equal(t, "password salah", err.Error())
}

type MockRepository struct{}

func (mr *MockRepository) Insert(newUser user.User) (user.User, error) {
	return user.User{ID: uint(1), Nama: "jerry"}, nil
}
func (mr *MockRepository) Login(nama string) (user.User, error) {
	return user.User{}, nil
}

type FalseMockRepository struct{}

func (fmr *FalseMockRepository) Insert(newUser user.User) (user.User, error) {
	return user.User{}, errors.New("something happend")
}
func (fmr *FalseMockRepository) Login(nama string) (user.User, error) {
	return user.User{}, nil
}

type MockEnkrip struct{}

func (me *MockEnkrip) Compare(hashed string, input string) error {
	return nil
}
func (me *MockEnkrip) HashPassword(input string) (string, error) {
	return "some string", nil
}
