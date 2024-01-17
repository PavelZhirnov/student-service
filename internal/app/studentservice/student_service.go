package studentservice

import (
	"github.com/pavelzhirnov/student-service/internal/repository"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
)

type Service struct {
	api.UnimplementedStudentServiceServer
	studentRepository repository.StudentRepository
}

func NewService(studentRepository repository.StudentRepository) *Service {
	return &Service{
		studentRepository: studentRepository,
	}
}
