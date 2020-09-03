package en

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat0(t *testing.T) {
	x := float64(0)
	assert.True(t, x == 0)
}
