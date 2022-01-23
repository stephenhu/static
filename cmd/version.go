package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	APP_VERSION				= "0.1"
)


func init() {

} // init


var versionCmd = &cobra.Command{
	Use: "version",
	Short: "static version number",
	Long: "static (opinionated static site compiler) version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("static v%s", APP_VERSION)
	},
}
