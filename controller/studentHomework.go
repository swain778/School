package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/go-chi/chi"
)

func CreateStudentHomework(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentHomeWorkService()
	a, _ := strconv.Atoi(r.FormValue("homeworkID"))
	b, _ := strconv.Atoi(r.FormValue("studentID"))
	homework, err := service.CreateStudentHomework(&models.StudentHomework{
		HomeworkID: uint(a),
		StudentID:  uint(b),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "Can't create student homework",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "successfully created",
		Data:    homework,
	})
}

func DeleteStudentHomework(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentHomeWorkService()
	homework, err := service.DeleteStudentHomework(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete student homework",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "successfully deleted",
		Data:    homework,
	})
}
