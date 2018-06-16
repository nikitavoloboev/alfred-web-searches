package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// TODO:
func updateList() {
	// Get websites list from GitHub
	res, err := http.Get("https://raw.githubusercontent.com/nikitavoloboev/alfred-web-searches/master/workflow/websites.csv")
	if err != nil {
		return
	}

	defer res.Body.Close()
	// body, err := ioutil.ReadAll(res.Body)
	// TODO: Cache it
}

// parseCSV parses CSV for links and arguments.
func parseCSV() map[string]string {
	var err error

	// Load file
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

	// Holds user's search arguments and an appropriate search URL
	links := make(map[string]string)

	for _, record := range records {
		links[record[0]] = record[1]
	}

	return links
}

// doSearch searches through the websites and returns results to Alfred.
func doSearch() error {
	showUpdateStatus()

	log.Printf("query=%s", query)

	links := parseCSV()

	var re1 = regexp.MustCompile(`.: `)
	var re2 = regexp.MustCompile(`(all)`)

	for key, value := range links {
		if strings.Contains(key, "r: ") {
			wf.NewItem(key).Valid(true).Var("URL", value).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(redditIcon).Var("RECENT", re2.ReplaceAllString(value, `week`)).Subtitle("⌘ = Search past week")
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

	query = os.Args[1]

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
	return nil
}
