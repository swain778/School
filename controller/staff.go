package controller

import (
	"net/http"
	"school/models"
	"school/service"

	"github.com/go-chi/chi"
)

func CreateStaff(w http.ResponseWriter, r *http.Request) {
	service := service.NewStaffService()
	staff, err := service.CreateStaff(&models.Staff{
		StaffType:   r.FormValue("stafftype"),
		Name:        r.FormValue("name"),
		DOB:         r.FormValue("dob"),
		JoiningDate: r.FormValue("joiningdate"),
		Session:     r.FormValue("session"),
		Aadharno:    r.FormValue("aadharno"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create staff",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Successfully created Staff",
		Data:    staff,
	})
}

func DeleteStaff(w http.ResponseWriter, r *http.Request) {

	service := service.NewStaffService()
	staff, err := service.DeleteStaff(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete staff",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Successfully deleted staff",
		Data:    staff,
	})
}

func GetStaff(w http.ResponseWriter, r *http.Request) {

	service := service.NewStaffService()
	staff, err := service.GetStaff(chi.URLParam(r, "staffType"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get staff details",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Successfully get staf details",
		Data:    staff,
	})
}

func GetStaffByID(w http.ResponseWriter, r *http.Request) {
	service := service.NewStaffService()
	staff, err := service.GetStaffByID(chi.URLParam(r, "staffType"), chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get staf by id",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "successfully get staf by id",
		Data:    staff,
	})
}
