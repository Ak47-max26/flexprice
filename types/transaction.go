package types

import (
	ierr "github.com/flexprice/flexprice/internal/errors"
)

// TransactionType represents the type of wallet transaction
type TransactionType string

const (
	TransactionTypeCredit TransactionType = "credit"
	TransactionTypeDebit  TransactionType = "debit"
)

func (t TransactionType) String() string {
	return string(t)
}

func (t TransactionType) Validate() error {
	allowedTypes := []TransactionType{TransactionTypeCredit, TransactionTypeDebit}
	for _, allowedType := range allowedTypes {
		if t == allowedType {
			return nil
		}
	}
	return ierr.NewError("invalid transaction type").
		WithHint("Please provide a valid transaction type").
		WithReportableDetails(map[string]any{
			"allowed": allowedTypes,
			"type":    t,
		}).
		Mark(ierr.ErrValidation)
}

// TransactionStatus represents the status of a wallet transaction
type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
)

func (t TransactionStatus) String() string {
	return string(t)
}

func (t TransactionStatus) Validate() error {
	allowedStatuses := []TransactionStatus{TransactionStatusPending, TransactionStatusCompleted, TransactionStatusFailed}
	for _, allowedStatus := range allowedStatuses {
		if t == allowedStatus {
			return nil
		}
	}
	return ierr.NewError("invalid transaction status").
		WithHint("Please provide a valid transaction status").
		WithReportableDetails(map[string]any{
			"allowed": allowedStatuses,
			"status":  t,
		}).
		Mark(ierr.ErrValidation)
}
