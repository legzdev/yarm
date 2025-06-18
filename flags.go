package yarm

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	FlagDir       bool
	FlagDryRun    bool
	FlagForce     bool
	FlagRecursive bool
	FlagVerbose   bool
)

var (
	FlagVersion bool
	Version     string
)

var cmd = &cobra.Command{
	Use: "yarm",
	Run: func(cmd *cobra.Command, args []string) {
		if FlagVersion {
			fmt.Println(Version)
			os.Exit(0)
		}

		if len(args) <= 1 {
			Fatal("yarm: requires at least 1 argument")
		}

		for _, arg := range args[1:] {
			err := MoveToTrash(arg)
			if err != nil {
				Fatal(err)
			}
		}
	},
}

func ParseArguments(args []string) error {
	cmd.SetArgs(args)

	cmd.Flags().BoolVarP(&FlagDir, "dir", "d", false, "remove empty directories")
	cmd.Flags().BoolVar(&FlagDryRun, "dry-run", false, "")
	cmd.Flags().BoolVarP(&FlagForce, "force", "f", false, "ignore nonexistent files and arguments, never prompt")
	cmd.Flags().BoolVarP(&FlagRecursive, "recursive", "r", false, "remove directories and their contents recursively")
	cmd.Flags().BoolVarP(&FlagRecursive, "recursive-deletion", "R", false, "alias for -r")
	cmd.Flags().BoolVarP(&FlagVerbose, "verbose", "v", false, "explain what is being done")
	cmd.Flags().BoolVar(&FlagVersion, "version", false, "output version information and exit")

	cmd.Flags().MarkHidden("recursive-deletion")

	return cmd.Execute()
}
