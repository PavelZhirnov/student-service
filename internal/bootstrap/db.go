package bootstrap

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib" // необходим из-за драйвера pgx
	"github.com/jmoiron/sqlx"
	"github.com/pavelzhirnov/student-service/internal/closer"
	"github.com/pavelzhirnov/student-service/internal/config"
)

func InitDB(cfg *config.Config) (*sqlx.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DBName)
	conn, err := sqlx.Connect("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("can't connect to pg instance, %v", err)
	}

	closer.Add(conn.Close)

	return conn.Unsafe(), nil
}
