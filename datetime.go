package uadmin

import "time"

func getTZ() *time.Location {
	if UTCTime {
		return time.UTC
	}
	return time.Local
}
