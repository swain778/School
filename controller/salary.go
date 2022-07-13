package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/go-chi/chi"
)

func CreateSalary(w http.ResponseWriter, r *http.Request) {

	a, _ := strconv.Atoi(r.FormValue("basic"))
	b, _ := strconv.Atoi(r.FormValue("bonus"))
	c, _ := strconv.Atoi(r.FormValue("pf"))
	d, _ := strconv.Atoi(r.FormValue("tax"))
	e, _ := strconv.Atoi(r.FormValue("netamount"))
	f, _ := strconv.Atoi(r.FormValue("staffID"))
	g, _ := strconv.Atoi(r.FormValue("bankdetailID"))

	service := service.NewSalaryService()
	salary, err := service.CreateSalaryService(&models.Salary{
		Basic:        a,
		Bonus:        b,
		PF:           c,
		Tax:          d,
		Month:        r.FormValue("month"),
		NetAmount:    e,
		StaffID:      uint(f),
		BankDetailID: uint(g),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "Can't create Salary",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "salary Created",
		Data:    salary,
	})
}

func DeleteSalary(w http.ResponseWriter, r *http.Request) {
	service := service.NewSalaryService()
	salary, err := service.DeleteSalaryService(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete salary",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "salary deleted",
		Data:    salary,
	})
}
