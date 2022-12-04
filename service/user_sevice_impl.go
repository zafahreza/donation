package service

import (
	"context"
	"donation/chache"
	"donation/entity/client"
	"donation/entity/domain"
	"donation/exception"
	"donation/helper"
	"donation/middleware"
	"donation/repository"
	"errors"
	"github.com/go-redis/redis/v9"
	mail "github.com/xhit/go-simple-mail/v2"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	UserChache     chache.UserChache
	DB             *gorm.DB
	Validate       *validator.Validate
	Smtp           *mail.SMTPClient
}

func NewUserService(userRepository repository.UserRepository, userChache chache.UserChache, DB *gorm.DB, validate *validator.Validate, smtp *mail.SMTPClient) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		UserChache:     userChache,
		DB:             DB,
		Validate:       validate,
		Smtp:           smtp,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request client.UserCreateRequest) client.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	userEmail, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	helper.PanicIfError(err)
	exception.PanicIfEmailUsed(request.Email, userEmail.Email)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	helper.PanicIfError(err)

	goodEmail := strings.ToLower(request.Email)
	user := domain.User{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Email:        goodEmail,
		Bio:          request.Bio,
		PasswordHash: string(passwordHash),
		IsActive:     false,
	}

	stringOtp := helper.GenerateOtp()

	otp := domain.OTP{
		Email: goodEmail,
		OTP:   stringOtp,
	}

	go service.UserChache.SetOtp(ctx, otp)
	go helper.SendOtp(otp, service.Smtp)

	newUser := service.UserRepository.Save(ctx, tx, user, otp)

	return helper.ToUserResponse(newUser)
}

func (service *UserServiceImpl) Update(ctx context.Context, request client.UserUpdateRequest) client.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	var user domain.User

	key := "userid" + strconv.Itoa(request.Id)
	user, err = service.UserChache.Get(ctx, key)
	if err == redis.Nil {
		user, err = service.UserRepository.FindById(ctx, tx, request.Id)
		helper.PanicIfError(err)
		exception.PanicIfNotFound(user.Id)
	}

	goodEmail := strings.ToLower(request.Email)

	if user.Email != goodEmail {
		userEmail, err := service.UserRepository.FindByEmail(ctx, tx, goodEmail)
		helper.PanicIfError(err)
		exception.PanicIfEmailUsed(goodEmail, userEmail.Email)
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = goodEmail
	user.Bio = request.Bio

	updatedUser := service.UserRepository.Update(ctx, tx, user)

	key2 := "userbyemail" + goodEmail
	go service.UserChache.Del(ctx, key, key2)

	return helper.ToUserResponse(updatedUser)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	var user domain.User

	key := "userid" + strconv.Itoa(userId)

	user, err := service.UserChache.Get(ctx, key)
	if err == redis.Nil {
		user, err = service.UserRepository.FindById(ctx, tx, userId)
		helper.PanicIfError(err)
		exception.PanicIfNotFound(user.Id)
	}

	service.UserRepository.Delete(ctx, tx, user)

	key2 := "userbyemail" + user.Email
	go service.UserChache.Del(ctx, key, key2)

}

func (service *UserServiceImpl) Session(ctx context.Context, request client.UserSessionRequest) client.UserLoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	var user domain.User

	key := "userbyemail" + request.Email

	user, err = service.UserChache.Get(ctx, key)
	if err == redis.Nil {
		user, err = service.UserRepository.FindByEmail(ctx, tx, request.Email)
		helper.PanicIfError(err)
		exception.PanicIfNotFound(user.Id)
		go service.UserChache.Set(ctx, user, key)
	}

	if user.IsActive == false {
		panic(exception.NewUnauthorizedError(errors.New("your account is not active, please activate")))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		panic(exception.NewWrongPasswordError(errors.New("wrong password")))
	}

	token := middleware.NewAuthMiddleware().GenerateToken(user.Id)

	return helper.ToUserLoginResponse(user, token)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) client.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	var user domain.User

	key := "userid" + strconv.Itoa(userId)

	user, err := service.UserChache.Get(ctx, key)
	if err == redis.Nil {
		user, err = service.UserRepository.FindById(ctx, tx, userId)
		helper.PanicIfError(err)
		exception.PanicIfNotFound(user.Id)
		go service.UserChache.Set(ctx, user, key)
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindByEmail(ctx context.Context, userEmail string) client.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	var user domain.User

	key := "userbyemail" + userEmail

	user, err := service.UserChache.Get(ctx, key)
	if err == redis.Nil {
		user, err = service.UserRepository.FindByEmail(ctx, tx, userEmail)
		helper.PanicIfError(err)
		exception.PanicIfNotFound(user.Id)
		go service.UserChache.Set(ctx, user, key)
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []client.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}

func (service *UserServiceImpl) FindOtp(ctx context.Context, request client.UserOtpRequest) client.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	goodEmail := strings.ToLower(request.Email)

	userOtp := domain.OTP{
		Email: goodEmail,
		OTP:   request.OTP,
	}

	otp, err := service.UserChache.GetOtp(ctx, userOtp)
	if err == redis.Nil {
		panic(exception.NewNotFoundError(errors.New("otp not found")))
	}

	if request.OTP != otp.OTP {
		panic(exception.NewWrongOtpError(errors.New("otp invalid")))
	}

	user := service.UserRepository.UpdateStatusEmail(ctx, tx, otp)

	key := "userid" + strconv.Itoa(user.Id)
	key1 := "userbyemail" + otp.Email
	key2 := "otpfor" + otp.Email
	go service.UserChache.Del(ctx, key, key1, key2)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) GetNewOtp(ctx context.Context, request client.UserGetNewOtpRequest) client.UserGetNewOtpResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	//cek permission

	goodEmail := strings.ToLower(request.Email)

	var user domain.User

	key := "userbyemail" + request.Email
	user, err = service.UserChache.Get(ctx, key)
	if err == redis.Nil {
		user, err = service.UserRepository.FindByEmail(ctx, tx, goodEmail)
		helper.PanicIfError(err)
		exception.PanicIfNotFound(user.Id)
	}

	if user.IsActive == true {
		panic(exception.NewWrongOtpError(errors.New("this user is already active")))
	}

	var otp domain.OTP

	stringOtp := helper.GenerateOtp()

	otp = domain.OTP{
		Email: goodEmail,
		OTP:   stringOtp,
	}

	go helper.SendOtp(otp, service.Smtp)
	go service.UserChache.SetOtp(ctx, otp)

	response := client.UserGetNewOtpResponse{
		Email: goodEmail,
		Msg:   "OTP sent successfully, check your email",
	}

	//set permission wait 5 menit

	return response
}
