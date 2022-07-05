package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func AddComplain(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("studentID"))
	service := service.NewComplainService()
	complain, err := service.CreateComplain(&models.Complain{
		StudentID: uint(a),
		Complain:  r.FormValue("complain"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create complain",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "complain created",
		Data:    complain,
	})
}

func DeleteComplain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewComplainService()
	complain, err := service.DeleteComplain(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete complain",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "complain deleted",
		Data:    complain,
	})

}

func GetComplain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewComplainService()
	complain, err := service.GetComplain(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get complain ",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "get complain",
		Data:    complain,
	})

}

func GetsComplain(w http.ResponseWriter, r *http.Request) {
	service := service.NewComplainService()
	complain, err := service.GetsComplain()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get complain",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    complain,
	})
}

func UpdateComplain(w http.ResponseWriter, r *http.Request) {
	service := service.NewComplainService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("studentID"))

	complain, err := service.UpdateComplain(&models.Complain{
		ID:        uint(a),
		Complain:  r.FormValue("complain"),
		StudentID: uint(b),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update complain",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    complain,
	})
}
