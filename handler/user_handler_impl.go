package handler

import (
	"donation/entity/client"
	"donation/helper"
	"donation/middleware"
	"donation/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserHandlerImpl struct {
	UserService    service.UserService
	AuthMiddleware middleware.AuthMidleware
}

func NewUserHanlder(userService service.UserService, authMiddleware middleware.AuthMidleware) UserHandler {
	return &UserHandlerImpl{
		UserService:    userService,
		AuthMiddleware: authMiddleware,
	}
}

func (handler *UserHandlerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userRequest := client.UserCreateRequest{}
	helper.ReadFromRequestBody(r, &userRequest)

	userResponse := handler.UserService.Create(r.Context(), userRequest)
	webResponse := client.UserAPIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (handler *UserHandlerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := handler.AuthMiddleware.ValidateToken(r)

	userRequest := client.UserUpdateRequest{}
	helper.ReadFromRequestBody(r, &userRequest)

	userRequest.Id = userId

	userResponse := handler.UserService.Update(r.Context(), userRequest)
	webResponse := client.UserAPIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (handler *UserHandlerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := handler.AuthMiddleware.ValidateToken(r)

	handler.UserService.Delete(r.Context(), userId)
	webResponse := client.UserAPIResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (handler *UserHandlerImpl) Session(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userRequest := client.UserSessionRequest{}
	helper.ReadFromRequestBody(r, &userRequest)

	userResponse := handler.UserService.Session(r.Context(), userRequest)
	webResponse := client.UserAPIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (handler *UserHandlerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("userId")
	userId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	userResponse := handler.UserService.FindById(r.Context(), userId)
	webResponse := client.UserAPIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (handler *UserHandlerImpl) FindByEmail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userEmail := params.ByName("userEmail")

	userResponse := handler.UserService.FindByEmail(r.Context(), userEmail)
	webResponse := client.UserAPIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (handler *UserHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userResponse := handler.UserService.FindAll(r.Context())
	webResponse := client.UserAPIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (handler *UserHandlerImpl) FindOtp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userRequest := client.UserOtpRequest{}
	helper.ReadFromRequestBody(r, &userRequest)

	userResponse := handler.UserService.FindOtp(r.Context(), userRequest)
	webResponse := client.UserAPIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (handler *UserHandlerImpl) GetNewOtp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userRequest := client.UserGetNewOtpRequest{}
	helper.ReadFromRequestBody(r, &userRequest)

	userResponse := handler.UserService.GetNewOtp(r.Context(), userRequest)
	webResponse := client.UserAPIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
