package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"git.deanishe.net/deanishe/awgo"
	"git.deanishe.net/deanishe/awgo/update"
	"gopkg.in/alecthomas/kingpin.v2"
)

// name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
	// kingpin and script
	app *kingpin.Application

	// application commands
	filterWebsitesCmd *kingpin.CmdClause

	// script options
	query string

	// icons
	redditIcon    = &aw.Icon{Value: "icons/reddit.png"}
	docIcon       = &aw.Icon{Value: "icons/doc.png"}
	gitubIcon     = &aw.Icon{Value: "icons/github.png"}
	forumsIcon    = &aw.Icon{Value: "icons/forums.png"}
	translateIcon = &aw.Icon{Value: "icons/translate.png"}
	stackIcon     = &aw.Icon{Value: "icons/stack.png"}
	iconAvailable = &aw.Icon{Value: "icons/update-available.png"}

	repo = "nikitavoloboev/alfred-web-searches"

	// workflow
	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo))

	app = kingpin.New("web-searches", "Search through customised list of websites")
	app.HelpFlag.Short('h')
	app.Version(wf.Version())

	filterWebsitesCmd = app.Command("websites", "filters websites")

	for _, cmd := range []*kingpin.CmdClause{filterWebsitesCmd} {
		cmd.Flag("query", "search query").Short('q').StringVar(&query)
	}

	// list action commands
	app.DefaultEnvars()
}

// _actions
// fills Alfred with hash map values and shows keys
func filterWebsites(links map[string]string) {

	var re1 = regexp.MustCompile(`.: `)
	var re2 = regexp.MustCompile(`(all)`)

	for key, value := range links {
		if strings.Contains(key, "r: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(redditIcon).Var("RECENT", re2.ReplaceAllString(value, `week`)).Subtitle("⌃ = search past week")
		} else if strings.Contains(key, "d: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(docIcon)
		} else if strings.Contains(key, "g: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(gitubIcon)
		} else if strings.Contains(key, "s: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(stackIcon)
		} else if strings.Contains(key, "f: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(forumsIcon)
		} else if strings.Contains(key, "t: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(translateIcon)
		} else {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key)
		}
	}
	wf.Filter(query)
	wf.SendFeedback()
}

// :TODO: add auto update
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
