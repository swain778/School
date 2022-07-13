package controller

import (
	"fmt"
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateFees(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("studentID"))
	b, _ := strconv.Atoi(r.FormValue("classID"))
	d, _ := strconv.Atoi(r.FormValue("totalfees"))
	c, _ := strconv.Atoi(r.FormValue("feespaid"))

	fmt.Println(c)
	fmt.Println(d)
	e := d - c
	fmt.Println(e)
	service := service.NewFeeService()
	fees, err := service.CreateFees(&models.Fees{
		StudentID: uint(a),
		ClassID:   uint(b),
		Session:   r.FormValue("session"),
		FeesPaid:  uint(c),
		TotalFees: uint(d),
		Pending:   uint(e),
		Month:     r.FormValue("month"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create fees",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Fees Created",
		Data:    fees,
	})
}

func DeletedFees(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewExamService()
	exam, err := service.DeleteExam(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "exam deleted",
		Data:    exam,
	})
}

func GetFees(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewFeeService()
	fees, err := service.GetFees(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get fees",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    fees,
	})
}

func GetsFees(w http.ResponseWriter, r *http.Request) {
	service := service.NewFeeService()
	fees, err := service.GetsFees()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get fees",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    fees,
	})
}

func UpdateFees(w http.ResponseWriter, r *http.Request) {
	service := service.NewFeeService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("studentID"))
	c, _ := strconv.Atoi(r.FormValue("classID"))
	d, _ := strconv.Atoi(r.FormValue("feespaid"))
	e, _ := strconv.Atoi(r.FormValue("totalfees"))

	fees, err := service.UpdateFees(&models.Fees{
		ID:        uint(a),
		StudentID: uint(b),
		ClassID:   uint(c),
		Session:   r.FormValue("session"),
		FeesPaid:  uint(d),
		TotalFees: uint(e),
		Pending:   uint(d - e),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't updates fees",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    fees,
	})
}
