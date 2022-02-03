package cmd

import (
	"fmt"
	"net/http"

	"github.com/eknkc/amber"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)


const (
	APP_ADDR			= "0.0.0.0"
	APP_PORT      = "8888"
)


const (
	ASSETS    = "/assets"
	CSS_DIR   = "/css/"
	JS_DIR    = "/js/"
	FAVICON   = "favicon.ico"
	INDEX			= "index"
)


var (
	addr  		string
	port      string
)


var testCmd = &cobra.Command{
	Use: "test",
	Short: "test site",
	Long: "test site in local server",
	Run: func(cmd *cobra.Command, args []string) {

		router := initRouter()
		fmt.Printf("Starting local webserver on %s:%s...", addr, port)

		fmt.Println(http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), router))

	},
}


func init() {

	testCmd.PersistentFlags().StringVarP(&addr, "addr", "a", APP_ADDR,
		"address to listen, defaults to 0.0.0.0")
	testCmd.PersistentFlags().StringVarP(&port, "port", "p", APP_PORT,
		"port, defaults to :8888")

} // init


func initRouter() *mux.Router {

	router := mux.NewRouter()
	
	router.HandleFunc("/", genericHandler)

	router.PathPrefix(CSS_DIR).Handler(http.FileServer(
		http.Dir(PWD)))
	router.PathPrefix(JS_DIR).Handler(http.FileServer(
		http.Dir(PWD)))

	return router

} // initRouter


func genericHandler(w http.ResponseWriter, r *http.Request) {
	
	switch r.Method {
	case http.MethodGet:

		compiler := amber.New()
		
		err := compiler.ParseFile("./index.amber")

		if err != nil {
			
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		} else {

			template, err := compiler.Compile()

			if err != nil {
				
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
	
			}
	
			initMaster()

			extractArticles()

			s := struct{
				Master [] Article
				Version string
			}{
				Master: master,
				Version: APP_VERSION,
			}

			template.Execute(w, s)
	
		}

	case http.MethodDelete:
	case http.MethodPost:
	case http.MethodPut:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

} // genericHandler
