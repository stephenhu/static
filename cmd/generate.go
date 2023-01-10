package cmd

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/djherbis/times"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
	
)


const (
	APP_EXCLUDE							= README
)


type Conf struct {
	Exclude				[]string				`json:"exclude"`
}


var (
	exFlag   			string
)


var config Conf


var generateCmd = &cobra.Command{
	Use: "generate",
	Short: "generates static site files or meta data index",
	Long: "generates static files or meta data index",
}


func init() {

	config = Conf{}

	generateCmd.PersistentFlags().StringVarP(&exFlag, "exclude", "e",
	  APP_EXCLUDE, "defaults to " + APP_EXCLUDE)

	generateCmd.AddCommand(buildCmd)
	generateCmd.AddCommand(indexCmd)

} // init


func hash(buf []byte) string {

	s := md5.Sum(buf)

	return fmt.Sprintf("%x", s)[:32]

} // hash


func fileExists(filename string) bool {

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}

} // fileExists


func shouldExclude(f string) bool {
	
	for _, e := range config.Exclude {

		if(strings.ToLower(e) == strings.ToLower(f)) {
			return true
		}
	
	}

	return false

} // shouldExclude


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

		sort.Sort(master)

		s := struct{
			Master [] Article
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


func getMarkdownFiles() []string {

	files, err := filepath.Glob(fmt.Sprintf("%s/*.%s", PWD, MD_EXT))

	if err != nil {
		log.Println(err)
		return nil
	} else {
		return files
	}

} // getMarkdownFiles


func parseDateFromCommit(c *object.Commit) GitObject {

	lines := strings.Split(c.String(), NEWLINE)

	s := strings.Trim(lines[DATE_INDEX], DATE_PREFIX)

	s = strings.TrimSpace(s)
	
	var t int64

	tmp, err := time.Parse(GIT_DATE_FMT, s)

	if err != nil {
		log.Println(err)
		t = 0
	} else {
		t = tmp.Unix()
	}

	return GitObject{
		Creation: s,
		Timestamp: t,
	}

} // parseDateFromCommit


func getGitDates(f string) (string, int64) {

	var g GitObject

	r, err := git.PlainOpen(PWD)

	if err != nil {
		
		log.Println(err)
		return EMPTY, 0

	} else {

		logs, err := r.Log(&git.LogOptions{
			FileName: &f,
		})

		if err != nil {
			
			log.Println(err)
			return EMPTY, 0

		} else {

			err = logs.ForEach(func(c *object.Commit) error {

				g = parseDateFromCommit(c)
					
				return nil

			})

		}

		return g.Creation, g.Timestamp

	}

} // getGitDates


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


func generateSummary(t string, c []byte) string {

	tmp := parseTagContents(t, c, EMPTY)

	tmp = strings.TrimSpace(tmp)

	if len(tmp) >= SUMMARY_MAX_CHAR {

		tmp = strings.TrimRight(tmp[0:SUMMARY_MAX_CHAR-1], SPACE)

		tmp = fmt.Sprintf("%s%s", tmp, ELLIPSES)

	}

	return tmp

} // generateSummary
