package cmd

import (
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding tasks",
	Long: `todocli add task name +project @kw/year

+project is optional - it will fall back to 'General'
@kw/year is currently needed - it is used to create a file under $taskdir/year/kw.todo

Planned for future is to make @kw/year optional as well - leaving it will use the current week and year.`,
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
	var kw = ""
	var year = ""

	for _, arg := range args {
		if strings.HasPrefix(arg, "+") {
			project = trimFirstRune(arg)
		} else if strings.HasPrefix(arg, "@") {
			dateinput := strings.Split(trimFirstRune(arg), "/")

			kw = dateinput[0]
			year = dateinput[1]
		} else {
			task += arg + " "
		}
	}

	if _, err := os.Stat(getYearFolder(year)); os.IsNotExist(err) {
		os.MkdirAll(getYearFolder(year), os.ModePerm)
	}

	err := ioutil.WriteFile(getFilePath(kw, year), prepareFileContent(task, project, kw, year), 0644)
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

func getFilePath(kw string, year string) string {
	return viper.GetString("taskdir") + "/" + year + "/" + kw + ".todo"
}

func prepareFileContent(task string, project string, kw string, year string) []byte {
	if _, err := os.Stat(getFilePath(kw, year)); os.IsNotExist(err) {
		return []byte(_prepareNewProject(task, project))
	}

	if !_checkForProject(project, kw, year) {
		return []byte(readFile(kw, year) + _prepareNewProject(task, project))
	}

	existingProjects := strings.Split(readFile(kw, year), "\n")
	var newFileContent = ""
	var hitProject = false

	for _, line := range existingProjects {
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "    [") && strings.HasSuffix(line, ":") {
			if hitProject {
				newFileContent += "    [] " + task + "\n"
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
		newFileContent += "    [] " + task + "\n"
	}

	return []byte(newFileContent)
}

func _prepareNewProject(task string, project string) string {
	return "\n" + project + ":\n" + "    [] " + task + "\n"
}

func _checkForProject(project string, kw string, year string) bool {
	return strings.Contains(readFile(kw, year), project+":\n")
}

func readFile(kw string, year string) string {
	content, err := ioutil.ReadFile(getFilePath(kw, year))
	check(err)

	return string(content)
}
