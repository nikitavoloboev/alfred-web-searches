//
// Copyright (c) 2017 Nikita Voloboev <nikita.voloboev@gmail.com>
//
// MIT Licence. See http://opensource.org/licenses/MIT
//
// Created on 2017-05-26
//

package main

import "gogs.deanishe.net/deanishe/awgo"

var (
	iconUpdate = &aw.Icon{Value: "update.png"} // TODO: check
	minScore   = 30.0                          // Default cut-off for search results

	wfopts *aw.Options
	sopts  *aw.SortOptions
	wf     *aw.Workflow
)

func init() {
	sopts = aw.NewSortOptions()
	sopts.SeparatorBonus = 10.0
	wfopts = &aw.Options{
		GitHub:      repo,
		SortOptions: sopts,
	}
	wf = aw.NewWorkflow(wfopts)
}

const updateJobName = "checkForUpdate"

const repo = "nikitavoloboev/alfred-web-searches"

func run() {

	aw.NewItem(" test")
	aw.SendFeedback()
}

func main() {
	aw.Run(run)
}
