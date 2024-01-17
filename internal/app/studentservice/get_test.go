package studentservice

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pavelzhirnov/student-service/internal/app/studentservice/adapters"
	"github.com/pavelzhirnov/student-service/internal/model"
	"github.com/pavelzhirnov/student-service/internal/repository"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestGetStudent(t *testing.T) {
	t.Run("validation Error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetStudentRequest{
			Id: "",
		}

		resource, err := te.studentService.GetStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		var expectedResponse *api.Student
		assert.Equal(t, expectedResponse, resource)
	})

	t.Run("repository Error", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.GetStudentRequest{
			Id: studentID,
		}

		te.studentRepository.EXPECT().Get(te.ctx, req.GetId()).Return(nil, errors.New("any catalog error"))

		resource, err := te.studentService.GetStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedResponse *api.Student
		assert.Equal(t, expectedResponse, resource)
	})

	t.Run("student not found", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.GetStudentRequest{
			Id: studentID,
		}

		te.studentRepository.EXPECT().Get(te.ctx, req.GetId()).Return(nil, repository.ErrEntityNotFound)

		student, err := te.studentService.GetStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
		var expectedResponse *api.Student
		assert.Equal(t, expectedResponse, student)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.GetStudentRequest{
			Id: studentID,
		}

		modelStudent := &model.Student{
			ID:       studentID,
			FullName: "Павел Жирнов",
			Age:      18,
			Salary:   123425,
		}

		te.studentRepository.EXPECT().Get(te.ctx, req.GetId()).Return(modelStudent, nil)

		student, err := te.studentService.GetStudent(te.ctx, req)
		assert.NoError(t, err)
		expectedStudent := adapters.StudentToPb(modelStudent)
		assert.Equal(t, expectedStudent, student)
	})
}
