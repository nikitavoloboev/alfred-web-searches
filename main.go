package main

import (
	"fmt"
	"log"
	"net/url"

	"gogs.deanishe.net/deanishe/awgo"
	"gopkg.in/alecthomas/kingpin.v2"
)

type URL struct {
	url  string // url to search for
	name string // name in Alfred
}

// urlList holds user's list of websites to search for
var urlList []string

var (
	// kingpin and script options
	app *kingpin.Application

	// application commands
	filterWebsitesCmd, openLinkCmd *kingpin.CmdClause

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

	filterWebsitesCmd = app.Command("websites", "List websites").Alias("la")
	openLinkCmd = app.Command("openLink", "openLink").Alias("ol")

	// list action commands
	app.DefaultEnvars()
}

// _actions
func openLink(url string) {
	log.Printf(url)
}

func filterWebsites(urls []string) {
	fmt.Println(urls)
}

func run() {
	var err error

	// URLs in the list
	urlList = append(urlList, "https://github.com/search?utf8=✓&q=test&type=")
	urlList = append(urlList, "https://www.reddit.com/search?q=test&sort=relevance&t=all")

	cmd, err := app.Parse(wf.Args())
	if err != nil {
		wf.FatalError(err)
	}

	wf.MaxResults = maxResults

	switch cmd {
	case openLinkCmd.FullCommand():
		openLink(urlList[0])
	case filterWebsitesCmd.FullCommand():
		filterWebsites(urlList)
	default:
		err = fmt.Errorf("unknown command: %s", cmd)
	}

	if err != nil {
		wf.FatalError(err)
	}

	// https://www.reddit.com/r/golang/search?q=test&restrict_sr=on&sort=relevance&t=all
	// aw.NewItem("https://github.com").Valid(true)
	// aw.SendFeedback()
}

func main() {
	aw.Run(run)
}
