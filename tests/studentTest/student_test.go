package student

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"school/config"
	"school/models"
	"strconv"
	"testing"
)

type StudentResponse struct {
	Code    int
	Message string
	Data    models.Student
}

func TestCreateStudent(t *testing.T) {

	resp, err := http.PostForm(
		config.ApiURL+"student",
		url.Values{
			"name":       {"Shuham"},
			"address":    {"#home"},
			"dob":        {"2020-12-03"},
			"fathername": {"Papa"},
			"mothername": {"Mamma"},
		},
	)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	Student := StudentResponse{}
	json.Unmarshal(body, &Student)

	if Student.Code != 200 {
		t.Errorf("Status code is %d", Student.Code)
	}
	if Student.Data.ID == 0 {
		log.Fatalf("Error while creating student %v", err)
	}
}

func TestSudentDelete(t *testing.T) {
	student := &StudentResponse{}

	str := strconv.Itoa(int(student.Data.ID))
	resp, err := http.Get(config.ApiURL + "student/" + str + "/delete")
	if err != nil {
		t.Error(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &student)
	fmt.Println(config.ApiURL + "student" + str + "delete")
	if student.Code != 200 {
		t.Errorf("Status code is %d", student.Code)
	}
}

func TestParallel(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", TestCreateStudent)
		t.Run("Test2", TestSudentDelete)
	})
}
