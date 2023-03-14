package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// TxOpt configure transaction
type TxOpt func(*pgx.TxOptions)

// IsolationLevel configures transaction to specific isolation level
func IsolationLevel(isolationLevel pgx.TxIsoLevel) TxOpt {
	return func(opts *pgx.TxOptions) {
		opts.IsoLevel = isolationLevel
	}
}

// ReadWrite configures transaction as both read and write
func ReadWrite() TxOpt {
	return func(opts *pgx.TxOptions) {
		opts.AccessMode = pgx.ReadWrite
	}
}

// TxFunc is function, that is run in transaction context
type TxFunc = func(Ops) error

// WithTx begins and commits transaction and handles errors
func (db *Database) WithTx(ctx context.Context, f TxFunc, opts ...TxOpt) error {
	var txOpt pgx.TxOptions
	for _, opt := range opts {
		opt(&txOpt)
	}

	return db.BeginTxFunc(ctx, txOpt, func(tx pgx.Tx) error {
		return f(tx)
	})
}
