package types

import (
	ierr "github.com/flexprice/flexprice/internal/errors"
	"github.com/samber/lo"
)

// PauseMode represents the mode of pausing a subscription
type PauseMode string

const (
	// PauseModeImmediate pauses the subscription immediately
	PauseModeImmediate PauseMode = "immediate"

	// PauseModeScheduled pauses the subscription at a scheduled time
	PauseModeScheduled PauseMode = "scheduled"

	// PauseModePeriodEnd pauses the subscription at the end of the current billing period
	// Not supported in Phase 0
	PauseModePeriodEnd PauseMode = "period_end"
)

// Validate validates the pause mode
func (m PauseMode) Validate() error {
	allowed := []PauseMode{
		PauseModeImmediate,
		PauseModeScheduled,
		PauseModePeriodEnd,
	}

	if !lo.Contains(allowed, m) {
		return ierr.NewError("invalid pause_mode").
			WithHint("Invalid pause mode").
			WithReportableDetails(map[string]any{
				"type":          m,
				"allowed_types": allowed,
			}).
			Mark(ierr.ErrValidation)
	}

	return nil
}

// String returns the string representation of the pause mode
func (m PauseMode) String() string {
	return string(m)
}

// ResumeMode represents the mode of resuming a subscription
type ResumeMode string

const (
	// ResumeModeImmediate resumes the subscription immediately
	ResumeModeImmediate ResumeMode = "immediate"

	// ResumeModeScheduled resumes the subscription at a scheduled time
	ResumeModeScheduled ResumeMode = "scheduled"

	// ResumeModeAuto resumes the subscription automatically at the end of the pause period
	ResumeModeAuto ResumeMode = "auto"
)

// Validate validates the resume mode
func (m ResumeMode) Validate() error {
	allowed := []ResumeMode{
		ResumeModeImmediate,
		ResumeModeScheduled,
		ResumeModeAuto,
	}

	if !lo.Contains(allowed, m) {
		return ierr.NewError("invalid resume mode").
			WithHint("Invalid resume mode").
			WithReportableDetails(map[string]any{
				"mode":          m,
				"allowed_modes": allowed,
			}).
			Mark(ierr.ErrValidation)
	}

	return nil
}

// String returns the string representation of the resume mode
func (m ResumeMode) String() string {
	return string(m)
}
