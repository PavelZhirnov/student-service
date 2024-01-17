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

func TestUpdateStudent(t *testing.T) {
	t.Run("validate error", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.UpdateStudentRequest{
			Id:       studentID,
			FullName: "",
			Age:      0,
			Salary:   0,
		}

		student, err := te.studentService.PatchStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		var expectedStudent *api.Student
		assert.Equal(t, expectedStudent, student)
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.UpdateStudentRequest{
			Id:       studentID,
			FullName: "Паша Жирнов",
			Age:      13,
			Salary:   123456,
		}

		expectedMockStudent := adapters.UpdateStudentFromPb(req)
		te.studentRepository.EXPECT().Update(te.ctx, expectedMockStudent).Return(nil, errors.New("any catalog error"))

		student, err := te.studentService.PatchStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedStudent *api.Student
		assert.Equal(t, expectedStudent, student)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		studentID := uuid.New().String()
		req := &api.UpdateStudentRequest{
			Id:       studentID,
			FullName: "Паша Жирнов",
			Age:      13,
			Salary:   123456,
		}

		expectedMockStudent := adapters.UpdateStudentFromPb(req)
		modelStudent := &model.Student{
			ID:       studentID,
			FullName: "name",
			Age:      12,
			Salary:   123123,
		}

		te.studentRepository.EXPECT().Update(te.ctx, expectedMockStudent).Return(modelStudent, nil)

		student, err := te.studentService.PatchStudent(te.ctx, req)
		assert.NoError(t, err)
		expectedStudent := adapters.StudentToPb(modelStudent)
		assert.Equal(t, expectedStudent, student)
	})
}
