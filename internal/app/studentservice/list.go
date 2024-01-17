package studentservice

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pavelzhirnov/student-service/internal/app/studentservice/adapters"
	"github.com/pavelzhirnov/student-service/pkg/logging"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListStudents(ctx context.Context, req *api.ListStudentRequest) (*api.ListStudentResponse, error) {
	if err := validateListStudentRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logging.GetLogger(ctx).Info("list student")
	studentList, err := s.studentRepository.List(ctx, adapters.ListFilterStudentFromPb(req))
	if err != nil {
		return nil, status.Error(codes.Internal, "error list student")
	}

	return &api.ListStudentResponse{
		Students: adapters.StudentsToPb(studentList),
	}, nil
}

func validateListStudentRequest(req *api.ListStudentRequest) error {
	err := validation.Errors{
		"student_ids": validation.Validate(req.GetStudentIds(), validation.Required),
	}.Filter()

	if err != nil {
		return err
	}
	return nil
}
