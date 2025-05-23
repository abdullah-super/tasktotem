/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Marks a task as complete",
	Long:  `Marks a task as complete. It takes a task ID as an argument.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile("tasks.csv", os.O_RDWR, 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		data, err := reader.ReadAll()
		if err != nil {
			fmt.Printf("Error reading CSV file: %v\n", err)
			return
		}

		taskID := args[0]
		var found bool

		// Find and update the task
		for i, row := range data {
			if i == 0 {
				continue
			}
			if row[0] == taskID {
				// Mark the task as complete
				row[3] = "true"
				found = true
				break
			}
		}

		// Move the check outside the loop
		if !found {
			fmt.Printf("Task with ID %s not found\n", taskID)
			return
		}

		// Truncate file before writing new content
		if err := file.Truncate(0); err != nil {
			fmt.Printf("Error truncating file: %v\n", err)
			return
		}

		// Write the updated data back to the CSV file
		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.WriteAll(data); err != nil {
			fmt.Printf("Error writing to CSV file: %v\n", err)
			return
		}

		fmt.Printf("Task with ID %s marked as complete\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
