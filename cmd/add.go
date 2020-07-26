package cmd

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding tasks",
	Long: `todocli add task name +project @cw/year

+project (optional) - it will fall back to 'General'
@cw/year (optional) - it will fall back to current week and year`,
	Run: func(cmd *cobra.Command, args []string) {
		addTask(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(args []string) {
	var task = ""
	var project = "General"
	var cw = ""
	var year = ""

	for _, arg := range args {
		if strings.HasPrefix(arg, "+") {
			project = trimFirstRune(arg)
		} else if strings.HasPrefix(arg, "@") {
			dateinput := strings.Split(trimFirstRune(arg), "/")

			cw = dateinput[0]
			year = dateinput[1]
		} else {
			task += arg + " "
		}
	}

	if cw == "" {
		var _, cwInt = time.Now().ISOWeek()
		cw = strconv.Itoa(cwInt)
	}

	if year == "" {
		var yearInt, _ = time.Now().ISOWeek()
		year = strconv.Itoa(yearInt)
	}

	if _, err := os.Stat(getYearFolder(year)); os.IsNotExist(err) {
		os.MkdirAll(getYearFolder(year), os.ModePerm)
	}

	err := ioutil.WriteFile(getFilePath(cw, year), prepareFileContent(task, project, cw, year), 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeLastRuneInString(s)
	return s[i:]
}

func getYearFolder(year string) string {
	return viper.GetString("taskdir") + "/" + year
}

func getFilePath(cw string, year string) string {
	return viper.GetString("taskdir") + "/" + year + "/" + cw + ".todo"
}

func prepareFileContent(task string, project string, cw string, year string) []byte {
	if _, err := os.Stat(getFilePath(cw, year)); os.IsNotExist(err) {
		return []byte(_prepareNewProject(task, project))
	}

	if !_checkForProject(project, cw, year) {
		return []byte(readFile(cw, year) + _prepareNewProject(task, project))
	}

	existingProjects := strings.Split(readFile(cw, year), "\n")
	var newFileContent = ""
	var hitProject = false

	for _, line := range existingProjects {
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "    [") && strings.HasSuffix(line, ":") {
			if hitProject {
				newFileContent += "    [ ] " + task + "\n"
				hitProject = false
			}

			if strings.Contains(line, project) {
				hitProject = true
			}

			newFileContent += "\n" + line + "\n"
			continue
		}

		newFileContent += line + "\n"
	}

	if hitProject {
		newFileContent += "    [ ] " + task + "\n"
	}

	return []byte(newFileContent)
}

func _prepareNewProject(task string, project string) string {
	return "\n" + project + ":\n" + "    [ ] " + task + "\n"
}

func _checkForProject(project string, cw string, year string) bool {
	return strings.Contains(readFile(cw, year), project+":\n")
}

func readFile(cw string, year string) string {
	content, err := ioutil.ReadFile(getFilePath(cw, year))
	check(err)

	return string(content)
}
