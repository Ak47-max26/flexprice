package types

import "time"

func ParseTime(t string) (time.Time, error) {
	return time.Parse(time.RFC3339, t)
}

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseYYYYMMDDToDate converts YYYYMMDD integer to time.Time with beginning of day time
// for ex 20250101 means the credits will expire on 2025-01-01 00:00:00 UTC
// hence they will be available for use until 2024-12-31 23:59:59 UTC
func ParseYYYYMMDDToDate(date *int) *time.Time {
	if date == nil {
		return nil
	}

	parsedTime := time.Date(
		*date/10000,                   // year
		time.Month((*date%10000)/100), // month
		*date%100,                     // day
		0, 0, 0, 0,                    // Set to beginning of day
		time.UTC,
	)
	return &parsedTime
}
