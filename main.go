package main

// Package is called aw
import "gogs.deanishe.net/deanishe/awgo"

func run() {
	aw.NewItem("or not")
	aw.SendFeedback()
}

func main() {
	aw.Run(run)
}
