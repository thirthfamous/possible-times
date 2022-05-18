package tests

import (
	"testing"
	"thirthfamous/possibility-hour/services"

	"github.com/stretchr/testify/assert"
)

func TestPossibleTimes(t *testing.T) {
	totalPossibility := services.PossibleTimes([]int{1, 2, 3, 4})
	assert.Equal(t, totalPossibility, 60)
}
