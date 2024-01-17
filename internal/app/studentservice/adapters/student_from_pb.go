package adapters

import (
	"github.com/pavelzhirnov/student-service/internal/app/teacherservice/adapters"
	"github.com/pavelzhirnov/student-service/internal/model"
	"github.com/pavelzhirnov/student-service/internal/repository"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"strings"
)

func CreateStudentFromPb(student *api.CreateStudentRequest) *model.Student {
	return &model.Student{
		FullName: strings.TrimSpace(student.GetFullName()),
		Age:      student.GetAge(),
		Salary:   student.GetSalary(),
		Teachers: CreateTeachersFromPb(student.GetTeachers()),
	}
}

func CreateTeachersFromPb(teachers []*api.CreateTeacherRequest) []*model.Teacher {
	modelTeachers := make([]*model.Teacher, 0, len(teachers))
	for _, teacher := range teachers {
		modelTeachers = append(modelTeachers, adapters.CreateTeacherFromPb(teacher))
	}

	return modelTeachers
}

func UpdateStudentFromPb(student *api.UpdateStudentRequest) *model.Student {
	return &model.Student{
		ID:       student.GetId(),
		FullName: strings.TrimSpace(student.GetFullName()),
		Age:      student.GetAge(),
		Salary:   student.GetSalary(),
		Teachers: UpdateTeachersFromPb(student.GetTeachers()),
	}
}

func UpdateTeachersFromPb(teachers []*api.UpdateTeacherRequest) []*model.Teacher {
	modelTeachers := make([]*model.Teacher, 0, len(teachers))
	for _, teacher := range teachers {
		modelTeachers = append(modelTeachers, adapters.UpdateTeacherFromPb(teacher))
	}

	return modelTeachers
}

func ListFilterStudentFromPb(filter *api.ListStudentRequest) *repository.StudentListFilter {
	return &repository.StudentListFilter{
		IDList: filter.GetStudentIds(),
	}
}
