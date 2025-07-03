package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const warningMsg = "NOT to be altered without THOROUGH impact assessment: used for binding, attributes and other critical functions"

func TestDefault(t *testing.T) {
	assert.Equal(t, "message", Message, warningMsg)
}
