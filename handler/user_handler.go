package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Session(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByEmail(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindOtp(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetNewOtp(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
