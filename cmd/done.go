package cmd

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var project = "General"
		var tasknumber = 0
		for _, arg := range args {
			if strings.Contains(arg, "+") {
				project = strings.Split(trimFirstRune(arg), "+")[0]
			} else {
				var taskInt, err = strconv.Atoi(arg)
				check(err)

				tasknumber = taskInt
			}
		}

		markTastAsDone(project, tasknumber)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func markTastAsDone(project string, tasknumber int) {
	var content = strings.Split(readCurrentFile(), "\n")
	var newFileContent = ""
	var hitProject = false
	var lineNumber = 1

	for _, line := range content {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "    [x]") {
			continue
		}

		if !strings.HasPrefix(line, "    [") && strings.HasSuffix(line, ":") {
			if hitProject {
				hitProject = false
			}

			if strings.Contains(line, project) {
				hitProject = true
			}

			newFileContent += "\n" + line + "\n"
			continue
		}

		if hitProject && lineNumber == tasknumber {
			newFileContent += _markTastAsDone(line) + "\n"
			lineNumber++
			continue
		}

		newFileContent += line + "\n"

		if hitProject {
			lineNumber++
		}
	}

	err := ioutil.WriteFile(getCurrentFilePath(), []byte(newFileContent), 0644)
	check(err)
}

func _markTastAsDone(task string) string {
	return strings.Replace(task, "    [ ]", "    [x]", 1)
}
