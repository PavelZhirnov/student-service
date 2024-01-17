package app

import (
	"context"
	"fmt"
	"github.com/pavelzhirnov/student-service/internal/app/studentservice"
	"github.com/pavelzhirnov/student-service/internal/app/teacherservice"
	"github.com/pavelzhirnov/student-service/internal/bootstrap"
	"github.com/pavelzhirnov/student-service/internal/closer"
	"github.com/pavelzhirnov/student-service/internal/config"
	"github.com/pavelzhirnov/student-service/internal/repository"
	"github.com/pavelzhirnov/student-service/pkg/logging"
	api "github.com/pavelzhirnov/student-service/pkg/studentServiceApi"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, cfg *config.Config, logger logging.Logger) error {
	s := grpc.NewServer()

	ctx, cancel := context.WithCancel(ctx)
	ctx = logging.ContextWithLogger(ctx, logger)

	logging.GetLogger(ctx).Infof("listen GRPC server to %s", cfg.Ports.GRPC)
	l, err := net.Listen("tcp", cfg.Ports.GRPC)
	if err != nil {
		cancel()
		logging.GetLogger(ctx).Fatalf("failed to listen tcp %s, %v\n", cfg.Ports.GRPC, err)
	}

	initServices(ctx, s, cfg)

	go func() {
		if err = s.Serve(l); err != nil {
			logging.GetLogger(ctx).Fatal("ERROR: ", err.Error())
		}
	}()

	gracefulShutdown(s, cancel)
	return nil
}

func initServices(ctx context.Context, s *grpc.Server, cfg *config.Config) {
	logging.GetLogger(ctx).Info("start connect to DB")
	conn, err := bootstrap.InitDB(cfg)
	if err != nil {
		logging.GetLogger(ctx).Fatalf("not connect to db :%v", err)
	}

	logging.GetLogger(ctx).Info("register services")
	api.RegisterStudentServiceServer(s, studentservice.NewService(
		repository.NewStudentRepository(
			conn,
			repository.NewTeacherRepository(conn),
		),
	),
	)
	api.RegisterTeacherServiceServer(s, teacherservice.NewService(
		repository.NewTeacherRepository(conn),
	),
	)
}

func gracefulShutdown(s *grpc.Server, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done")
	fmt.Println(errorMessage)
	s.GracefulStop()
	cancel()
	closer.CloseAll()
}
