package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/pivotal-cf/servicescli/command"

	"github.com/jessevdk/go-flags"
	_ "github.com/pivotal-cf/ism/cli"
	_ "github.com/pivotal-cf/servicehealth/cli"
)

func main() {
	parser := flags.NewNamedParser("cli", flags.HelpFlag|flags.PassDoubleDash)

	var aggregatedCommands []command.Command
	for _, commandBuilder := range command.CommandBuilders {
		commands, err := commandBuilder()
		if err != nil {
			panic(err)
		}
		aggregatedCommands = append(aggregatedCommands, commands...)
	}

	sort.Sort(command.ByCommandLength(aggregatedCommands))

	for _, command := range aggregatedCommands {
		if len(command.Command) == 1 {
			parser.AddCommand(command.Command[0], command.ShortDescription, command.LongDescription, command.Executor)
		} else {
			parentCommand := findParent(command.Command[0], parser.Commands())
			parentCommand.AddCommand(command.Command[1], command.ShortDescription, command.LongDescription, command.Executor)
		}
	}

	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}

	_, err := parser.Parse()

	if err != nil {
		if outErr, ok := err.(*flags.Error); ok && outErr.Type == flags.ErrHelp {
			fmt.Println(err.Error())
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func findParent(parentCommand string, commands []*flags.Command) *flags.Command {
	for _, command := range commands {
		if command.Name == parentCommand {
			return command
		}
	}
	panic("cant find parent " + parentCommand)
}
