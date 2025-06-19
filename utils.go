package yarm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func confirm(msg string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	msg = strings.TrimSpace(msg) + " (y/N) "
	fmt.Print(msg)

	prompt, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	prompt = strings.TrimSuffix(prompt, "\n")

	return prompt == "y", nil
}

func confirmf(format string, v ...any) (bool, error) {
	return confirm(fmt.Sprintf(format, v...))
}

func verbose(msg string) {
	if FlagDryRun || FlagVerbose {
		fmt.Println(msg)
	}
}

func verbosef(format string, v ...any) {
	verbose(fmt.Sprintf(format, v...))
}
