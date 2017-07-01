//
// Copyright (c) 2017 Nikita Voloboev <nikita.voloboev@gmail.com>
//
// MIT Licence. See http://opensource.org/licenses/MIT
//
// Created on 2017-05-26
//

package main

import "gogs.deanishe.net/deanishe/awgo"

func run() {
	// https://www.reddit.com/r/golang/search?q=test&restrict_sr=on&sort=relevance&t=all
	aw.NewItem("r: wow")
	aw.SendFeedback()
}

func main() {
	aw.Run(run)
}
