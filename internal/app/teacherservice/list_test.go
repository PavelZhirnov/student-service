package teacherservice

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pavelzhirnov/student-service/internal/app/teacherservice/adapters"
	"github.com/pavelzhirnov/student-service/internal/model"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestListTeacher(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListTeacherRequest{
			TeacherIds: []string{},
		}

		teachers, err := te.teacherService.ListTeachers(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		var expectedTeachers *api.ListTeacherResponse
		assert.Equal(t, expectedTeachers, teachers)
	})

	t.Run("repository Error", func(t *testing.T) {
		te := newTestEnv(t)

		teacherID := uuid.New().String()
		req := &api.ListTeacherRequest{
			TeacherIds: []string{teacherID},
		}

		expectedMockStudentIds := adapters.ListFilterTeacherFromPb(req)
		te.teacherRepository.EXPECT().List(te.ctx, expectedMockStudentIds).Return(nil, errors.New("any catalog error"))

		teachers, err := te.teacherService.ListTeachers(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedTeachers *api.ListTeacherResponse
		assert.Equal(t, expectedTeachers, teachers)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		teacherID := uuid.New().String()
		req := &api.ListTeacherRequest{
			TeacherIds: []string{teacherID},
		}
		expectedMockTeacherIds := adapters.ListFilterTeacherFromPb(req)
		modelTeachers := []*model.Teacher{
			{
				ID:           teacherID,
				FullName:     "name",
				PositionType: 1,
			},
		}

		te.teacherRepository.EXPECT().List(te.ctx, expectedMockTeacherIds).Return(modelTeachers, nil)

		resource, err := te.teacherService.ListTeachers(te.ctx, req)
		assert.NoError(t, err)
		expectedResponse := &api.ListTeacherResponse{
			Teachers: adapters.TeachersToPb(modelTeachers),
		}
		assert.Equal(t, expectedResponse, resource)
	})
}
