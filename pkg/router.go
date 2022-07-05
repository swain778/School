package pkg

import (
	"school/controller"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/student", controller.AddStudent).Methods("POST")
	r.HandleFunc("/student/{id}/delete", controller.DeleteStudent).Methods("GET")
	// r.HandleFunc("/student/{id}/update", controller.UpdateStudent).Methods("POST")
	r.HandleFunc("/student/{id}", controller.GetStudent).Methods("GET")
	r.HandleFunc("/students", controller.GetsStudent).Methods("GET")
	r.HandleFunc("/student/{id}/update", controller.UpdateStudent).Methods("POST")

	r.HandleFunc("/teacher", controller.AddTeacher).Methods("POST")
	r.HandleFunc("/teacher/{id}/delete", controller.DeleteTeacher).Methods("GET")
	r.HandleFunc("/teacher/{id}", controller.GetTeacher).Methods("GET")
	r.HandleFunc("/teachers", controller.GetsTeacher).Methods("GET")
	r.HandleFunc("/teacher/{id}/update", controller.UpdateTeacher).Methods("POST")

	r.HandleFunc("/class", controller.AddClass).Methods("POST")
	r.HandleFunc("/class/{id}/delete", controller.DeleteClass).Methods("GET")
	r.HandleFunc("/class/{id}", controller.GetClass).Methods("GET")
	r.HandleFunc("/classes", controller.GetClasses).Methods("GET")
	r.HandleFunc("/class/{id}/update", controller.UpdateClass).Methods("POST")

	r.HandleFunc("/complain", controller.AddComplain).Methods("POST")
	r.HandleFunc("/complain/{id}/delete", controller.DeleteComplain).Methods("GET")
	r.HandleFunc("/complain/{id}", controller.GetComplain).Methods("GET")
	r.HandleFunc("/complains", controller.GetsComplain).Methods("GET")
	r.HandleFunc("/complain/{id}/update", controller.UpdateComplain).Methods("POST")

	r.HandleFunc("/fees", controller.CreateFees).Methods("POST")
	r.HandleFunc("/fees/{id}/delete", controller.DeletedFees).Methods("GET")
	r.HandleFunc("/fees/{id}", controller.GetFees).Methods("GET")
	r.HandleFunc("/feess", controller.GetsFees).Methods("GET")
	r.HandleFunc("/fees/{id}/update", controller.UpdateFees).Methods("POST")

	r.HandleFunc("/exam", controller.CreateExam).Methods("POST")
	r.HandleFunc("/exam/{id}/delete", controller.DeleteExam).Methods("GET")
	r.HandleFunc("/exam/{id}", controller.GetExam).Methods("GET")
	r.HandleFunc("/exams", controller.GetsExam).Methods("GET")
	r.HandleFunc("/exam/{id}/update", controller.UpdateExam).Methods("POST")

	r.HandleFunc("/studentattendence", controller.CreateAttendence).Methods("POST")
	r.HandleFunc("/studentattendence/{id}/delete", controller.DeleteTeacherAttendence).Methods("GET")
	r.HandleFunc("/studentattendence/{id}", controller.GetStudentAttendence).Methods("GET")
	r.HandleFunc("/studentattendences", controller.GetsStudentAttendence).Methods("GET")
	r.HandleFunc("/studentattendence/{id}/update", controller.UpdateStudentAttendence).Methods("POST")

	r.HandleFunc("/studentsession", controller.CreateStudentSession).Methods("POST")
	r.HandleFunc("/studentsession/{id}/delete", controller.DeleteStudentSession).Methods("GET")
	r.HandleFunc("/studentsession/{id}", controller.GetStudentSession).Methods("GET")
	r.HandleFunc("/studentssessions", controller.GetsStudentSession).Methods("GET")
	r.HandleFunc("/studentsession/{id}/update", controller.UpdateStudentSession).Methods("POST")

	r.HandleFunc("/subject", controller.AddSubject).Methods("POST").Methods("POST")
	r.HandleFunc("/subject/{id}/delete", controller.DeleteStudent).Methods("GET")
	r.HandleFunc("/subject/{id}", controller.GetSubject).Methods("GET")
	r.HandleFunc("/subjects", controller.GetsSubject).Methods("GET")
	r.HandleFunc("/subject/{id}/update", controller.UpdateSubject).Methods("POST")

	r.HandleFunc("/teacherattendence", controller.TeacherAttendence).Methods("POST")
	r.HandleFunc("/teacherattendence/{id}/delete", controller.DeleteTeacherAttendence).Methods("GET")
	r.HandleFunc("/teacherattendence/{id}", controller.GetTeacherAttendence).Methods("GET")
	r.HandleFunc("/teacherattendences", controller.GetsTeacherAttendence).Methods("GET")
	r.HandleFunc("/teacherattendence/{id}/update", controller.UpdateTeacherAttendence).Methods("POST")

	r.HandleFunc("/classteacher", controller.ClassTeacher).Methods("POST")
	r.HandleFunc("/classteacher/{id}/delete", controller.DeleteClassTeacher).Methods("GET")
	r.HandleFunc("/classteacher/{id}", controller.GetClassTeacher).Methods("GET")
	r.HandleFunc("/classteachers", controller.GetsClassTeacher).Methods("GET")
	r.HandleFunc("/classteacher/{id}/update", controller.UpdateClassTeacher).Methods("POST")

	r.HandleFunc("/studentexam", controller.StudentExam).Methods("POST")
	r.HandleFunc("/studentexam/{id}/delete", controller.DeleteStudentExam).Methods("GET")
	r.HandleFunc("/studentexam/{id}", controller.StudentExam).Methods("GET")
	r.HandleFunc("/studentexams", controller.GetsStudentExam).Methods("GET")
	r.HandleFunc("/studentexam/{id}/update", controller.UpdateStudentExam).Methods("POST")

	r.HandleFunc("/studentclass", controller.StudentClass).Methods("POST")
	r.HandleFunc("/studentclass/{id}delete", controller.DeleteStudentClass).Methods("GET")
	r.HandleFunc("/studentclass/{id}", controller.GetStudentClass).Methods("GET")
	r.HandleFunc("/studentclasses", controller.GetsStudentClass).Methods("GET")
	r.HandleFunc("/studentclass/{id}/update", controller.UpdateStudentClass).Methods("POST")
}
