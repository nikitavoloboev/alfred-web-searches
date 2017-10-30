package main

import (
	"encoding/csv"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/docopt/docopt-go"
)

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var usage = `alfred-web-searches [search|check] [<query>]

Search through any website on Earth.

Usage:
	alfred-web-searches search [<query>]
	alfred-web-searches check
    alfred-web-searches -h

Options:
    -h, --help    Show this message and exit.
`

var (
	// icons
	iconAvailable = &aw.Icon{Value: "icons/update.png"}
	redditIcon    = &aw.Icon{Value: "icons/reddit.png"}
	githubIcon    = &aw.Icon{Value: "icons/github.png"}
	translateIcon = &aw.Icon{Value: "icons/translate.png"}
	forumsIcon    = &aw.Icon{Value: "icons/forums.png"}
	stackIcon     = &aw.Icon{Value: "icons/stack.png"}
	docIcon       = &aw.Icon{Value: "icons/doc.png"}

	repo = "nikitavoloboev/alfred-web-searches"
	wf   *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo))
}

func run() {

	// Pass wf.Args() to docopt because our update logic relies on
	// AwGo's magic actions.
	args, _ := docopt.Parse(usage, wf.Args(), true, wf.Version(), false, true)

	// alternate action: get available releases from remote
	if args["check"] != false {
		wf.TextErrors = true
		log.Println("checking for updates...")
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}
		return
	}

	// _script filter
	var query string
	if args["<query>"] != nil {
		query = args["<query>"].(string)
	}

	log.Printf("query=%s", query)

	// call self with "check" command if an update is due and a
	// check job isn't already running.
	if wf.UpdateCheckDue() && !aw.IsRunning(updateJobName) {
		log.Println("running update check in background...")
		cmd := exec.Command("./alfred-web-searches", "check")
		if err := aw.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("error starting update check: %s", err)
		}
	}

	if query == "" { // Only show update status if query is empty
		// Send update status to Alfred
		if wf.UpdateAvailable() {
			wf.NewItem("update available!").
				Subtitle("↩ to install").
				Autocomplete("workflow:update").
				Valid(false).
				Icon(iconAvailable)
		}
	}

	links := parseCSV()

	var re1 = regexp.MustCompile(`.: `)
	var re2 = regexp.MustCompile(`(all)`)

	for key, value := range links {
		if strings.Contains(key, "r: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(redditIcon).Var("RECENT", re2.ReplaceAllString(value, `week`)).Subtitle("⌃ = search past week")
		} else if strings.Contains(key, "d: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(docIcon)
		} else if strings.Contains(key, "g: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(githubIcon)
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

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
}

// parses CSV of links and arguments
func parseCSV() map[string]string {
	var err error

	// load values from file to a hash map
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

	return links

}

func main() {
	wf.Run(run)
}
