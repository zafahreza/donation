package exception

import (
	"donation/entity/client"
	"donation/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if unauthorizedError(w, r, err) {
		return
	}

	if notFoundError(w, r, err) {
		return
	}
	if validationEror(w, r, err) {
		return
	}
	if emailUsedError(w, r, err) {
		return
	}
	if wrongPasswordError(w, r, err) {
		return
	}
	if wrongOtpError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := client.UserAPIResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   nil,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := client.UserAPIResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false

}

func validationEror(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := client.UserAPIResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}

func emailUsedError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(EmailUsedError)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := client.UserAPIResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}

func wrongPasswordError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(WrongPasswordError)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := client.UserAPIResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}
func wrongOtpError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(WrongOtpError)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := client.UserAPIResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}

func unauthorizedError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := client.UserAPIResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}
