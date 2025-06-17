package yarm

import (
	"os"
	"path"
)

var TrashDir string

func CheckTrashDir() error {
	XDGDataHome := os.Getenv("XDG_DATA_HOME")

	if XDGDataHome != "" {
		TrashDir = path.Join(XDGDataHome, "Trash")
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		TrashDir = path.Join(home, ".local", "share", "Trash")
	}

	err := os.MkdirAll(path.Join(TrashDir, "files"), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Join(TrashDir, "info"), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
