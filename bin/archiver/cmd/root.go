package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Short: "Archiver Chukchlandia",
}

func Execute {
	if err := rootCmd.Execute(); err != nil {
		handler.handlerErr(err)
	}
}