package main

import (
	"os"

	"github.com/legzdev/yarm"
)

func main() {
	err := yarm.CheckTrashDir()
	if err != nil {
		yarm.Fatal(err)
		return
	}

	err = yarm.ParseArguments(os.Args)
	if err != nil {
		yarm.Fatal(err)
		return
	}
}
