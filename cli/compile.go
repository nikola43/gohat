package cli

import (
	"context"
	"fmt"
)

type CompileCommand struct {
	args *Args
}

func newCompileCommand() Command {
	return Command{
		Name:        "compile",
		Description: "compile solidity code and generate abi and bin files",
		Exec:        &CompileCommand{},
	}
}

func (c *CompileCommand) ExecCommand(ctx context.Context, args []string) error {
	c.args = &Args{args}
	if len(args) == 0 {
		return ErrorFromString(fmt.Sprintf("file not found"))
	}
	fmt.Println("args", args)

	return nil
}
