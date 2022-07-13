package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentService()

	student, err := service.CreateStudent(&models.Student{
		Name:        r.FormValue("name"),
		Address:     r.FormValue("address"),
		DOB:         r.FormValue("dob"),
		Father_Name: r.FormValue("fathername"),
		Mother_Name: r.FormValue("mothername"),
		Status:      r.FormValue("status"),
	})

	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    student,
	})
}

func GetsStudent(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentService()
	student, err := service.GetsStudent()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get's student id",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "get's student id",
		Data:    student,
	})

}

func GetStudent(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(*models.Login)

	if user.Role.Role != string(models.AdminUser) && user.Role.Role != string(models.TeacherUser) {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    "User Type " + user.Role.Role + " not allowed",
		})
		return
	}

	service := service.NewStudentService()

	student, err := service.GetStudent(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    student,
	})
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	student, err := service.UpdateStudent(&models.Student{
		ID:          uint(a),
		Name:        r.FormValue("name"),
		Address:     r.FormValue("address"),
		DOB:         r.FormValue("dob"),
		Father_Name: r.FormValue("father_Name"),
		Mother_Name: r.FormValue("mother_Name"),
		Status:      r.FormValue("status"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get's student id",
			Data:    nil,
		})
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "can't update student",
		Data:    student,
	})
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentService()
	params := mux.Vars(r)
	student, err := service.DeleteStudent(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can delete student ",
			Data:    nil,
		})
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Student deleted",
		Data:    student,
	})
}
