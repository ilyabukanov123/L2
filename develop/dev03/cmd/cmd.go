package cmd

import (
	"bufio"
	"fmt"
	"io"
	"l2/develop/dev03/strsort"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "strsort [flags] FILE...",
	Short: "strsort - sort lines of text files",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetInt("key")
		numeric, _ := cmd.Flags().GetBool("numeric-sort")
		reverse, _ := cmd.Flags().GetBool("reverse")
		unique, _ := cmd.Flags().GetBool("unique")

		strs, err := readFiles(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		strs = strsort.Sort(strs, key, numeric, reverse, unique)

		fmt.Fprintln(os.Stdin, strings.Join(strs, "\n"))
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
	Cmd.Flags().IntP("key", "k", 0, "указание колонки для сортировки")
	Cmd.Flags().BoolP("numeric-sort", "n", false, "сортировать по числовому значению")
	Cmd.Flags().BoolP("reverse", "r", false, "сортировать в обратном порядке")
	Cmd.Flags().BoolP("unique", "u", false, "не выводить повторяющиеся строки")

}

func readFiles(files []string) ([]string, error) {
	var result []string

	for _, filepath := range files {
		file, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			line = line[:len(line)-1]
			result = append(result, line)
		}
	}
	return result, nil
}
