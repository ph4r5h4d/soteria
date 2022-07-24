package cmd

import (
	"fmt"
	"github.com/ph4r5h4d/soteria/models"
	"github.com/ph4r5h4d/soteria/pkg/files/filesPathParser"
	"github.com/ph4r5h4d/soteria/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

type dependencies struct {
	logger models.LogInterface
}

var di dependencies

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "soteria",
	Short: "I'll do my best to keep your custom configuration safe",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.soteria.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	setupDependencies()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".soteria" (without extension).
		viper.AddConfigPath(home + "/.soteria")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	//viper.SetEnvPrefix("SOTERIA_")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// setting correct file path before app start
	files, err := filesPathParser.ParseFilesPath(viper.GetStringSlice("files"))
	cobra.CheckErr(err)
	viper.Set("files", files)
}

func setupDependencies() {
	l, err := logger.BuildLogger("zap")
	if err != nil {
		cobra.CheckErr(err)
	}
	di.logger = l
}
