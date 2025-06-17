package yarm

import (
	"github.com/spf13/cobra"
)

var (
	FlagDryRun    bool
	FlagForce     bool
	FlagRecursive bool
	FlagVerbose   bool
)

var cmd = &cobra.Command{
	Use: "yarm",
	Run: func(cmd *cobra.Command, args []string) {
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

	cmd.Flags().BoolVar(&FlagDryRun, "dry-run", false, "")
	cmd.Flags().BoolVarP(&FlagForce, "force", "f", false, "ignore nonexistent files and arguments, never prompt")
	cmd.Flags().BoolVarP(&FlagRecursive, "recursive", "r", false, "remove directories and their contents recursively")
	cmd.Flags().BoolVarP(&FlagRecursive, "recursive-deletion", "R", false, "alias for -r")
	cmd.Flags().BoolVarP(&FlagVerbose, "verbose", "v", false, "explain what is being done")

	cmd.Flags().MarkHidden("recursive-deletion")

	return cmd.Execute()
}
