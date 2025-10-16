package types

import (
	"database/sql/driver"
	"encoding/json"

	ierr "github.com/flexprice/flexprice/internal/errors"
)

// Metadata represents a JSONB field for storing key-value pairs
type Metadata map[string]string

// Scan implements the sql.Scanner interface for Metadata
func (m *Metadata) Scan(value interface{}) error {
	if value == nil {
		*m = make(Metadata)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return ierr.NewError("failed to unmarshal JSONB value").
			WithHint("Please provide a valid JSON value").
			Mark(ierr.ErrValidation)
	}

	result := make(Metadata)
	err := json.Unmarshal(bytes, &result)
	*m = result
	return err
}

// Value implements the driver.Valuer interface for Metadata
func (m Metadata) Value() (driver.Value, error) {
	if m == nil {
		return json.Marshal(make(Metadata))
	}
	return json.Marshal(m)
}
