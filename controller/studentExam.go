package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
)

func StudentExam(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("studentID"))
	b, _ := strconv.Atoi(r.FormValue("examID"))
	c, _ := strconv.Atoi(r.FormValue("sessionyearID"))
	d, _ := strconv.Atoi(r.FormValue("internalmarks"))
	e, _ := strconv.Atoi(r.FormValue("theorymarks"))
	f, _ := strconv.Atoi(r.FormValue("practicalmarks"))
	g, _ := strconv.Atoi(r.FormValue("maximummarks"))

	service := service.NewStudentExamService()
	studentExam, err := service.SaveStudentExamInfo(&models.StudentExam{
		StudentID:      uint(a),
		ExamID:         uint(b),
		SessionYearID:  uint(c),
		InternalMarks:  uint(d),
		TheoryMarks:    uint(e),
		PracticalMarks: uint(f),
		MaximumMarks:   uint(g),
		ExamType:       r.FormValue("examtype"),
	})

	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student exam created",
		Data:    studentExam,
	})
}

func DeleteStudentExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewStudentExamService()
	studentExam, err := service.DeleteStudentExam(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student exam deleted",
		Data:    studentExam,
	})

}

func GetStudentExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewStudentExamService()
	studentExam, err := service.GetStudentExam(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentExam,
	})
}

func GetsStudentExam(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentExamService()
	studentExam, err := service.GetsStudentExam()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentExam,
	})
}

func UpdateStudentExam(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentExamService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("studentID"))
	c, _ := strconv.Atoi(r.FormValue("examID"))
	d, _ := strconv.Atoi(r.FormValue("sessionyearID"))

	studentExam, err := service.UpdateStudentExam(&models.StudentExam{
		ID:            uint(a),
		StudentID:     uint(b),
		ExamID:        uint(c),
		SessionYearID: uint(d),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "CAn't update student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentExam,
	})
}

func GetStudentExamByID(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentExamService()
	exam, err := service.GetStudentExamByID(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    200,
			Message: "can't get details",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    900,
		Message: "get student data",
		Data:    exam,
	})
}
