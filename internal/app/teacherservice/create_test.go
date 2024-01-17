package teacherservice

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/pavelzhirnov/student-service/internal/app/teacherservice/adapters"
	"github.com/pavelzhirnov/student-service/internal/model"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestCreateTeacher(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.CreateTeacherRequest{
			PositionType: 0,
			FullName:     "",
			StudentId:    studentID,
		}

		teacher, err := te.teacherService.CreateTeacher(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		var expectedTeacher *api.Teacher
		assert.Equal(t, expectedTeacher, teacher)
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.CreateTeacherRequest{
			FullName:     "name",
			PositionType: 1,
			StudentId:    studentID,
		}

		expectedMockTeacher := adapters.CreateTeacherFromPb(req)
		te.teacherRepository.EXPECT().Create(te.ctx, []*model.Teacher{expectedMockTeacher}).Return(nil, errors.New("any catalog error"))

		teacher, err := te.teacherService.CreateTeacher(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedTeacher *api.Teacher
		assert.Equal(t, expectedTeacher, teacher)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		teacherID := uuid.New().String()
		req := &api.CreateTeacherRequest{
			FullName:     "name",
			PositionType: 1,
			StudentId:    studentID,
		}
		expectedMockTeacher := adapters.CreateTeacherFromPb(req)
		modelTeachers := []*model.Teacher{
			{
				ID:           teacherID,
				FullName:     "name",
				PositionType: 1,
				StudentID:    studentID,
			},
		}

		te.teacherRepository.EXPECT().Create(te.ctx, []*model.Teacher{expectedMockTeacher}).Return(modelTeachers, nil)

		teacher, err := te.teacherService.CreateTeacher(te.ctx, req)
		assert.NoError(t, err)
		expectedTeacher := adapters.TeacherToPb(modelTeachers[len(modelTeachers)-1:][0])
		assert.Equal(t, expectedTeacher, teacher)
	})
}

func TestValidateCreateTeacherRequest(t *testing.T) {
	type args struct {
		req *api.CreateTeacherRequest
	}

	studentID := uuid.New().String()

	tests := []struct {
		name       string
		args       args
		errorField string
		errorCode  string
	}{
		{
			name: "full_name no filled",
			args: args{
				&api.CreateTeacherRequest{
					FullName:     "",
					PositionType: 1,
					StudentId:    studentID,
				},
			},
			errorField: "full_name",
			errorCode:  validation.ErrRequired.Code(),
		},
		{
			name: "position_type in invalid",
			args: args{
				&api.CreateTeacherRequest{
					FullName:     "name",
					PositionType: 134,
					StudentId:    studentID,
				},
			},
			errorField: "position_type",
			errorCode:  validation.ErrInInvalid.Code(),
		},
		{
			name: "student_id no filled",
			args: args{
				&api.CreateTeacherRequest{
					FullName:     "name",
					PositionType: 1,
					StudentId:    "",
				},
			},
			errorField: "student_id",
			errorCode:  validation.ErrRequired.Code(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateCreateTeacherRequest(test.args.req, true)
			if !assert.Error(t, err, "error expected") {
				assert.FailNow(t, "error expected")
			}
			validationErr, ok := err.(validation.Errors)
			if !ok {
				assert.FailNow(t, fmt.Sprintf("expected validation errors, but not received: [%+v]", err))
			}
			fieldErr, ok := validationErr[test.errorField]
			if !ok {
				assert.FailNow(t, fmt.Sprintf("expected error in field [%s] errors, but received: [%+v]", test.errorField, err))
			}

			errObject, ok := fieldErr.(validation.ErrorObject)
			if !ok {
				assert.FailNow(t, fmt.Sprintf("expected error object, but not received: [%+v]", fieldErr))
			}
			assert.Equal(t, test.errorCode, errObject.Code(), "error code must be equal to expected")
		})
	}
}
