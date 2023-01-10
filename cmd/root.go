package cmd

import (
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)


var (
	eFlag  				string
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

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(versionCmd)

	parseConfig()
	
} // init


func parseConfig() {

  _, err := os.Stat(STATIC_CONF)

  if os.IsNotExist(err) {
		config.Exclude = []string{"README.md"}
  } else {

		file, err := os.Open(STATIC_CONF)

		if err != nil {
			log.Println(err)
		} else {

			buf, err := ioutil.ReadAll(file)

			if err != nil {
				log.Println(err)
			} else {

				err := json.Unmarshal(buf, &config)
	
				if err != nil {
					log.Println(err)
				}

			}
			
		}

  }

} // parseConfig
