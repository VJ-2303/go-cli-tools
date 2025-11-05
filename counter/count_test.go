package count_test

import (
	"bytes"
	"testing"

	"github.com/VJ-2303/count"
)

func TestCount(t *testing.T) {
	t.Parallel()
	c := count.NewCounter()
	c.Input = bytes.NewBufferString("1\n2\n3")

	want := 3

	got := c.Count()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
