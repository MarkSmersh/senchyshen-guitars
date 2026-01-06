package api

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/MarkSmersh/senchyshen-guitars/tables"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init() {
	dbUrl := os.Getenv("DB_URL")

	if len(dbUrl) <= 0 {
		slog.Error("DB_URL enviroment variable is missing")
		os.Exit(1)
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	s := NewServer(conn, "", 1488)

	tables.Init(conn)
	s.SetRouters()
	s.Start()

	slog.Info(
		fmt.Sprintf("Server is started on port :%d", s.Port),
	)
}
