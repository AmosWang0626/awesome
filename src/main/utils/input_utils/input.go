package input_utils

import (
	"regexp"
)

func IsNum(str string) (result bool) {
	pattern := "^(\\d+)$"
	result, _ = regexp.MatchString(pattern, str)

	return
}
