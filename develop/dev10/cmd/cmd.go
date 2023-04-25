package cmd

import (
	"fmt"
	"l2/develop/dev10/telnet"
	"os"

	"github.com/spf13/cobra"
)

var telnetArgs telnet.Args

var Cmd = &cobra.Command{
	Use:   "go-telnet [flags] HOST PORT",
	Short: "go-telnet - подключается к серверу",
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		err := telnet.Run(args[0], args[1], telnetArgs)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func Execute() {
	err := Cmd.Execute()
	if err != nil {
		Cmd.Help()
		os.Exit(1)
	}
}

func init() {
	Cmd.Flags().BoolP("help", "h", false, "помощь по go-telnet")
	Cmd.Flags().UintVar(&telnetArgs.Timeout, "timeout", 0, "таймаут на подключение")

}
