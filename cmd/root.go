package cmd

import "github.com/spf13/cobra"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "WechatMsgAnalyzer",
	Short: "Analyze chat messages from an Excel file",
	Long:  `Analyze chat messages from an Excel file and generate a report.`,
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(analyzeCmd)
}
