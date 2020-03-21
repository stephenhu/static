package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"html/template"
	
	"github.com/eknkc/amber"
	"github.com/fatih/color"
	"github.com/russross/blackfriday"
)

func fileExists(filename string) bool {

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}

} // fileExists

func toHtml(filename string) string {

	var str string

	ext 	:= filepath.Ext(filename)
	base	:= filepath.Base(filename)

	switch ext {
	case AMBER:
		str = strings.TrimSuffix(base, AMBER)

	case MARKDOWN:		
		str = strings.TrimSuffix(base, MARKDOWN)

	case MD:
		str = strings.TrimSuffix(base, MD)
		
	}

	if len(*cmdOut) == 0 {
		return fmt.Sprintf("%s.html", str)	
	} else {
		return fmt.Sprintf("%s/%s.html", *cmdOut, str)
	}

} // toHtml

func writeHtml(filename string, t *template.Template) {

	if len(*cmdOut) > 0 {
		
		if !fileExists(*cmdOut) {
			color.Red(
				"Output directory does not exist, aborting.  Please create directory.")
		}

	}
	fh, err := os.Create(filename)
	
	defer fh.Close()

	if err != nil {
		color.Red("[Error] writeHtml(): %s", err)
	} else {

		color.Green("%s created...", filename)

		err := t.Execute(fh, nil)

		if err != nil {
			color.Red("[Error] writeHtml(): %s", err)
		} else {
			tagFile(filename)
		}

	}

} // writeHtml

func getFiles() []string {

	dir := *cmdSrc

	if len(dir) == 0 {
		dir = PWD
	}

	path := fmt.Sprintf("%s/*[.amber|.md]", dir)

	files, err := filepath.Glob(path)

	if err != nil {
		color.Red("Error: %s", err)
		return nil
	}

	return files

} // getFiles

func compile() {

  filenames := getFiles()

	compiler := amber.New()

	for _, f := range(filenames) {

		if isExcluded(f) {
			continue
		}

		if strings.Contains(f, AMBER) {
			
			err := compiler.ParseFile(f)

			if err != nil {
				color.Red("[Error] compile(): %s", err)
			} else {

				t, err := compiler.Compile()

				if err != nil {
					color.Red("[Error] compile(): %s", err)
				} else {

					htmlFilename := toHtml(f)

					if fileExists(htmlFilename) {

						if *cmdForce {
							writeHtml(htmlFilename, t)
						} else {
							color.Yellow(
								"File %s exists, skipping.  Use --force to overwrite....",
								htmlFilename)
						}

					} else {
						writeHtml(htmlFilename, t)
					}

				}

			}
	
		} else if strings.Contains(f, MD) || strings.Contains(f, MARKDOWN) {

			// TODO: incorporate layout.amber with body

			file, err := os.Open(f)

			if err != nil {
				color.Red("[Error] compile(): %s", err)
			} else {

				buf, err := ioutil.ReadAll(file)
				
				if err != nil {
					color.Red("[Error] compile(): %s", err)
				}
	
				_ = blackfriday.MarkdownCommon(buf)
				
			}
			
		}
		
	}

} // compile
