package yarm

import (
	"fmt"
	"os"
)

func Fatal(v ...any) (int, error) {
	n, err := fmt.Println(v...)
	if err != nil {
		return n, err
	}

	os.Exit(1)
	return 0, nil
}

func Fatalf(format string, v ...any) (int, error) {
	n, err := fmt.Printf(format, v...)
	if err != nil {
		return n, err
	}

	os.Exit(1)
	return 0, nil
}
