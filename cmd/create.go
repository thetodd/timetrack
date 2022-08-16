package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/thetodd/timetrack/src/domain"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("You need to provide name and id of the project")
			os.Exit(1)
		}

		dirname, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		file, _ := ioutil.ReadFile(dirname + "/.timetrack/projects.json")

		projects := domain.ProjectList{}

		_ = json.Unmarshal([]byte(file), &projects)

		projects.Projects = append(projects.Projects, domain.Project{Name: args[0], Id: args[1]})

		file, _ = json.MarshalIndent(projects, "", " ")

		os.MkdirAll(dirname+"/.timetrack/", os.ModePerm)
		err = ioutil.WriteFile(dirname+"/.timetrack/projects.json", file, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println("Added project", args[0], "with id", args[1])
	},
}

func init() {
	projectCmd.AddCommand(createCmd)
}
