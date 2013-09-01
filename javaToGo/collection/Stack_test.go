package collection

import "testing"

import (
	collection "."
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	var s collection.Stack
	s.Push("ok22")

	assert.True(t, s.Size() == 1)
	assert.Equal(t, s.Pop(), "ok22")

	assert.True(t, s.Size() == 0)

}
