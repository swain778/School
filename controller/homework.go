package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/go-chi/chi"
)

func CreateHomework(w http.ResponseWriter, r *http.Request) {
	service := service.NewHomeworkService()
	a, _ := strconv.Atoi(r.FormValue("subjectID"))
	b, _ := strconv.Atoi(r.FormValue("studentID"))
	c, _ := strconv.Atoi(r.FormValue("teacherID"))
	homework, err := service.CreateHomework(&models.Homework{
		SubjectID:      uint(a),
		StudentID:      uint(b),
		TeacherID:      uint(c),
		Class:          r.FormValue("class"),
		Homework:       r.FormValue("homework"),
		SubmissionDate: r.FormValue("submissiondate"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create homework",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Homework created",
		Data:    homework,
	})
}

func DeleteHomework(w http.ResponseWriter, r *http.Request) {
	service := service.NewHomeworkService()
	homework, err := service.DeleteHomework(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "Can't delete homework",
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
