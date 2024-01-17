package teacherservice

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pavelzhirnov/student-service/internal/app/teacherservice/adapters"
	"github.com/pavelzhirnov/student-service/pkg/logging"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *Service) PatchTeacher(ctx context.Context, req *api.UpdateTeacherRequest) (*api.Teacher, error) {
	if err := ValidateUpdateTeacherRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logging.GetLogger(ctx).Info("update teacher")
	teacher, err := s.teacherRepository.Update(ctx, adapters.UpdateTeacherFromPb(req))
	if err != nil {
		return nil, status.Error(codes.Internal, "error update teacher")
	}

	return adapters.TeacherToPb(teacher), nil
}

func ValidateUpdateTeacherRequest(req *api.UpdateTeacherRequest) error {
	err := validation.Errors{
		"id":            validation.Validate(req.GetId(), validation.Required),
		"position_type": validation.Validate(req.GetPositionType(), positionTypeValidationRule),
		"full_name":     validation.Validate(strings.TrimSpace(req.GetFullName()), validation.Required, validation.Length(1, 0)),
	}.Filter()

	if err != nil {
		return err
	}
	return nil
}
