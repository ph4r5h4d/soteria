package cmd

import (
	"github.com/ph4r5h4d/soteria/pkg/validation"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validating your configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		validation.Validate(viper.GetStringSlice("files"), di.logger)
	},
}

func init() {
	configCmd.AddCommand(validateCmd)
}
