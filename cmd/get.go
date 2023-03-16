// Package cmd for parsing command line arguments
package cmd

import (
	command "github.com/rawdaGastan/gridify/internal/cmd"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get deployed project domains",
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		err = command.Get(debug)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
