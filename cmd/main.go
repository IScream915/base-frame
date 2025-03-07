package main

import (
	"base_frame/pkg/cmd"
	"base_frame/pkg/program"
)

func main() {
	if err := cmd.NewApiCmd().Exec(); err != nil {
		program.ExitWithError(err)
	}
}
