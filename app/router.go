package app

import (
	"school/controller"

	"github.com/go-chi/chi"
)

func (ctn *Container) LoadRoutes() {
	r := ctn.Mux
	r.Route("/user", func(r chi.Router) {
		r.Post("/signUp/", controller.CreateUser)
		r.Post("/login", controller.UserLogin)
	})

	r.With(Auth, Acl).Route("/api", func(r chi.Router) {
		r.Get("/student/{id}", controller.GetStudent)
		r.Post("/student", controller.AddStudent)
		r.Get("/student/{id}/delete", controller.DeleteStudent)
		r.Get("/students", controller.GetsStudent)
		r.Post("/student/{id}/update", controller.UpdateStudent)

		r.Post("/teacher", controller.AddTeacher)
		r.Get("/teacher/{id}/delete", controller.DeleteTeacher)
		r.Get("/teacher/{id}", controller.GetTeacher)
		r.Get("/teachers", controller.GetsTeacher)
		r.Post("/teacher/{id}/update", controller.UpdateTeacher)

		r.Post("/class", controller.AddClass)
		r.Get("/class/{id}/delete", controller.DeleteClass)
		r.Get("/class/{id}", controller.GetClass)
		r.Get("/classes", controller.GetClasses)
		r.Post("/class/{id}/update", controller.UpdateClass)

		r.Post("/complain", controller.AddComplain)
		r.Get("/complain/{id}/delete", controller.DeleteComplain)
		r.Get("/complain/{id}", controller.GetComplain)
		r.Get("/complains", controller.GetsComplain)
		r.Post("/complain/{id}/update", controller.UpdateComplain)

		r.Post("/fees", controller.CreateFees)
		r.Get("/fees/{id}/delete", controller.DeletedFees)
		r.Get("/fees/{id}", controller.GetFees)
		r.Get("/feess", controller.GetsFees)
		r.Post("/fees/{id}/update", controller.UpdateFees)

		r.Post("/exam", controller.CreateExam)
		r.Get("/exam/{id}/delete", controller.DeleteExam)
		r.Get("/exam/{id}", controller.GetExam)
		r.Get("/exams", controller.GetsExam)
		r.Post("/exam/{id}/update", controller.UpdateExam)

		r.Post("/studentattendence", controller.CreateAttendence)
		r.Get("/studentattendence/{id}/delete", controller.DeleteTeacherAttendence)
		r.Get("/studentattendence/{id}", controller.GetStudentAttendence)
		r.Get("/studentattendences", controller.GetsStudentAttendence)
		r.Post("/studentattendence/{id}/update", controller.UpdateStudentAttendence)

		r.Post("/sessionyear", controller.CreateSessionYear)
		r.Get("/sessionyear/{id}/delete", controller.DeleteSessionYear)
		r.Get("/sessionyear/{id}", controller.GetSessionYear)
		r.Get("/studentssessions", controller.GetsSessionYear)
		r.Post("/sessionyear/{id}/update", controller.UpdateSessionYear)

		r.Post("/subject", controller.AddSubject)
		r.Get("/subject/{id}/delete", controller.DeleteSubject)
		r.Get("/subject/{id}", controller.GetSubject)
		r.Get("/subjects", controller.GetsSubject)
		r.Post("/subject/{id}/update", controller.UpdateSubject)

		r.Post("/teacherattendence", controller.TeacherAttendence)
		r.Get("/teacherattendence/{id}/delete", controller.DeleteTeacherAttendence)
		r.Get("/teacherattendence/{id}/{date}", controller.GetTeacherAttendence)
		r.Get("/teacherattendences", controller.GetsTeacherAttendence)
		r.Post("/teacherattendence/{id}/update", controller.UpdateTeacherAttendence)

		r.Post("/classteacher", controller.ClassTeacher)
		r.Get("/classteacher/{id}/delete", controller.DeleteClassTeacher)
		r.Get("/classteacher/{id}", controller.GetClassTeacher)
		r.Get("/classteachers", controller.GetsClassTeacher)
		r.Post("/classteacher/{id}/update", controller.UpdateClassTeacher)

		r.Post("/studentexam", controller.StudentExam)
		r.Get("/studentexam/{id}/delete", controller.DeleteStudentExam)
		//r.Get("/studentexam/{id}", controller.StudentExam)
		r.Get("/studentexams", controller.GetsStudentExam)
		r.Post("/studentexam/{id}/update", controller.UpdateStudentExam)
		r.Get("/studentexam/{id}", controller.GetStudentExamByID)

		r.Post("/studentclass", controller.StudentClass)
		r.Get("/studentclass/{id}/delete", controller.DeleteStudentClass)
		r.Get("/studentclass/{id}", controller.GetStudentClass)
		r.Get("/studentclasses", controller.GetsStudentClass)
		r.Post("/studentclass/{id}/update", controller.UpdateStudentClass)

		r.Post("/forgotpassword", controller.ForgotPassword)

		r.Post("/reset_password/{resetKey}", controller.ResetPassword)

		r.Post("/staff", controller.CreateStaff)
		r.Get("/staff/{id}/delete", controller.DeleteStaff)
		r.Get("/staff/{staffType}", controller.GetStaff)
		r.Get("/staff/{staffType}/{id}", controller.GetStaffByID)

		r.Post("/BankDetail", controller.CreateBankDetail)
		r.Get("/BankDetail/{id}/delete", controller.DeleteBankDetail)

		r.Post("/salary", controller.CreateSalary)
		r.Get("/salary/{id}/delete", controller.DeleteSalary)

		r.Post("/homework", controller.CreateHomework)
		r.Get("/homework/{id}/delete", controller.DeleteHomework)

		r.Post("/studenthomework", controller.CreateHomework)
		r.Get("/studenthomework/{id}/delete", controller.DeleteHomework)

		r.Post("/guardian", controller.CreateGuardian)
		r.Get("/guardian/{id}/delete", controller.DeleteGuardian)

		// r.Post("/classsubject", controller.CreateClassSubject)
		// r.Get("/classsubject/{id}/delete", controller.DeleteClassSubject)

	})
}
