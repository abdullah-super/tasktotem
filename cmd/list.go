/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var (
	allTasks bool
)

func readTasksFromCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening task data: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading task data: %w", err)
	}
	return data, nil
}

func displayTasks(w *tabwriter.Writer, data [][]string, allTasks bool) {
	// Print header
	if allTasks {
		fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
	} else {
		fmt.Fprintln(w, "ID\tTask\tCreated")
	}

	// Print data
	for i, row := range data {
		if i == 0 { // Skip header row
			continue
		}

		// Skip incomplete tasks when allTasks is false
		if !allTasks && row[3] != "false" {
			continue
		}

		createdAt, err := time.Parse(time.RFC3339, row[2])
		if err != nil {
			fmt.Printf("Error parsing time for task %s: %v\n", row[0], err)
			continue
		}

		timeAgo := timediff.TimeDiff(createdAt)

		if allTasks {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", row[0], row[1], timeAgo, row[3])
		} else {
			fmt.Fprintf(w, "%s\t%s\t%s\n", row[0], row[1], timeAgo)
		}
	}
}

// listCmd represents the list commands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all tasks",
	Long: `It will list all task if the -a flag is specified,
and if its not given then,
it will show only the completed tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := readTasksFromCSV("tasks.csv")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 8, 8, 2, ' ', 0)
		defer w.Flush()

		displayTasks(w, data, allTasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&allTasks, "all", "a", false, "Show all tasks (including incomplete ones)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
