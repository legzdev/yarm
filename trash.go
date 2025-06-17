package yarm

import (
	"fmt"
	"os"
	"path"
)

func MoveToTrash(target string) error {
	err := CheckTarget(target)
	if err != nil {
		return err
	}

	trashinfoFile, baseName, err := CreateTrashInfoFile(target)
	if err != nil {
		return err
	}
	defer trashinfoFile.Close()

	trashPath := path.Join(TrashDir, "files", baseName)

	if !FlagDryRun {
		moveToTrash(target, trashPath, trashinfoFile)
	}

	if FlagVerbose || FlagDryRun {
		fmt.Printf("trashed '%s' => '%s'\n", target, trashPath)
	}

	return nil
}

func moveToTrash(target, trashPath string, trashInfo *os.File) error {
	err := os.Rename(target, trashPath)
	if err != nil {
		return err
	}

	err = WriteTrashInfo(trashInfo, target)
	if err != nil {
		return err
	}

	return nil
}

func CheckTarget(target string) error {
	targetInfo, err := os.Stat(target)
	if err != nil {
		return err
	}

	if targetInfo.IsDir() && !FlagRecursive {
		return &ErrIsDir{Target: target}
	}

	return nil
}
