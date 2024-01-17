package adapters

import (
	"github.com/pavelzhirnov/student-service/internal/app/teacherservice/adapters"
	"github.com/pavelzhirnov/student-service/internal/model"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
)

func StudentToPb(student *model.Student) *api.Student {
	return &api.Student{
		Id:       student.ID,
		FullName: student.FullName,
		Age:      student.Age,
		Salary:   student.Salary,
		Teachers: adapters.TeachersToPb(student.Teachers),
	}
}

func StudentsToPb(students []*model.Student) []*api.Student {
	apiStudents := make([]*api.Student, 0, len(students))

	for _, item := range students {
		apiStudents = append(apiStudents, StudentToPb(item))
	}
	return apiStudents
}
