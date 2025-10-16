package types

import (
	ierr "github.com/flexprice/flexprice/internal/errors"
	"github.com/samber/lo"
)

type ResetUsage string

const (
	ResetUsageBillingPeriod ResetUsage = "BILLING_PERIOD"
	ResetUsageNever         ResetUsage = "NEVER"
	// TODO: support more reset values like MONTHLY, WEEKLY, DAILY in future
)

// Validate ensures the ResetUsage value is valid
func (r ResetUsage) Validate() error {
	if r == "" {
		return nil
	}

	allowedValues := []ResetUsage{
		ResetUsageBillingPeriod,
		ResetUsageNever,
	}

	if !lo.Contains(allowedValues, r) {
		return ierr.NewError("invalid reset usage").
			WithHint("Invalid reset usage").
			WithReportableDetails(map[string]any{
				"allowed_values": allowedValues,
				"provided_value": r,
			}).
			Mark(ierr.ErrValidation)
	}

	return nil
}
