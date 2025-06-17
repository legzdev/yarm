package yarm

import (
	"fmt"
	"os"
)

type ErrIsDir struct {
	Target string
}

func (e *ErrIsDir) Error() string {
	return fmt.Sprintf("yarm: cannot remove '%s': Is a directory", e.Target)
}

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
