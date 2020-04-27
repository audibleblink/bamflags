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

func TestSignificance(t *testing.T) {
	places := []int{1, 2, 3, 4, 5, 6, 7, 8}
	values := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for idx, col := range places {
		got := Significance(col)
		expect := values[idx]
		assert.Equal(t, expect, got)
	}
}

func TestContains(t *testing.T) {
	input := int64(514)
	expect := true
	got, err := Contains(input, int64(2))

	assert.Equal(t, expect, got)
	assert.Nil(t, err)

	input = int64(520)
	expect = false
	got, err = Contains(input, int64(2))

	assert.Equal(t, expect, got)
	assert.Nil(t, err)
}
