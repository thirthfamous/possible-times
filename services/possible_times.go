package services

import (
	"fmt"
	"thirthfamous/possibility-hour/utils"
)

func PossibleTimes(digits []int) int {
	var hours = make([]string, 0)
	hours = utils.MakeHour(digits)

	// the minute
	possibleTimes := utils.AddMinutes(hours, digits)
	fmt.Println(possibleTimes)
	return len(possibleTimes)
}
