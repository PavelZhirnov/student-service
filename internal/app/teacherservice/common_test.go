package teacherservice

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pavelzhirnov/student-service/internal/repository"
	"testing"
)

type testEnv struct {
	ctx  context.Context
	ctrl *gomock.Controller

	teacherRepository *repository.MockTeacherRepository

	teacherService *Service
}

func newTestEnv(t *testing.T) *testEnv {
	tEnv := &testEnv{}
	tEnv.ctx = context.Background()
	tEnv.ctrl = gomock.NewController(t)

	tEnv.teacherRepository = repository.NewMockTeacherRepository(tEnv.ctrl)

	tEnv.teacherService = NewService(tEnv.teacherRepository)
	return tEnv
}
