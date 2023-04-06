package datetime

import "time"

// GetDateString Returns date in YYYYMMDDHHmmss format
func GetDateString() string {
	return time.Now().Format("20060102150405")
}
