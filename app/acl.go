package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"school/controller"
	"school/models"
	"school/pkg"

	"github.com/go-chi/chi"
)

type acl struct {
	Admin   []string `json:"admin"`
	Teacher []string `json:"teacher"`
	Student []string `json:"student"`
}

func Acl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.Login)
		aclData := &acl{}
		aclFile, err := os.Open("acl.json")
		aclBytes, _ := ioutil.ReadAll(aclFile)
		json.Unmarshal(aclBytes, &aclData)

		if err != nil {
			fmt.Println(err)
		}
		defer aclFile.Close()

		acl := []string{}
		if user.Role.Role == string(models.AdminUser) {
			acl = aclData.Admin
		} else if user.Role.Role == string(models.TeacherUser) {
			acl = aclData.Teacher
		} else if user.Role.Role == string(models.StudentUser) {
			acl = aclData.Student
		}

		next.ServeHTTP(w, r.WithContext(r.Context()))
		route := chi.RouteContext(r.Context()).RoutePattern()
		if !pkg.Contains(acl, route) {
			controller.ApiResponse(w, &controller.Res{
				Code:    900,
				Message: "error",
				Data:    "Api not allowed to this user role",
			})
		}
	})
}
