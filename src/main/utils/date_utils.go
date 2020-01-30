package utils

import (
	"strings"
	"time"
)

var Year2Day = "yyyy-MM-dd"
var Year2Second = "yyyy-MM-dd hh:mm:ss"

/**
 * base date 2006/01/02 15:04:05
 */
func Format(time time.Time, format string) string {
	format = strings.ReplaceAll(format, "yyyy", "2006")
	format = strings.ReplaceAll(format, "MM", "01")
	format = strings.ReplaceAll(format, "dd", "02")
	format = strings.ReplaceAll(format, "hh", "15")
	format = strings.ReplaceAll(format, "mm", "04")
	format = strings.ReplaceAll(format, "ss", "05")

	return time.Format(format)
}
