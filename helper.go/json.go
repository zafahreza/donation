package helper

import (
	"donation/entity/domain"
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "aplication/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func UnMarshal(result string) domain.User {
	var user domain.User
	err := json.Unmarshal([]byte(result), &user)
	PanicIfError(err)

	return user
}
