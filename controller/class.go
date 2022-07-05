package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func AddClass(w http.ResponseWriter, r *http.Request) {
	service := service.NewClassService()
	class, err := service.CreateClass(&models.Class{
		Class: r.FormValue("class"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create class",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "class Created",
		Data:    class,
	})
}

func DeleteClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewClassService()
	class, err := service.DeleteClass(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete class",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "class deleted",
		Data:    class,
	})
}

func GetClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewClassService()
	class, err := service.GetClass(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get class",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "get class id",
		Data:    class,
	})
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	service := service.NewClassService()
	class, err := service.GetClasses()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get's class",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "get's class",
		Data:    class,
	})
}

func UpdateClass(w http.ResponseWriter, r *http.Request) {
	service := service.NewClassService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	class, err := service.UpdateClass(&models.Class{
		ID:    uint(a),
		Class: r.FormValue("class"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    class,
	})
}
