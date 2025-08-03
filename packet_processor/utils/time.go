package utils

import (
	"time"
)

func GetNowAsFormattedString() string {
	return time.Now().Format("2006-01-02_15-04-05")
}
