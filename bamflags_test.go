package bamflags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	input := int64(515)
	expect := []int{512, 2, 1}
	got, err := ParseInt(input)

	assert.Equal(t, expect, got)
	assert.Nil(t, err)
}
