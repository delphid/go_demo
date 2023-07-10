package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

type fooCommand struct {
}

func (f *fooCommand) Help() string {
	return "The help message of fooCommand."
}

func (f *fooCommand) Run(args []string) int {
	fmt.Println("fooCommand running")
	return 0
}

func (f *fooCommand) Synopsis() string {
	return "This is fooCommand."
}

func fooCommandFactory() (cli.Command, error) {
	return &fooCommand{}, nil
}

func main() {
	c := cli.NewCLI("app", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"foo": fooCommandFactory,
		//"bar": barCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
