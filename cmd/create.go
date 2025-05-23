/*
Copyright © 2025 CORENEBULA <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func getNextID(f *os.File) int {
	// Return to start of file
	f.Seek(0, 0)

	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		return 1
	}

	// If file is empty or only has header
	if len(data) <= 1 {
		return 1
	}

	// Get last row's ID and increment
	lastRow := data[len(data)-1]
	lastID := 0
	fmt.Sscanf(lastRow[0], "%d", &lastID)

	return lastID + 1
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates And Adds A New Task",
	Long: `Creates a new task and adds it to the task lists. 
	It takes a string that will be the description of the task`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile("tasks.csv", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		// Get current time in RFC3339 format
		timestamp := time.Now().Format(time.RFC3339)

		// Create the row with the timestamp
		row := []string{
			fmt.Sprintf("%d", getNextID(file)),
			args[0],
			timestamp,
			"false",
		}

		err = writer.Write(row)
		if err != nil {
			fmt.Printf("Error writing task: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
