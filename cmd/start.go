/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/thetodd/timetrack/src/domain"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		description := ""
		projectId := 0

		if len(args) == 1 {
			description = args[0]
		} else if len(args) == 2 {
			description = args[0]
			projectId, _ = strconv.Atoi(args[1])
		}

		currentTime := time.Now()
		dirname, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		filename := dirname + "/.timetrack/" + currentTime.Format("2006-01-02") + ".json"
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			err = ioutil.WriteFile(filename, []byte{}, 0644)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		entries := domain.TimeEntries{}

		_ = json.Unmarshal([]byte(file), &entries)

		entries.TimeEntries = append(entries.TimeEntries, domain.TimeEntry{
			Start:       time.Now(),
			Description: description,
			ProjectId:   projectId,
		})

		file, _ = json.MarshalIndent(entries, "", " ")

		err = ioutil.WriteFile(filename, file, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
