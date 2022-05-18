package utils

import (
	"strconv"
)

func MakeHour(digits []int) []string {
	var returnHours = make([]string, 0)
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			if i != j {
				value, _ := strconv.Atoi(strconv.Itoa(digits[i]) + strconv.Itoa(digits[j]))
				reverseValue, _ := strconv.Atoi(strconv.Itoa(digits[j]) + strconv.Itoa(digits[i]))
				if value < 24 && !IsAppended(value, returnHours) {
					if len(strconv.Itoa(value)) != 1 {
						returnHours = append(returnHours, strconv.Itoa(value))
					}
				}
				if reverseValue < 24 && !IsAppended(reverseValue, returnHours) {
					if len(strconv.Itoa(reverseValue)) != 1 {
						returnHours = append(returnHours, strconv.Itoa(reverseValue))
					}

				}
			}
		}
	}
	return returnHours
}

func AddMinutes(hours []string, digits []int) []string {
	var hourAndMinutes = make([]string, 0)
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			if i != j {
				value, _ := strconv.Atoi(strconv.Itoa(digits[i]) + strconv.Itoa(digits[j]))
				reverseValue, _ := strconv.Atoi(strconv.Itoa(digits[j]) + strconv.Itoa(digits[i]))
				if value < 60 {
					if len(strconv.Itoa(value)) != 1 {
						hourAndMinutes = appendMinutesToHours(hours, value, hourAndMinutes)
					}
				}
				if reverseValue < 60 {
					if len(strconv.Itoa(reverseValue)) != 1 {
						hourAndMinutes = appendMinutesToHours(hours, reverseValue, hourAndMinutes)
					}
				}
			}
		}
	}

	return hourAndMinutes
}

func appendMinutesToHours(hours []string, value int, hourAndMinutes []string) []string {
	for j := 0; j < len(hours); j++ {
		if len(hourAndMinutes) > 0 {
			if !checkForDuplicateTime(hourAndMinutes, hours[j]+":"+strconv.Itoa(value)) {
				hourAndMinutes = append(hourAndMinutes, (hours[j])+":"+strconv.Itoa(value))
			}
		} else {
			hourAndMinutes = append(hourAndMinutes, (hours[j])+":"+strconv.Itoa(value))
		}
	}
	return hourAndMinutes
}

func checkForDuplicateTime(hourAndMinutes []string, time string) bool {
	for i := 0; i < len(hourAndMinutes); i++ {
		if hourAndMinutes[i] == time {
			return true //found
		}
	}
	return false // not found, safe to append
}
