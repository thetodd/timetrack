/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/thetodd/timetrack/src/domain"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		file, err := ioutil.ReadFile(dirname + "/.timetrack/projects.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		projects := domain.ProjectList{}

		_ = json.Unmarshal([]byte(file), &projects)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name"})

		for i := 0; i < len(projects.Projects); i++ {
			table.Append([]string{projects.Projects[i].Id, projects.Projects[i].Name})
		}
		table.SetCenterSeparator("┼")
		table.SetColumnSeparator("│")
		table.SetRowSeparator("─")
		table.SetHeaderColor(tablewriter.Colors{tablewriter.FgHiYellowColor, tablewriter.Bold}, tablewriter.Colors{tablewriter.FgHiYellowColor, tablewriter.Bold})

		fmt.Println("Available projects:")
		table.Render() // Send output
	},
}

func init() {
	projectCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
