package tests

import (
	"testing"
	"thirthfamous/possibility-hour/utils"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfTimeHasAppendedMustTrue(t *testing.T) {
	mustTrue := utils.IsAppended(12, []string{"12", "23", "21"})
	assert.Equal(t, mustTrue, true)
}

func TestCheckIfTimeHasAppendedMustFalse(t *testing.T) {
	mustFalse := utils.IsAppended(11, []string{"12", "23", "21"})
	assert.Equal(t, mustFalse, false)
}
