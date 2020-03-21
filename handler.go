package main;

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/eknkc/amber"

)

func genericHandler(w http.ResponseWriter, r *http.Request) {
	
		switch r.Method {
		case http.MethodGet:
		
			var entry string

			compiler := amber.New()
	
			location := strings.ToLower(r.URL.Path[1:])

			if location == "" {
				entry = "index"
			} else {
				entry = location
			}

			file := fmt.Sprintf("%s.amber", entry)

			parseErr := compiler.ParseFile(file)
	
			if parseErr != nil {
				
				log.Printf("[%s][Error] %s", VERSION, parseErr)
				w.WriteHeader(http.StatusInternalServerError)
				return
	
			}
	
			template, compileErr := compiler.Compile()
	
			if compileErr != nil {
				
				log.Printf("[%s][Error] %s", VERSION, compileErr)
				w.WriteHeader(http.StatusInternalServerError)
				return
	
			}
	
			template.Execute(w, nil)

		case http.MethodDelete:
		case http.MethodPost:
		case http.MethodPut:
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	
	} // genericHandler
	