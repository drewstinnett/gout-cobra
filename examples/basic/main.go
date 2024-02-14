package main

import (
	"log"

	goutbra "github.com/drewstinnett/gout-cobra"
	"github.com/drewstinnett/gout/v2"
	"github.com/spf13/cobra"
)

func newCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "basic-example",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			goutbra.Cmd(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			gout.MustPrint([]struct {
				First string
				Last  string
			}{
				{First: "betty", Last: "blue"},
				{First: "joe", Last: "schmoe"},
			})
		},
	}
	goutbra.Bind(cmd, goutbra.WithHelp("Some help"))

	return cmd
}

func main() {
	if err := newCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
