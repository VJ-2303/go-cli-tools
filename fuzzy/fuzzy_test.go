package fuzzy_test

import (
	"bytes"
	"testing"

	"github.com/VJ-2303/fuzzy"
	"github.com/stretchr/testify/assert"
)

func TestFuzzy(t *testing.T) {
	buf := bytes.NewBufferString("hello guys i am vj\nvj is smart\nsl is smart\nvj is cute")

	count, lines := fuzzy.FuzzyFind(buf, "vj")

	want_count := 3
	want_lines := []int{1, 2, 3}

	assert.Equal(t, count, want_count)
	assert.Equal(t, lines, want_lines)
}
