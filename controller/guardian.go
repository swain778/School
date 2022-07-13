package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/go-chi/chi"
)

func CreateGuardian(w http.ResponseWriter, r *http.Request) {
	service := service.NewGuardianService()
	a, _ := strconv.Atoi(r.FormValue("studentID"))
	guardian, err := service.CreateGuardian(&models.Guardian{
		GuardianType: r.FormValue("guardiantype"),
		GuardianName: r.FormValue("guardianname"),
		StudentID:    uint(a),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create guardian",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Guardian created",
		Data:    guardian,
	})
}

func DeleteGuardian(w http.ResponseWriter, r *http.Request) {
	service := service.NewGuardianService()
	guardian, err := service.DeleteGuardian(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete guardian",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Guardian deleted",
		Data:    guardian,
	})
}
