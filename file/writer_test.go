package writer_test

import (
	"os"
	"testing"

	"github.com/VJ-2303/writer"
	"github.com/google/go-cmp/cmp"
)

func TestWriteToFile_WritesGivenDataToFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/write_test.txt"

	want := []byte{1, 2, 3, 4}
	err := writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWriteToFile_ReturnsErrorForUnwritableFile(t *testing.T) {
	t.Parallel()
	path := "notexists/write_test.txt"
	err := writer.WriteToFile(path, []byte{})
	if err == nil {
		t.Fatal("want error when file not writable")
	}
}

func TestWriteToFile_ClobbersExistingFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/clobber_test.txt"

	err := os.WriteFile(path, []byte{4, 5, 6}, 0o600)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{1, 2, 3}
	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
