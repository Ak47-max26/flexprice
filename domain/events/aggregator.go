package events

import (
	"context"

	"github.com/flexprice/flexprice/internal/types"
)

type Aggregator interface {
	// GetQuery returns the query for this aggregation
	GetQuery(ctx context.Context, params *UsageParams) string

	// GetType returns the aggregation type
	GetType() types.AggregationType
}
