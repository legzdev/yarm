package yarm

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
)

func MoveToTrash(target string) error {
	info, err := os.Stat(target)
	if err != nil {
		fmt.Println(target)
		return err
	}

	if info.IsDir() && !FlagRecursive {
		return fmt.Errorf("'%s' is a directory", target)
	}

	baseName := path.Base(target)
	var trashPath string

	for {
		trashPath = path.Join(TrashDir, "files", baseName)

		_, err := os.Stat(trashPath)
		if errors.Is(err, fs.ErrNotExist) {
			break
		}

		if err != nil {
			return err
		}

		baseName = path.Base(target) + "." + GenerateRandomID()
	}

	if !FlagDryRun {
		err = os.Rename(target, trashPath)
		if err != nil {
			return err
		}

		err = WriteTrashInfo(baseName, target)
		if err != nil {
			return err
		}
	}

	if FlagVerbose || FlagDryRun {
		fmt.Printf("trashed '%s' => '%s'\n", target, trashPath)
	}

	return nil
}
