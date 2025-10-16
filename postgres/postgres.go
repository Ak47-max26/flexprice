package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/flexprice/flexprice/internal/config"
	"github.com/flexprice/flexprice/internal/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB wraps sqlx.DB to provide transaction management
type DB struct {
	*sqlx.DB
	logger *logger.Logger
}

// Querier interface defines all database operations
// Both *sqlx.DB and *sqlx.Tx implement these methods
type Querier interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
}

// NewDB creates a new DB instance
func NewDB(config *config.Configuration, logger *logger.Logger) (*DB, error) {
	dsn := config.Postgres.GetDSN()
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &DB{DB: db, logger: logger}, nil
}

// Close closes the database connection
func (db *DB) Close() {
	if err := db.DB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}

// GetQuerier returns either the transaction from context or the base DB
func (db *DB) GetQuerier(ctx context.Context) Querier {
	if tx, ok := GetTx(ctx); ok {
		return NewTracedQuerier(tx.Tx, db.logger, tx.ID)
	}
	return NewTracedQuerier(db.DB, db.logger, "")
}

// NamedExecContext is a helper method that wraps NamedExec with context
func (db *DB) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	q := db.GetQuerier(ctx)
	return q.NamedExec(query, arg)
}

// NamedQueryContext is a helper method that wraps NamedQuery with context
func (db *DB) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	q := db.GetQuerier(ctx)
	return q.NamedQuery(query, arg)
}
