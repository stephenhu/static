package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
)

var (

	//conf			string

	rootCmd = &cobra.Command{
		Use: "static",
		Short: "static site tool",
		Long: "opinionated static site tool",
	}

)


func Execute() error {
	return rootCmd.Execute()
} // Execute


func init() {

	cobra.OnInitialize()

	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(versionCmd)
	
} // init
