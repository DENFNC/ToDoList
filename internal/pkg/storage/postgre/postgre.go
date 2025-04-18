package psql

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	DB   *pgxpool.Pool
	Dial *goqu.DialectWrapper
}

func New(conn string) (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, conn)
	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(ctx); err != nil {
		return nil, err
	}

	dialect := goqu.Dialect("postgres")

	return &Storage{
		DB:   dbpool,
		Dial: &dialect,
	}, nil
}

func (s *Storage) Stop() {
	s.DB.Close()
}
