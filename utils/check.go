package utils

import "strconv"

func IsAppended(value int, returnHours []string) bool {
	if len(returnHours) != 0 {
		for i := 0; i < len(returnHours); i++ {
			val, _ := strconv.Atoi(returnHours[i])
			if val == value {
				return true
			}
		}
	}
	return false
}
