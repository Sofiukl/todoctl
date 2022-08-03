/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/sofiukl/todoctl/utils"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View all todo",
	Long:  `View all of your todo`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := utils.ReadAPI()
		for _, t := range tasks {
			fmt.Printf("%+v\n", t)
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
