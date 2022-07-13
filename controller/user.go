package controller

import (
	"errors"
	"fmt"
	"net/http"
	"school/models"
	"school/pkg"
	"school/service"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("password") != r.FormValue("confirmPassword") {
		errors.New("password did not matched")
	}

	service := service.NewUserService()
	User, err := service.CreateUser(&models.User{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
		UserType: r.FormValue("usertype"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "unable to log in",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Success",
		Data:    User,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewUserService()
	User, err := service.DeleteUser(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete User",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Success",
		Data:    User,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewUserService()
	User, err := service.GetUser(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get details",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Success",
		Data:    User,
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	service := service.NewUserService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	User, err := service.UpdateUser(&models.User{
		ID:       uint(a),
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update User details",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Success",
		Data:    User,
	})
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	service := service.NewUserService()
	user, err := service.UserLogin(
		r.FormValue("username"),
		r.FormValue("password"),
	)
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't login",
			Data:    err.Error(),
		})
		return
	}

	token, err := service.GenerateToken(user.ID)
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "unable to generate token",
			Data:    err.Error(),
		})
		return
	}

	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    token,
	})
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	user, err := service.NewUserService().GetUserByModel(&models.User{
		UserName: username,
	})

	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    err,
		})
	}

	resetKey := pkg.RandomString()

	service.NewResetPasswordService().SaveRestKey(&models.ResetPassword{
		UserID:    user.ID,
		ResetKey:  resetKey,
		ValidTill: time.Now().Add(time.Hour + 12),
		IsUsed:    false,
	})

	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    "Http://localhost:8000/reset_password/" + resetKey,
	})
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("password") != r.FormValue("conformpassword") {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    "password mismatch",
		})
	}

	params := mux.Vars(r)
	fmt.Println(params["resetKey"])
	userService := service.NewUserService()
	userReset, err := userService.GetResetKey(params["resetKey"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    "key not found or expired",
		})
		return
	}
	userService.UpdateResetKey(&models.ResetPassword{
		ID: userReset.ID,
	})
	_, err = userService.UpdateUser(&models.User{
		ID:       userReset.UserID,
		Password: r.FormValue("password"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    "unable to update password",
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Success",
		Data:    nil,
	})
}
