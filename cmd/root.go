package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todocli",
	Short: "Manage your tasks from your favorite place - the CLI",
	Long:  `Sort your tasks by project and date. And the best is: dont worry about your backlog.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(getCurrentFilePath()); os.IsNotExist(err) {
			fmt.Println("You currently have nothing todo :)")

			return
		}

		content, err := ioutil.ReadFile(getCurrentFilePath())
		if err != nil {
			panic(err)
		}

		fmt.Println(string(content))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todocli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".todocli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".todocli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading in config file")
	}
}

func getCurrentFilePath() string {
	var year, _ = time.Now().ISOWeek()
	var _, kw = time.Now().ISOWeek()

	return viper.GetString("taskdir") + "/" + strconv.Itoa(year) + "/" + strconv.Itoa(kw) + ".todo"
}
