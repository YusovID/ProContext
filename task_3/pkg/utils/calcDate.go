package utils

import (
	"time"
)

const dateFormatLayout = "02/01/2006"

func CalcDate(days int) string {
	now := time.Now()
	resultTime := now.AddDate(0, 0, days)

	resultDateString := resultTime.Format(dateFormatLayout)

	return resultDateString
}
