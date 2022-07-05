package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func ClassTeacher(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("classID"))
	b, _ := strconv.Atoi(r.FormValue("teacherID"))
	service := service.NewClassTeacherService()
	classteacher, err := service.CreateClassTeacher(&models.ClassTeacher{
		ClassID:   uint(a),
		TeacherID: uint(b),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create class teacher",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "class teacher created",
		Data:    classteacher,
	})
}

func DeleteClassTeacher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewClassTeacherService()
	classTeacher, err := service.DeleteClassTeacher(params["id"])
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
		Data:    classTeacher,
	})

}

func GetClassTeacher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewClassTeacherService()
	classTeacher, err := service.GetClassTeacher(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get class teacher",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "can't get class teacher",
		Data:    classTeacher,
	})
}

func GetsClassTeacher(w http.ResponseWriter, r *http.Request) {
	service := service.NewClassTeacherService()
	classTeacher, err := service.GetsClassTeacher()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get teacher",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    classTeacher,
	})
}

func UpdateClassTeacher(w http.ResponseWriter, r *http.Request) {
	service := service.NewClassTeacherService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("classID"))
	c, _ := strconv.Atoi(r.FormValue("teacherID"))
	classTeacher, err := service.UpdateClassTeacher(&models.ClassTeacher{
		ID:        uint(a),
		ClassID:   uint(b),
		TeacherID: uint(c),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update class teacher",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "class teacher updated",
		Data:    classTeacher,
	})
}
