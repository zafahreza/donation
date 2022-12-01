package service

import (
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

func (service UserServiceImpl) Create(request client.UserCreateRequest) client.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	userEmail, err := service.UserRepository.FindByEmail(tx, request.Email)
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

	newUser := service.UserRepository.Save(tx, user)

	return helper.ToUserResponse(newUser)
}

func (service UserServiceImpl) Update(request client.UserUpdateRequest) client.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(tx, request.Id)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	userEmail, err := service.UserRepository.FindByEmail(tx, request.Email)
	helper.PanicIfError(err)
	exception.PanicIfEmailUsed(request.Email, userEmail.Email)

	goodEmail := strings.ToLower(request.Email)
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = goodEmail

	updatedUser := service.UserRepository.Update(tx, user)

	return helper.ToUserResponse(updatedUser)
}

func (service UserServiceImpl) Delete(userId int) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(tx, userId)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	service.UserRepository.Delete(tx, user)
}

func (service UserServiceImpl) Session(request client.UserSessionRequest) client.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(tx, request.Email)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) FindById(userId int) client.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(tx, userId)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) FindByEmail(userEmail string) client.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(tx, userEmail)
	helper.PanicIfError(err)
	exception.PanicIfNotFound(user.Id)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) FindAll() []client.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(tx)

	return helper.ToUserResponses(users)
}
