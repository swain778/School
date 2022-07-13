package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateSessionYear(w http.ResponseWriter, r *http.Request) {
	service := service.NewSessionYear()
	SessionYear, err := service.CreateSessionYear(&models.SessionYear{
		SessionYear: r.FormValue("sessionyear"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create student session",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student session created",
		Data:    SessionYear,
	})
}

func DeleteSessionYear(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewSessionYear()
	SessionYear, err := service.DeleteSessionYear(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete student session",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student session deleted",
		Data:    SessionYear,
	})
}

func GetSessionYear(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewSessionYear()
	SessionYear, err := service.GetSessionYear(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "get student session",
		Data:    SessionYear,
	})
}

func GetsSessionYear(w http.ResponseWriter, r *http.Request) {
	//service := service.NewSessionYear()
	SessionYear, err := service.NewSessionYear().GetsSessionYear()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student session",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    SessionYear,
	})
}

func UpdateSessionYear(w http.ResponseWriter, r *http.Request) {
	service := service.NewSessionYear()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])

	SessionYear, err := service.UpdateSessionYear(&models.SessionYear{
		ID:          uint(a),
		SessionYear: r.FormValue("sessionyear"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update student session",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    SessionYear,
	})
}
