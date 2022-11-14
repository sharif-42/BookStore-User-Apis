package time_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func getNowStandard() time.Time {
	return time.Now().UTC()
}

func getNowLocal() time.Time {
	return time.Now()
}

func GetLocalNowTimeString() string {
	// for server local time
	return getNowLocal().Format(apiDateLayout)
}

func getNowStandardTimeString() string {
	// for standard time
	return getNowStandard().Format(apiDateLayout)
}
