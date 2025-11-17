package writer

import (
	"flag"
	"fmt"
	"os"
)

func Main() {
	size := flag.Int("s", 0, "Size of the file in bytes")
	file := flag.String("f", "", "File name to create")
	flag.Parse()
	if *file == "" {
		fmt.Println("enter the file name")
		os.Exit(1)
	}
	err := WriteToFile(*file, make([]byte, *size))
	if err != nil {
		fmt.Println(err)
	}
}

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0o600)
	if err != nil {
		return err
	}
	return os.Chmod(path, 0o600)
}
