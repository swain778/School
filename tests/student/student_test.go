package student

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"school/config"
	"school/models"
	"testing"
)

type StudentResponse struct {
	Code    int
	Message string
	Data    models.Student
}

func TestSudentAdd(t *testing.T) {
	resp, err := http.Get(config.ApiURL + "student/17")
	resp.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	student := StudentResponse{}
	json.Unmarshal(body, &student)

	if student.Code != 200 {
		t.Errorf("Status code is %d", student.Code)
	}
}
