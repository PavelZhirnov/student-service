package studentservice

import (
	"errors"
	"github.com/google/uuid"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestDeleteStudent(t *testing.T) {
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

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.GetStudentRequest{
			Id: studentID,
		}

		te.studentRepository.EXPECT().Delete(te.ctx, req.GetId()).Return(errors.New("any catalog error"))

		resource, err := te.studentService.DeleteStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedResponse *api.SimpleResponse
		assert.Equal(t, expectedResponse, resource)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.GetStudentRequest{
			Id: studentID,
		}

		te.studentRepository.EXPECT().Delete(te.ctx, req.GetId()).Return(nil)

		resource, err := te.studentService.DeleteStudent(te.ctx, req)
		assert.NoError(t, err)
		expectedResponse := &api.SimpleResponse{}
		assert.Equal(t, expectedResponse, resource)
	})
}
