package studentservice

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pavelzhirnov/student-service/internal/app/studentservice/adapters"
	"github.com/pavelzhirnov/student-service/internal/repository"
	"github.com/pavelzhirnov/student-service/pkg/logging"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetStudent(ctx context.Context, req *api.GetStudentRequest) (*api.Student, error) {
	if err := validateStudentIDRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logging.GetLogger(ctx).Info("get student")
	student, err := s.studentRepository.Get(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, repository.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, "student not found")
		}
		return nil, status.Error(codes.Internal, "error get student")
	}

	return adapters.StudentToPb(student), nil
}

func validateStudentIDRequest(req *api.GetStudentRequest) error {
	err := validation.Errors{
		"student_id": validation.Validate(req.GetId(), validation.Required),
	}.Filter()
	if err != nil {
		return err
	}
	return nil
}
