package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	// "strings"

	"gogs.deanishe.net/deanishe/awgo"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// kingpin and script options
	app *kingpin.Application

	// application commands
	filterWebsitesCmd *kingpin.CmdClause

	// script options (populated by kingpin application)
	searchURL  *url.URL
	maxResults int

	// workflow
	wf *aw.Workflow
)

// sets up kingpin flags
func init() {
	wf = aw.NewWorkflow(nil)

	app = kingpin.New("web-searches", "Search through customised list of websites")
	app.HelpFlag.Short('h')
	app.Version(wf.Version())

	filterWebsitesCmd = app.Command("websites", "filters websites").Alias("fl")

	// list action commands
	app.DefaultEnvars()
}

// _actions

// fills Alfred with hash map values and shows keys
func filterWebsites(links map[string]string) {

	var re = regexp.MustCompile(`.: `)

	for key, value := range links {
		wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re.ReplaceAllString(key, ``)).SortKey(key)
		log.Println(key)
		log.Println(value)
	}
	wf.SendFeedback()
}

func run() {
	var err error

	// load values from websites.csv to a hash map
	f, err := os.Open("websites.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// holds user's search arguments and an appropriate search URL
	links := make(map[string]string)

	for _, record := range records {
		links[record[0]] = record[1]
	}

	// _arg parsing
	cmd, err := app.Parse(wf.Args())
	if err != nil {
		wf.FatalError(err)
	}

	wf.MaxResults = maxResults

	switch cmd {
	case filterWebsitesCmd.FullCommand():
		filterWebsites(links)
	default:
		err = fmt.Errorf("unknown command: %s", cmd)
	}

	if err != nil {
		wf.FatalError(err)
	}
}

func main() {
	aw.Run(run)
}
