package cmd

import (
	"fmt"
	"log"

	"github.com/antoniofmoliveira/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

var addCmd = &cobra.Command{
	Use:   "add <text> [flags]",
	Short: "Add a new todo item3",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}
	for _, x := range args {
		item := todo.Item{Text: x, Priority: priority}
		if dueDateOpt != "" {
			item.DueDate = dueDateOpt
		}
		if createdOpt != "" {
			item.Created = createdOpt
		}
		items = append(items, item)
	}
	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
}
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&dueDateOpt, "due", "d", "", "mm/dd/yyyy")
	addCmd.Flags().StringVarP(&createdOpt, "created", "c", "", "mm/dd/yyyy")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2 or 3")
}
