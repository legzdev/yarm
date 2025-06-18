package yarm

import (
	"fmt"
	"os"
)

type ErrCannotRemove struct {
	Target string
	Msg    string
}

func (e *ErrCannotRemove) Error() string {
	return fmt.Sprintf("yarm: cannot remove '%s': %s", e.Target, e.Msg)
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
