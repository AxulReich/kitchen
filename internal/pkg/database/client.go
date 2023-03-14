package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	pingTimeout = 10 * time.Second
)

type Tx interface {
	WithTx(ctx context.Context, f TxFunc, opts ...TxOpt) error
}

type Ops interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type DB interface {
	Ops
	Tx
}

type Options struct {
	DSN string
}

type beginTx interface {
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type Database struct {
	Ops
	beginTx
}

func NewDB(ctx context.Context, opts Options) (*Database, error) {
	db, err := pgxpool.Connect(context.Background(), opts.DSN)
	if err != nil {
		return nil, fmt.Errorf("can't init pool: %w", err)
	}
	defer db.Close()

	pingCtx, cancel := context.WithTimeout(ctx, pingTimeout)
	defer cancel()

	err = db.Ping(pingCtx)
	if err != nil {
		return nil, fmt.Errorf("database is unreachable: %w", err)
	}
	return &Database{
		Ops:     db,
		beginTx: db,
	}, nil
}