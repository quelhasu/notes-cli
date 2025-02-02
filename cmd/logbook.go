/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"time"

	"github.com/quelhasu/notes-cli/utils"
	"github.com/spf13/cobra"
)

// logbookCmd represents the logbook command
var logbookCmd = &cobra.Command{
	Use:   "logbook",
	Short: "Create a new logbook entry",
	Long: `Create a new logbook entry with current date, for example:
	
notes-cli logbook --category "intern"
	
This command will create a new logbook entry if it doesn't exist
for the category named "intern"`,
	Run: func(cmd *cobra.Command, args []string) {
		dt := time.Now()
		home := utils.GoEnvVariable("HOME_NOTES_CLI")

		utils.CreateDirIfNotExist(category)

		filename := dt.Format("01-02-2006") + ".md"
		file := utils.CreateFileIfNotExist(home, category+"/"+filename, "Log")
		file.Close()

		utils.OpenEditor(home + "/" + category + "/" + filename)
	},
}

func init() {
	rootCmd.AddCommand(logbookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logbookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logbookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
