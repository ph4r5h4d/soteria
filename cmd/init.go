package cmd

import (
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the application",
	Long: `Application will read the configuration and then tries to clone the remote git
repository, in order to do so, the application uses ssh-agent, so make sure this is 
configured in your machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := di.storage.Init()
		if err != nil {
			di.logger.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
