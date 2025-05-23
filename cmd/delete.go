/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by its ID",
	Long: `Delete a task by its ID. It takes a task ID as an argument.
	This command will remove the task from the CSV file.`,
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
				data = slices.Delete(data, i, i+1)
				found = true
				break
			}
		}

		for i, row := range data {
			if i == 0 {
				continue
			}
			row[0] = fmt.Sprintf("%d", i)
		}

		// If task not found, print a message
		// and exit without modifying the file
		if !found {
			fmt.Printf("Task with ID %s not found\n", taskID)
			return
		}

		// Truncate file before writing new content
		if err := file.Truncate(0); err != nil {
			fmt.Printf("Error truncating file: %v\n", err)
			return
		}

		// Seek back to start of file after truncating
		if _, err := file.Seek(0, 0); err != nil {
			fmt.Printf("Error seeking file: %v\n", err)
			return
		}

		// Write the updated data back to the CSV file
		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.WriteAll(data); err != nil {
			fmt.Printf("Error writing to CSV file: %v\n", err)
			return
		}

		fmt.Printf("Task with ID %s deleted\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
