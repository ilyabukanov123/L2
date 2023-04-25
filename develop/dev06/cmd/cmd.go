package cmd

import (
	"bufio"
	"fmt"
	"io"
	"l2/develop/dev06/cut"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "cut [flags] TEXT",
	Short: "cut - Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные",
	Run: func(cmd *cobra.Command, args []string) {
		delimiter, _ := cmd.Flags().GetString("delimiter")
		fields, _ := cmd.Flags().GetUint("fields")
		separated, _ := cmd.Flags().GetBool("separated")

		reader := bufio.NewReader(os.Stdin)
		for {
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			result := cut.Cut(str[:len(str)-1], delimiter, fields, separated)

			fmt.Println(result)
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
	Cmd.Flags().BoolP("help", "h", false, "помощь по strsort")
	Cmd.Flags().UintP("fields", "f", 0, "выбрать поля (колонки). 0 по умолчанию")
	Cmd.Flags().StringP("delimiter", "d", "\t", "использовать другой разделитель")
	Cmd.Flags().BoolP("separated", "s", false, "только строки c разделителем")

}
