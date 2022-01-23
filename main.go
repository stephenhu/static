package main

import(
	//"fmt"
	//"net/http"

	"github.com/stephenhu/static/cmd"
	//"github.com/gorilla/mux"
)

/*
var (

	cmdClean    = kingpin.Command("clean", "Remove compiled html files only")
	cmdCleanAll = cmdClean.Flag("all",
		"Remove all compiled html and configuration files").Short('a').Bool()
	
	cmdInit			= kingpin.Command("init", "Initialize site")
	cmdVersion	=	kingpin.Command("version", "static version")

	cmdConfig   = kingpin.Command("config", "static configuration")	
	cmdExclude  = cmdConfig.Flag("exclude",
		"Comma delimited list of files to exclude from compilation").Short(
			'e').String()
	cmdList     = cmdConfig.Flag("list",
	  "List all configuration parameters.").Bool()

	cmdTest		  =	kingpin.Command("test",
		"Starts up local http server for testing")
	cmdPort			=	cmdTest.Flag("port",
		"Port used by static server").String()

	cmdCompile	=	kingpin.Command("compile", "Compile templates")
	cmdSrc			= cmdCompile.Flag(
		"src", "Source file location").Short('d').String()
	cmdOut	    = cmdCompile.Flag(
		"out", "Output location").Short('o').String()
	cmdForce    = cmdCompile.Flag(
		"force", "Overwrite if files exist?").Short('f').Bool()

)
*/

/*
func address() string {

	if len(*cmdPort) == 0 {
		return fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, DEFAULT_PORT)
	} else {
		return fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, *cmdPort)
	}
	
} // address


func initRouter() *mux.Router {

	router := mux.NewRouter()

	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets",
		http.FileServer(http.Dir("."))))
	
	router.HandleFunc("/", genericHandler)
	router.HandleFunc("/{resource}", genericHandler)

	return router

} // initRouter
*/


func main() {
	
  cmd.Execute()
	
} // main
