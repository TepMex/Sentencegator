package main

import (
	"./web_interface"
)

var WaniKaniApiKey string
var Levels = ""
var Vocabular []string
var Sentences []string
var Result []string

var includeB = false

func main() {

	web_interface.CreateWebInterface(web_interface.DEFAULT_PORT)

}
