package services

// import (
// 	"context"
// 	"testing"

// 	"github.com/duyledat197/interview-hao/models"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type userLoginRequest struct {
// 	Email    string
// 	Password string
// }

// func TestRegister_Service_InvalidEmail(t *testing.T) {
// 	var userRegister models.User = models.User{
// 		Email: "test002",
// 	}
// 	userRepo := &mockUserRepository{}
// 	service := NewService(userRepo)
// 	err := service.Register(&userRegister, "password")
// 	assert.Equal(t, models.ErrInvalidEmail, err)
// }

// func TestRegister_Service_InvalidPassword(t *testing.T) {
// 	var userRegister models.User = models.User{
// 		Email: "test002@yopmail.com",
// 	}
// 	userRepo := &mockUserRepository{}
// 	service := NewService(userRepo)
// 	err := service.Register(&userRegister, "")
// 	assert.Equal(t, models.ErrInvalidPassword, err)
// }

// func TestRegister_Service_ExistEmail(t *testing.T) {
// 	var userRegister models.User = models.User{
// 		Email: "test001@yopmail.com",
// 	}
// 	userRepo := &mockUserRepository{}
// 	service := NewService(userRepo)
// 	userRepo.On("Create", &userRegister).Return(models.ErrUserAlreadyExist, nil)
// 	userRepo.On("FindByEmail", userRegister.Email).Return(&userRegister, nil, nil)
// 	err := service.Register(&userRegister, "password")
// 	assert.Equal(t, models.ErrUserAlreadyExist, err)
// }

// func TestRegister_Service_Success(t *testing.T) {
// 	var userRegister models.User = models.User{
// 		Email: "test002@yopmail.com",
// 	}
// 	userRepo := &mockUserRepository{}
// 	service := NewService(userRepo)
// 	userRepo.On("Create", &userRegister).Return(nil, nil)
// 	userRepo.On("FindByEmail", userRegister.Email).Return(nil, models.ErrUnknowUser, nil)
// 	err := service.Register(&userRegister, "password")
// 	assert.Equal(t, nil, err)
// }

// func TestLogin_Service_InvalidEmail(t *testing.T) {
// 	var userLogin userLoginRequest = userLoginRequest{
// 		Email:    "test001yopmail.com",
// 		Password: "Test@123",
// 	}
// 	userRepo := &mockUserRepository{}
// 	service := NewService(userRepo)
// 	user, token, err := service.Login(userLogin.Email, userLogin.Password)
// 	assert.Equal(t, models.ErrInvalidEmail, err)
// 	assert.Equal(t, models.User{}, user)
// 	assert.Equal(t, "", token)
// }

// func TestLogin_Service_EmptyEmail(t *testing.T) {
// 	var userLogin userLoginRequest = userLoginRequest{
// 		Email:    "",
// 		Password: "Test@123",
// 	}
// 	userRepo := &mockUserRepository{}
// 	service := NewService(userRepo)
// 	user, token, err := service.Login(userLogin.Email, userLogin.Password)
// 	assert.Equal(t, models.ErrInvalidEmail, err)
// 	assert.Equal(t, models.User{}, user)
// 	assert.Equal(t, "", token)
// }

// func TestLogin_Service_NotExistEmail(t *testing.T) {
// 	var userLogin userLoginRequest = userLoginRequest{
// 		Email:    "test001@yopmail.com",
// 		Password: "Test@123",
// 	}
// 	userRepo := &mockUserRepository{}
// 	service := NewService(userRepo)

// 	userRepo.On("FindByEmail", userLogin.Email).Return(nil, models.ErrUnknowUser, nil)
// 	user, token, err := service.Login(userLogin.Email, userLogin.Password)
// 	assert.Equal(t, models.ErrUnknowUser, err)
// 	assert.Equal(t, "", token)
// 	assert.Equal(t, models.User{}, user)
// }

// func TestLogin_Service_WrongPassword(t *testing.T) {
// 	var userLogin userLoginRequest = userLoginRequest{
// 		Email:    "test001@yopmail.com",
// 		Password: "wrongpassword",
// 	}
// 	rightPassword := "Test@123"
// 	hashPassword, err := models.HashPassword(rightPassword)
// 	expectUser := models.User{
// 		Email:          userLogin.Email,
// 		HashedPassword: hashPassword,
// 	}
// 	userRepo := &mockUserRepository{}
// 	service := NewService(userRepo)
// 	userRepo.On("FindByEmail", userLogin.Email).Return(&expectUser, nil, nil)
// 	user, token, err := service.Login(userLogin.Email, userLogin.Password)
// 	assert.Equal(t, models.ErrWrongEmailOrPassword, err)
// 	assert.Equal(t, "", token)
// 	assert.Equal(t, models.User{}, user)
// }

// func TestLogin_Service_Success(t *testing.T) {
// 	var userLogin userLoginRequest = userLoginRequest{
// 		Email:    "test001@yopmail.com",
// 		Password: "Test@123",
// 	}
// 	hashPassword, err := models.HashPassword(userLogin.Password)
// 	expectUser := models.User{
// 		Email:          userLogin.Email,
// 		HashedPassword: hashPassword,
// 	}
// 	userRepo := &mockUserRepository{}
// 	service := NewService(userRepo)

// 	userRepo.On("FindByEmail", userLogin.Email).Return(&expectUser, nil, nil)
// 	user, token, err := service.Login(userLogin.Email, userLogin.Password)
// 	assert.Equal(t, nil, err)
// 	assert.NotEqual(t, "", token)
// 	assert.NotEqual(t, models.User{}, user)
// }

// // mock
// type mockUserRepository struct {
// 	mock.Mock
// }

// func (m *mockUserRepository) Create(ctx context.Context, user *models.User) error {
// 	args := m.Called(user)
// 	return args.Error(1)
// }

// func (m *mockUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
// 	args := m.Called(email)
// 	var r0 *models.User
// 	if rf, ok := args.Get(0).(func(string) *models.User); ok {
// 		r0 = rf(email)
// 	} else {
// 		if args.Get(0) != nil {
// 			r0 = args.Get(0).(*models.User)
// 		}
// 	}

// 	var r1 error
// 	if rf, ok := args.Get(1).(func(string) error); ok {
// 		r1 = rf(email)
// 	} else {
// 		r1 = args.Error(1)
// 	}

// 	return r0, r1
// }
// func (m *mockUserRepository) FindByUserID(ctx context.Context, userID string) (*models.User, error) {
// 	return nil, nil
// }
// func (m *mockUserRepository) FindAll() (ctx context.Context,[]*models.User, error) {
// 	return []*models.User{}, nil
// }
// func (m *mockUserRepository) Update(ctx context.Context,userID string, user *models.User) error {
// 	return nil
// }
// func (m *mockUserRepository) Delete(ctx context.Context,userID string) error {
// 	return nil
// }
