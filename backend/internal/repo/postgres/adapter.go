package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RetryAdapterIntf interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Rows
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

type RetryAdapter struct {
	db              *pgxpool.Pool
	numberOfRetries int
	sleepTime       time.Duration
}

func NewRetryAdapter(db *pgxpool.Pool, retryNumber int, sleepTime time.Duration) *RetryAdapter {
	return &RetryAdapter{
		db:              db,
		numberOfRetries: retryNumber,
		sleepTime:       sleepTime,
	}
}

func (p *RetryAdapter) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	for i := 0; i < p.numberOfRetries; i++ {
		commTag, err := p.db.Exec(ctx, sql, arguments...)
		if err == nil {
			return commTag, nil
		}
		time.Sleep(p.sleepTime)
	}
	return pgconn.CommandTag{}, err
}

func (p *RetryAdapter) QueryRow(ctx context.Context, sql string, args ...any) pgx.Rows {
	var rows pgx.Rows
	for i := 0; i < p.numberOfRetries; i++ {
		rows, err := p.db.Query(ctx, sql, args...)
		if err == nil {
			rows.Next()
			return rows
		}
		time.Sleep(p.sleepTime)
	}
	return rows
}

func (p *RetryAdapter) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	var (
		rows pgx.Rows
		err  error
	)
	for i := 0; i < p.numberOfRetries; i++ {
		rows, err = p.db.Query(ctx, sql, args...)
		if err == nil {
			return rows, nil
		}
		time.Sleep(p.sleepTime)
	}
	return rows, err
}
