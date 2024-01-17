package studentservice

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pavelzhirnov/student-service/internal/repository"
	"testing"
)

type testEnv struct {
	ctx  context.Context
	ctrl *gomock.Controller

	studentRepository *repository.MockStudentRepository

	studentService *Service
}

func newTestEnv(t *testing.T) *testEnv {
	tEnv := &testEnv{}
	tEnv.ctx = context.Background()
	tEnv.ctrl = gomock.NewController(t)

	tEnv.studentRepository = repository.NewMockStudentRepository(tEnv.ctrl)

	tEnv.studentService = NewService(tEnv.studentRepository)
	return tEnv
}
