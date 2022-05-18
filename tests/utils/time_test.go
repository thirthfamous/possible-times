package tests

import (
	"fmt"
	"testing"
	"thirthfamous/possibility-hour/utils"

	"github.com/stretchr/testify/assert"
)

func TestMakeHour1(t *testing.T) {
	digits := []int{2, 3, 1, 5}
	collectionOfHour := utils.MakeHour(digits)
	assert.Equal(t, "23", collectionOfHour[0])
	assert.Equal(t, "21", collectionOfHour[1])
}

func TestMakeHour2(t *testing.T) {
	digits := []int{4, 2, 1, 1}
	collectionOfHour := utils.MakeHour(digits)
	fmt.Println(collectionOfHour)
	assert.Equal(t, "14", collectionOfHour[0])
	assert.Equal(t, "21", collectionOfHour[1])
}

func TestAddMinuteToHour(t *testing.T) {
	collectionOfTimes := utils.AddMinutes([]string{"14", "21", "12", "23", "13"}, []int{4, 2, 1, 3})
	fmt.Println(collectionOfTimes)
}
