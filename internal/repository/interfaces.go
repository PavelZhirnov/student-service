package repository

import (
	"context"
	"github.com/pavelzhirnov/student-service/internal/model"
)

type StudentRepository interface {
	Create(context.Context, *model.Student) (*model.Student, error)
	List(context.Context, *StudentListFilter) ([]*model.Student, error)
	Get(context.Context, string) (*model.Student, error)
	Update(context.Context, *model.Student) (*model.Student, error)
	Delete(context.Context, string) error
}

type TeacherRepository interface {
	Get(context.Context, string) (*model.Teacher, error)
	Create(context.Context, []*model.Teacher) ([]*model.Teacher, error)
	Update(context.Context, *model.Teacher) (*model.Teacher, error)
	UpdateTeachers(context.Context, string, []*model.Teacher) ([]*model.Teacher, error)
	DeleteByStudentID(context.Context, string) error
	List(context.Context, *TeacherListFilter) ([]*model.Teacher, error)
}
