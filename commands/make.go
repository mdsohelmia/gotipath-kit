package commands

import (
	"github.com/MdSohelMia/gotipath-kit/core"
	"github.com/spf13/cobra"
)

type MakeControllerCommand struct{}

func (s *MakeControllerCommand) Short() string {
	return "Create a new controller skeleton."
}

func (s *MakeControllerCommand) Setup(cmd *cobra.Command) {
	println("Hello")
}

func (MakeControllerCommand *MakeControllerCommand) Run() core.CommandRunner {
	return func(logger core.Logger) {
		// logger.Info("make controller")
	}
}

func NewMakeControllerCommandCommand() *MakeControllerCommand {
	return &MakeControllerCommand{}
}
