package yarm

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
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

func IsDirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func InTrashDir(target string) bool {
	if !path.IsAbs(target) {
		target, _ = filepath.Abs(target)
	}

	return strings.HasPrefix(target, TrashDir)
}
