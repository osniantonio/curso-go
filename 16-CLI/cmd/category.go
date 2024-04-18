/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("category called with name: " + category)
		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println("category called with exists: " + fmt.Sprint(exists))
		id, _ := cmd.Flags().GetInt16("id")
		fmt.Println("category called with id: " + fmt.Sprint(id))
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado antes do Run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado depois do Run")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("ocorreu um erro")
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "Category name")

	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if category exists")
}
