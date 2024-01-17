package studentservice

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pavelzhirnov/student-service/internal/app/studentservice/adapters"
	"github.com/pavelzhirnov/student-service/internal/model"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestListStudent(t *testing.T) {
	t.Run("validate error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListStudentRequest{
			StudentIds: []string{},
		}

		students, err := te.studentService.ListStudents(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		var expectedStudents *api.ListStudentResponse
		assert.Equal(t, expectedStudents, students)
	})

	t.Run("repository Error", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.ListStudentRequest{
			StudentIds: []string{studentID},
		}

		expectedMockStudentIds := adapters.ListFilterStudentFromPb(req)
		te.studentRepository.EXPECT().List(te.ctx, expectedMockStudentIds).Return(nil, errors.New("any catalog error"))

		students, err := te.studentService.ListStudents(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedStudents *api.ListStudentResponse
		assert.Equal(t, expectedStudents, students)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.ListStudentRequest{
			StudentIds: []string{studentID},
		}

		expectedMockStudentIds := adapters.ListFilterStudentFromPb(req)
		modelStudents := []*model.Student{
			{
				ID:       studentID,
				FullName: "Павел Жирнов",
				Age:      19,
				Salary:   12345,
			},
		}

		te.studentRepository.EXPECT().List(te.ctx, expectedMockStudentIds).Return(modelStudents, nil)

		resource, err := te.studentService.ListStudents(te.ctx, req)
		assert.NoError(t, err)
		expectedResponse := &api.ListStudentResponse{
			Students: adapters.StudentsToPb(modelStudents),
		}
		assert.Equal(t, expectedResponse, resource)
	})
}
