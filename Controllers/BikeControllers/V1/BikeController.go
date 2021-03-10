package v1

import (
	"encoding/json"
	"net/http"
	acceptBll "redis-cache-api/Business/BikeBll/V1"
	acceptModel "redis-cache-api/Business/BikeBll/V1/Models"
	helper "redis-cache-api/Helper"
	constants "redis-cache-api/Helper/Constants"

	"github.com/gorilla/mux"
)

// GetAcceptTask func(w ResponseWriter, r *Request)
func GetAcceptTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msg := constants.AcceptTaskSuccess
	result, notFound := acceptBll.StampTask(vars["taskId"])
	if result {
		result = acceptBll.GetTask(vars["taskId"])
		helper.InitialResponseSuccess(w, result, http.StatusOK, msg)
	} else {
		if notFound {
			helper.InitialResponseError(w, http.StatusNotFound, constants.NotFoundData, result)
		} else {
			msg = constants.AcceptTaskFail
			helper.InitialResponseSuccess(w, result, http.StatusOK, msg)
		}
	}
}

// PostAcceptTask func(w ResponseWriter, r *Request)
func PostAcceptTask(w http.ResponseWriter, r *http.Request) {
	var post acceptModel.AcceptTaskRequest
	_ = json.NewDecoder(r.Body).Decode(&post)
	helper.InitialResponseSuccess(w, acceptBll.AcceptTask(post), http.StatusCreated)
}

// FetchData func(w ResponseWriter, r *Request)
func FetchData(w http.ResponseWriter, r *http.Request) {
	acceptBll.FetchData()
	helper.InitialResponseSuccess(w, "FetchData", http.StatusOK)
}

// VerifyTask func(w ResponseWriter, r *Request)
func VerifyTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := acceptBll.VerifyTask(vars["taskId"])
	msg := constants.NotFoundData
	if result {
		msg = constants.SuccessMessage
	}
	helper.InitialResponseSuccess(w, result, http.StatusOK, msg)
}
