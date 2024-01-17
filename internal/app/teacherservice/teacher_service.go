package teacherservice

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pavelzhirnov/student-service/internal/repository"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
)

var positionTypeValidationRule = validation.In(
	api.PositionType_POSTGRADUATE,
	api.PositionType_ASSISTANT,
	api.PositionType_DEAN,
)

type Service struct {
	api.UnimplementedTeacherServiceServer
	teacherRepository repository.TeacherRepository
}

func NewService(teacherRepository repository.TeacherRepository) *Service {
	return &Service{
		teacherRepository: teacherRepository,
	}
}
