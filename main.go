package main;

import(
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (

	CMD_CLEAN     = "clean"
	CMD_COMPILE   = "compile"
	CMD_CONFIG    = "config"
	CMD_INIT			= "init"
	CMD_TEST      = "test"
	CMD_VERSION		= "version"

)

const (
	AMBER   					= ".amber"
	MARKDOWN					= ".markdown"
	MD      					= ".md"
)

const (
	CLEAN_FILE        = ".static.clean"
	CONFIG_FILE       = ".static.conf"
	DEFAULT_ADDRESS   = "localhost"
	DEFAULT_PORT      = "8888"
	EXCLUDE           = "exclude"
	PWD								= "."
	VERSION 					= "0.1.0"
)

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
		http.FileServer(http.Dir("./assets"))))
	
	router.HandleFunc("/", genericHandler)
	router.HandleFunc("/{resource}", genericHandler)

	return router

} // initRouter

func main() {
	
	switch kingpin.Parse() {
	case CMD_CLEAN:
		
		color.Green("Deleting files...")
		cleanFiles()

	case CMD_CONFIG:
		
		if len(*cmdExclude) == 0 || *cmdList {
			showConfig()			
		} else {
			setExclude()
		}

	case CMD_INIT:
		color.Green("Creating static site contents...")
	
	case CMD_TEST:

		color.Green("Please open your browser to %s...", address())

		router := initRouter()

		color.Red("Fatal: %s", http.ListenAndServe(address(), router))

	case CMD_VERSION:
		color.Green("static v%s", VERSION)
	
	case CMD_COMPILE:
		color.White("Compiling...")
		compile()
	}
	
} // main
