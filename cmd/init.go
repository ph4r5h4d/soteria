package cmd

import (
	"github.com/ph4r5h4d/soteria/pkg/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the application",
	Long: `Application will read the configuration and then tries to clone the remote git
repository, in order to do so, the application uses ssh-agent, so make sure this is 
configured in your machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := git.Clone(di.logger, viper.GetString("git.repository"))
		if err == nil {
			di.logger.Info("operation completed successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
