package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/antoniofmoliveira/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dueDateOpt string
var createdOpt string

var editCmd = &cobra.Command{
	Use:   "edit <label> <text>",
	Short: "Update the item with the given label",
	Long:  `Update the item with the given label`,
	Run:   editRun,
}

func updateText(args []string) error {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Fatalln(err)
	}
	if args[1] == "" {
		log.Fatalln("no text given")
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label")
	}
	if i < 1 || i > len(items) {
		log.Fatalln("no item with label", i, "found")
	}
	items[i-1].Text = args[1]
	fmt.Println("Item", i, "edited done")
	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func updatFlags(args []string) error {
	fmt.Println(priority)
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Fatalln(err)
	}
	for _, i := range args {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatalln(i, "is not a valid label")
		}
		if j < 1 || j > len(items) {
			log.Fatalln("no item with label", j, "found")
		}
		if priority >= 1 && priority <= 3 {
			items[j-1].Priority = priority
		}
		if dueDateOpt != "" {
			_, err := time.Parse("01/02/2006", dueDateOpt)
			if err != nil {
				log.Fatalln("Invalid due date format.", err)
				return nil
			}
			items[j-1].DueDate = dueDateOpt
		}
		if createdOpt != "" {
			_, err := time.Parse("01/02/2006", createdOpt)
			if err != nil {
				log.Fatalln("Invalid created date format.", err)
				return nil
			}
			items[j-1].Created = createdOpt
		}
		fmt.Println("Item", j, "edited done")
	}
	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func editRun(cmd *cobra.Command, args []string) {
	if len(args) == 2 {
		err := updateText(args)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
	if len(args) == 1 || len(args) > 2 {
		err := updatFlags(args)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVarP(&dueDateOpt, "due", "d", "", "mm/dd/yyyy")
	editCmd.Flags().StringVarP(&createdOpt, "created", "c", "", "mm/dd/yyyy")
	editCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2 or 3")

}
