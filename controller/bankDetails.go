package controller

import (
	"log"
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/go-chi/chi"
)

func CreateBankDetail(w http.ResponseWriter, r *http.Request) {
	service := service.NewBankDetailervice()

	staffID, _ := strconv.Atoi(chi.URLParam(r, "staff_id"))
	bankAccountData, err := service.GetsBankDetail(uint(staffID))
	if err != nil {
		log.Fatalln(err)
	}
	isDefault := true
	if len(*bankAccountData) > 0 {
		isDefault = false
	}

	b, _ := strconv.Atoi(r.FormValue("staffID"))
	bank, err := service.CreateBankDetail(&models.BankDetail{
		StaffID:     uint(b),
		Name:        r.FormValue("name"),
		Bank:        r.FormValue("bank"),
		BankAccount: r.FormValue("bankaccount"),
		IFSC:        r.FormValue("ifsc"),
		BranchCode:  r.FormValue("branchcode"),
		IsDefault:   isDefault,
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create bank",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Successfullt created",
		Data:    bank,
	})
}

func DeleteBankDetail(w http.ResponseWriter, r *http.Request) {
	service := service.NewBankDetailervice()
	bank, err := service.DeleteBankDetail(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete bank details",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Successfullt deleted data",
		Data:    bank,
	})
}
