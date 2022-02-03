package cmd

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	
	
	"github.com/djherbis/times"
	"github.com/eknkc/amber"
	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday/v2"
	"github.com/spf13/cobra"
)


type Article struct {
	ID            string            `json:"id"`
	Title					string						`json:"title"`
	Summary       string            `json:"summary"`
	Creation      string        		`json:"creation"`
	Contents      string        		`json:"contents"`
	Image         string            `json:"image"`
	Tags          map[string]int    `json:"tags"`
}


const (
	AMBER_EXT   		= "amber"
	ELLIPSES        = "..."
	EMPTY           = ""
	HEAD1           = "h1"
	IMG             = "img"
	INDEX_TEMPLATE  = "index.amber"
	INDEX_HTML      = "index.html"
	MARKDOWN				= "markdown"
	META_SRC        = "src"
	MD_EXT      		= "md"
	P               = "p"
	PWD							= "."
	README        	= "README.md"
	SPACE           = " "
	STATIC_IGNORE   = ".staticignore"
)


const (
	SUMMARY_MAX_CHAR    = 200
)


const (
	ERR_NO_TAG		  = "Error: tag does not exist"
)


var master map[string] Article


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


func init() {

	buildCmd.PersistentFlags().StringVarP(&srcDir, "source directory", "d",
		PWD, "default is the current directory")
	buildCmd.PersistentFlags().StringVarP(&outDir, "output directory", "o",
		PWD, "default is the current directory")
	buildCmd.PersistentFlags().BoolVarP(&overwrite, "overwrite", "f",
		false, "overwrite existing files")

} // init


func hash(buf []byte) string {

	s := md5.Sum(buf)

	return fmt.Sprintf("%x", s)[:32]

} // hash


func ignoreFiles() []string {

  _, err := os.Stat(STATIC_IGNORE)

  if os.IsNotExist(err) {
		return nil
  } else {
	
	/*
	file, err := os.Open(STATIC_IGNORE)

	if err != nil {
		color.Red("static: %s", err)
	} else {


	}
	*/

		return nil

  }

} // ignoreFiles


func shouldExclude(f string) bool {
	
	if(strings.ToLower(README) == strings.ToLower(f)) {
		return true
	} else {
		return false
	}

} // shouldExclude


func fileExists(filename string) bool {

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}

} // fileExists


func writeHtml(filename string, t *template.Template) {

	if len(outDir) > 0 {
		
		if !fileExists(outDir) {
			color.Red(
				"Output directory does not exist, aborting.  Please create directory.")
		}

	}

	fh, err := os.Create(filename)
	
	defer fh.Close()

	if err != nil {
		color.Red("[Error] writeHtml(): %s", err)
	} else {

		color.Green(filename)

		s := struct{
			Master map[string] Article
			Version string
		}{
			Master: master,
			Version: APP_VERSION,
		}

		err := t.Execute(fh, s)

		if err != nil {
			color.Red("[Error] writeHtml(): %s", err)
		}

	}

} // writeHtml


func getFiles() []string {

	dir := srcDir
	log.Println(dir)
	if len(dir) == 0 {
		dir = PWD
	}

	path := fmt.Sprintf("%s/*.[amber|md]", dir)

	files, err := filepath.Glob(path)

	if err != nil {
		color.Red("Error: %s", err)
		return nil
	}

	for _, f := range files {
		log.Println(f)
	}

	return files

} // getFiles


func getCreationDate(f string) string {

	if fileExists(f) {

		fh, err := times.Stat(f)

		if err != nil {
			color.Red("getCreationDate(): %s", err)
			return fmt.Sprintf("%d", time.Now().Unix())
		} else {

			if fh.HasBirthTime() {
				return fmt.Sprintf("%d", fh.BirthTime().Unix())
			} else {
				return fmt.Sprintf("%d", time.Now().Unix())
			}
			
		}

	} else {
		return fmt.Sprintf("%d", time.Now().Unix())
	}
	
} // getCreationDate


func initMaster() {
	master = make(map[string] Article)
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


func generateSummary(t string, c []byte) string {

	tmp := parseTagContents(t, c, EMPTY)

	tmp = strings.TrimSpace(tmp)

	if len(tmp) >= SUMMARY_MAX_CHAR {

		tmp = strings.TrimRight(tmp[0:SUMMARY_MAX_CHAR-1], SPACE)

		tmp = fmt.Sprintf("%s%s", tmp, ELLIPSES)

	}

	return tmp

} // generateSummary


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

				h := hash(buf)

				a := Article {
					Contents: string(content),
					Title: parseTagContents(HEAD1, content, EMPTY),
					Summary: generateSummary(P, content),
					Image: extractImage(content),
					Creation: getCreationDate(f),
					ID: h, 
				}

				master[h] = a

			}
	
		}

	}

} // extractArticles


func compile() {

	initMaster()

	extractArticles()

	buildPage()

} // compile
