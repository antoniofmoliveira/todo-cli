package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/antoniofmoliveira/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchCmd = &cobra.Command{
	Use:   "search <text>",
	Short: "Search todo items",
	Long:  `Search will search for todo items in the given file`,
	Run:   searchRun,
}

func searchRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	m := make(map[string]todo.Item)
	for _, x := range args {
		for _, item := range items {
			z, err := strconv.Atoi(x)
			if err == nil && z == item.Priority {
				m[item.Text] = item
			}
			if strings.Contains(strings.ToLower(item.Text), strings.ToLower(x)) || strings.Contains(item.DueDate, x) || strings.Contains(item.Created, x) {
				m[item.Text] = item
			}
		}
	}
	for _, item := range m {
		fmt.Fprintln(w, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.DueDate+"\t"+item.Created+"\t"+item.Text+"\t")

	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
