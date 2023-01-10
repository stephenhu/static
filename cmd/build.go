package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	
	"github.com/eknkc/amber"
	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday/v2"
	"github.com/spf13/cobra"
)


type GitObject struct {
  Name 					string						`json:"name"`
	Creation		  string						`json:"creation"`
	Timestamp			int64						  `json:"timestamp"`
}


type Article struct {
	ID            string            `json:"id"`
	Title					string						`json:"title"`
	Summary       string            `json:"summary"`
	Creation      string        		`json:"creation"`
	Timestamp     int64            	`json:"timestamp"`
	Contents      string        		`json:"contents"`
	Image         string            `json:"image"`
	Tags          map[string]int    `json:"tags"`
}


type Articles []Article


const (
	AMBER_EXT   		= "amber"
	DATE_PREFIX     = "Date:"
	ELLIPSES        = "..."
	EMPTY           = ""
	GIT_DATE_FMT    = "Mon Jan 02 15:04:05 2006 -0700"
	HEAD1           = "h1"
	IMG             = "img"
	INDEX_TEMPLATE  = "index.amber"
	INDEX_HTML      = "index.html"
	MARKDOWN				= "markdown"
	NEWLINE         = "\n"
	META_SRC        = "src"
	MD_EXT      		= "md"
	MKD_EXT         = "mkd"
	P               = "p"
	PWD							= "."
	README        	= "README.md"
	SPACE           = " "
	STATIC_CONF   	= ".static"
)


const (
	SUMMARY_MAX_CHAR    = 200
	DATE_INDEX          = 2
)


const (
	ERR_NO_TAG		  					= "Error: tag does not exist"
	ERR_NO_MARKDOWN_FILES			= "Error: no markdown files found"
)


var master Articles


var (
	srcDir					string
	outDir					string
	overwrite 			bool
)


var buildCmd = &cobra.Command{
	Use: "build",
	Short: "build site from templates",
	Long: "build site from templates",
	Run: func(cmd *cobra.Command, args []string) {
	  compile()
	},
}


func (a Articles) Len() int { return len(a) }
func (a Articles) Less(i, j int) bool { return a[i].Timestamp > a[j].Timestamp }
func (a Articles) Swap(i, j int) { a[i], a[j] = a[j], a[i] }


func init() {

	buildCmd.PersistentFlags().StringVarP(&srcDir, "source directory", "d",
		PWD, "default is the current directory")
	buildCmd.PersistentFlags().StringVarP(&outDir, "output directory", "o",
		PWD, "default is the current directory")
	buildCmd.PersistentFlags().BoolVarP(&overwrite, "overwrite", "f",
		false, "overwrite existing files")

} // init


func initMaster() {
	master = []Article{}
} // initMaster


// expects to find all the h1 headers and take the first one
func parseTagContents(tag string, buf []byte, attr string) string {

	if len(buf) == 0 {
		color.Red("parseTitle(): content is empty and cannot be parsed")
		return ""
	} else {

		reader := bytes.NewReader(buf)

		doc, err := goquery.NewDocumentFromReader(reader)

		if err != nil {
			log.Println("parseTagContents: ", err)
			return ERR_NO_TAG
		}

    var out string

    doc.Find(tag).Each(func(i int, s *goquery.Selection) {

			// img's are embedded in p, so the text of a p is empty

			if len(attr) > 0 {

				val, exists := s.Attr(attr)

				if !exists {
					out = ERR_NO_TAG
				} else {
					out = val
				}
        
      } else {

				if len(s.Text()) > 0 {

					if len(out) == 0 {
						out = s.Text()
					}
					
				}

			}

    })

		if len(out) == 0 {
			return ERR_NO_TAG
		} else {
			return out
		}

	}

} // parseTagContents


func checkTemplates() bool {

	if fileExists(fmt.Sprintf("%s/%s", ".", "layout.amber")) &&
	  fileExists(fmt.Sprintf("%s/%s", ".", "index.amber")) {
		return true
	} else {
		return false
	}

} // checkTemplates


func buildPage() {

	if checkTemplates() {

		compiler := amber.New()

		err := compiler.ParseFile(fmt.Sprintf("%s/%s", PWD, INDEX_TEMPLATE))
	
		if err != nil {
			color.Red("buildPage(): %s", err)
		} else {
	
			buf, err := compiler.Compile()

			if err != nil {
				color.Red("buildPage(): %s", err)
			} else {
				writeHtml(fmt.Sprintf("%s/%s", PWD, INDEX_HTML), buf)
			}	

		}

	} else {
		color.Red("buildPage(): layout.amber and index.amber not found, these are required files.")
	}

} // buildPage


func extractImage(c []byte) string {

	href := parseTagContents(IMG, c, META_SRC)

	if href == ERR_NO_TAG {
		return EMPTY
	} else {
		return href
	}

} // extractImage


func extractArticles() {

	files, err := filepath.Glob(fmt.Sprintf("%s/*.%s", PWD, MD_EXT))

	if err != nil {
		log.Println(err)
		return
	}

  for _, f := range files {

		if shouldExclude(f) {
			continue
		}

		file, err := os.Open(f)

		if err != nil {
			color.Red("extractArticles(): %s", err)
		} else {
	
			buf, err := ioutil.ReadAll(file)
			
			if err != nil {
				color.Red("extractArticles(): %s", err)
			} else if len(buf) == 0 {
				color.Red("extractArticles(): markdown file contents is empty, " +
				  "skipping %s", f)
			} else {

				content := blackfriday.Run(buf)

				creation, timestamp := getGitDates(f)

				a := Article {
					Contents: string(content),
					Title: parseTagContents(HEAD1, content, EMPTY),
					Summary: generateSummary(P, content),
					Image: extractImage(content),
					Creation: creation,
					Timestamp: timestamp,
					ID: hash(buf), 
				}

				master = append(master, a)

			}
	
		}

	}

} // extractArticles


func compile() {

	initMaster()

	extractArticles()

	buildPage()

} // compile
