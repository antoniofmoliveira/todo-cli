package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/antoniofmoliveira/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var undoneCmd = &cobra.Command{
	Use:   "undone <label>",
	Short: "Mark item as undone",
	Long:  `Mark item as undone`,
	Run:   undoneRun,
}

func undoneRun(cmd *cobra.Command, args []string) {
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label")
		return
	}
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Fatalln(err)
	}
	if i < 1 || i > len(items) {
		log.Fatalln("no item with label", i, "found")
	}
	items[i-1].Done = false
	fmt.Println("Marked item", i, "as undone")
	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(undoneCmd)
}
