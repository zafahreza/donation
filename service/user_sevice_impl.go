package service

import (
	"context"
	"donation/entity/client"
	"donation/entity/domain"
	"donation/exception"
	"donation/helper.go"
	"donation/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service UserServiceImpl) Create(ctx context.Context, request client.UserCreateRequest) client.UserResponse {
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
		PasswordHash: string(passwordHash),
	}

	newUser := service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(newUser)
}

func (service UserServiceImpl) Update(ctx context.Context, request client.UserUpdateRequest) client.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	goodEmail := strings.ToLower(request.Email)

	if user.Email != goodEmail {
		userEmail, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
		helper.PanicIfError(err)
		exception.PanicIfEmailUsed(request.Email, userEmail.Email)
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = goodEmail

	updatedUser := service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(updatedUser)
}

func (service UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	service.UserRepository.Delete(ctx, tx, user)
}

func (service UserServiceImpl) Session(ctx context.Context, request client.UserSessionRequest) client.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) FindById(ctx context.Context, userId int) client.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) FindByEmail(ctx context.Context, userEmail string) client.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, userEmail)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) FindAll(ctx context.Context) []client.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}
