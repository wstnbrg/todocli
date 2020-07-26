package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "Handle your config parameters",
	Long: `Handle your config parameters

get config parameters by running

todocli conf get configParameter


set config parameters by running

todocli conf set configParameter configValue

configValue can be multiple words seperated by a space


Current supported config parameters:
- taskdir		= directory to save tasks to`,
	Run: func(cmd *cobra.Command, args []string) {
		var configMode = ""
		var configPath = ""
		var configValue = ""
		var argCounter = 1
		for _, arg := range args {
			if argCounter == 1 {
				configMode = arg
			} else if argCounter == 2 {
				configPath = arg
			} else {
				configValue += arg + " "
			}
			argCounter++
		}

		if configMode == "set" {
			if configPath != "" && configValue != "" {
				viper.Set(configPath, strings.Trim(configValue, " "))
				viper.WriteConfig()
			} else {
				fmt.Println("Please enter valid arguments. Check 'todocli help conf' for help.")
			}
		} else if configMode == "get" {
			if configPath != "" {
				fmt.Println("Config path for "+configPath+" is: ", viper.Get(configPath))
			} else {
				fmt.Println("Please enter valid arguments. Check 'todocli help conf' for help.")
			}
		} else {
			fmt.Println("Please enter valid arguments. Check 'todocli help conf' for help.")
		}
	},
}

func init() {
	rootCmd.AddCommand(confCmd)
}
