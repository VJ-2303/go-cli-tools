package count_test

import (
	"bytes"
	"testing"

	"github.com/VJ-2303/count"
)

func TestLinesCountsLinesInInput(t *testing.T) {
	t.Parallel()

	buf := bytes.NewBufferString("1\n2\n3\n")

	c, _ := count.NewCounter(count.WithInput(buf))

	want := 3

	got := c.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
