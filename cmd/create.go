/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/sofiukl/todoctl/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a todo",
	Long:  `This command will create todo`,
	Run: func(cmd *cobra.Command, args []string) {

		taskTitle, _ := cmd.Flags().GetString("title")
		taskDescription, _ := cmd.Flags().GetString("description")
		task := utils.Task{
			Title:       taskTitle,
			Description: taskDescription,
		}
		fmt.Printf("Creating task %+v\n", task)

		// Storing task in file, you can call REST APIs also
		resp := utils.WriteAPI(task)
		fmt.Println("Task created with ID:", resp.TaskID)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("title", "t", "", "specify task title / heading")
	createCmd.Flags().StringP("description", "d", "", "specify task description")
}
