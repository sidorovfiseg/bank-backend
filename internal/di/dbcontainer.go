package di

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseContainer struct {
	dbpool *pgxpool.Pool
}

// New creates a new database container connected to postgres db 
func New(ctx context.Context) *DatabaseContainer {

	databaseURL := os.Getenv("DATABASE_URL") 
	if databaseURL == "" {
		slog.Error("DATABASE_URL is not set")
	}

	dbpool, err := pgxpool.New(ctx, databaseURL) 
	
	if (err != nil) {
		slog.Error("create connection pool", err)
	}
	
	
	if err:= dbpool.Ping(ctx); err != nil {
		slog.Error("ping", err, dbpool)
	}
	
	return &DatabaseContainer{
			dbpool: dbpool,
	}
}

func (dbcnt *DatabaseContainer) CloseConnection() {
	dbcnt.dbpool.Close()
	slog.Info("Database connection closed")
}

