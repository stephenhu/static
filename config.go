package main;

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
)

func isExcluded(filename string) bool {
	
	conf := readConfig()

	list := strings.Split(conf[EXCLUDE], ",")

	for _, f := range list {
		if strings.Trim(filename, " ") == strings.Trim(f, " ") {
			return true
		}
	}

	return false

} // isExcluded


func readConfig() map[string]string {

	conf := map[string]string{}
	
	buf, _ := ioutil.ReadFile(CONFIG_FILE)

	if len(buf) > 0 {
		
		err := json.Unmarshal(buf, &conf)
		
		if err != nil {
			color.Red("[Error] readConfig(): %s", err)
		}
		
	}

	return conf

} // readConfig

func toCsv(files []string) string {

	ret := ""

	for i, f := range files {

		if i == 0 {
			ret = fmt.Sprintf("%s", f)
		} else {
			ret = fmt.Sprintf("%s,%s", ret, f)
		}

	}

	return ret

} // toCsv

func setExclude() {

	files 		:= []string{}
	newfiles 	:= strings.Split(*cmdExclude, ",")
	
	conf := readConfig()

 	ex, ok := conf[EXCLUDE]

	if ok {
		files = strings.Split(ex, ",")
	}

	for _, nf := range newfiles {

		add := true

		for _, of := range files {
			
			if strings.Trim(of, " ") == strings.Trim(nf, " ") {
				add = false
				break
			}

		}

		if add {
			files = append(files, strings.Trim(nf, " "))
		}

	}		
		
	conf[EXCLUDE] = toCsv(files)
		
	j, err := json.Marshal(conf)

	if err != nil {
		color.Red("[Error] setConfig(): %s", err)
	} else {

		err := ioutil.WriteFile(CONFIG_FILE, j, 0755)
		
		if err != nil {
			color.Red("[Error] setConfig(): %s", err)
		}

	}

} // setConfig

func showConfig() {

	buf, _ := ioutil.ReadFile(CONFIG_FILE)

	color.White(string(buf))
	
} // showConfig
