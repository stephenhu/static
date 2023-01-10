package cmd

import (
	"encoding/json"
	"log"
	"os"
	"sort"

	"github.com/spf13/cobra"
)


const (
	APP_INDEX								= "meta.json"
)


type Meta struct {
  ID							string		`json:"id"`
	Creation				string		`json:"creation"`
  Timestamp       int64     `json:"timestamp"`
	Image           string    `json:"image"`
	Author          string    `json:"author"`
}


type Index struct {
  Content         []Meta    `json:"content"`
	Version         string    `json:"version"`
}


type Metas []Meta


var index Metas

var (
	metaFile  		string
)


var indexCmd = &cobra.Command{
	Use: "index",
	Short: "Creates index file",
	Long: "Creates index of markdown files and creation dates",
	Run: func(cmd *cobra.Command, args []string) {
		createIndex()
	},
}


func (m Metas) Len() int { return len(m) }
func (m Metas) Less(i, j int) bool { return m[i].Timestamp > m[j].Timestamp }
func (m Metas) Swap(i, j int) { m[i], m[j] = m[j], m[i] }


func init() {

	index = Metas{}

	indexCmd.PersistentFlags().StringVarP(&metaFile, "index", "m", APP_INDEX,
		"defaults to " + APP_INDEX)

} // init


func writeIndex() {

	sort.Sort(index)

	s := Index{
		Content: index,
		Version: APP_VERSION,
	}

	j, err := json.Marshal(s)

	if err != nil {
		log.Println(err)
	} else {

		err := os.WriteFile(APP_INDEX, j, 0644)

		if err != nil {
			log.Println(err)
		}

	}

} // writeIndex


func createIndex() {

	files := getMarkdownFiles()

	log.Println(files)

	index = []Meta{}

	for _, f := range files {

		if !shouldExclude(f) {

			_, s := getGitDates(f)

			m := Meta{
				ID: f,
				Timestamp: s,
			}

		  index = append(index, m)

		}

	}

	writeIndex()

} // createIndex
