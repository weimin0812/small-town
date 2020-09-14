package en

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	numCourses := 2
	pre := [][]int{{1, 0}}
	assert.True(t, canFinish(numCourses, pre))
	numCourses = 3
	pre = [][]int{{1, 0}, {2, 1}}
	assert.True(t, canFinish(numCourses, pre))
}
