package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/flexprice/flexprice/internal/logger"
	"github.com/jmoiron/sqlx"
)

// QueryTracer wraps database operations with tracing and logging
type QueryTracer struct {
	logger *logger.Logger
	query  string
	params interface{}
	start  time.Time
	txID   string
}

// NewQueryTracer creates a new query tracer
func NewQueryTracer(logger *logger.Logger, query string, params interface{}, txID string) *QueryTracer {
	return &QueryTracer{
		logger: logger,
		query:  query,
		params: params,
		start:  time.Now(),
		txID:   txID,
	}
}

// Done logs the query completion
func (qt *QueryTracer) Done(err error) {
	duration := time.Since(qt.start)
	fields := []interface{}{
		"duration_ms", duration.Milliseconds(),
		"query", qt.query,
		"params", fmt.Sprintf("%+v", qt.params),
	}
	if qt.txID != "" {
		fields = append(fields, "tx_id", qt.txID)
	}
	if err != nil {
		fields = append(fields, "error", err.Error())
		qt.logger.Errorw("database query failed", fields...)
		return
	}
	qt.logger.Debugw("database query completed", fields...)
}

// TracedQuerier wraps a Querier with tracing
type TracedQuerier struct {
	Querier
	logger *logger.Logger
	txID   string
}

// NewTracedQuerier creates a new traced querier
func NewTracedQuerier(q Querier, logger *logger.Logger, txID string) *TracedQuerier {
	return &TracedQuerier{
		Querier: q,
		logger:  logger,
		txID:    txID,
	}
}

// ExecContext traces ExecContext calls
func (tq *TracedQuerier) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tracer := NewQueryTracer(tq.logger, query, args, tq.txID)
	result, err := tq.Querier.ExecContext(ctx, query, args...)
	tracer.Done(err)
	return result, err
}

// NamedExec traces NamedExec calls
func (tq *TracedQuerier) NamedExec(query string, arg interface{}) (sql.Result, error) {
	tracer := NewQueryTracer(tq.logger, query, arg, tq.txID)
	result, err := tq.Querier.NamedExec(query, arg)
	tracer.Done(err)
	return result, err
}

// QueryContext traces QueryContext calls
func (tq *TracedQuerier) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	tracer := NewQueryTracer(tq.logger, query, args, tq.txID)
	rows, err := tq.Querier.QueryContext(ctx, query, args...)
	tracer.Done(err)
	return rows, err
}

// NamedQuery traces NamedQuery calls
func (tq *TracedQuerier) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	tracer := NewQueryTracer(tq.logger, query, arg, tq.txID)
	rows, err := tq.Querier.NamedQuery(query, arg)
	tracer.Done(err)
	return rows, err
}
