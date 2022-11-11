package uadmin

import "time"

func getTZ() *time.Location {
	if TimeZone == "" || TimeZone == "local" {
		return time.Local
	}
	loc, err := time.LoadLocation(TimeZone)
	if err != nil {
		Trail(ERROR, "Unknown time zone %s. System will use local time zone instead", TimeZone)
		return time.Local
	}
	return loc
}
