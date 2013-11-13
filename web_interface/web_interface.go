package web_interface

import (
	"../assets"
	"../sentencegator_utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

const DEFAULT_PORT = "80"

const HOME_DIR = "webcontent/"

const FAVICON_REQUEST = "/favicon.ico"
const BACKGROUND_REQUEST = "/background.jpg"
const SENTENCES_REQUEST = "/sentences"

func requestHandler(w http.ResponseWriter, r *http.Request) {

	var file string

	switch r.RequestURI {
	case FAVICON_REQUEST:
		file = HOME_DIR + "img/favicon.ico"
	case BACKGROUND_REQUEST:
		file = HOME_DIR + "img/background.jpg"
	case SENTENCES_REQUEST:
		apik := r.FormValue("apik")
		withB := r.FormValue("bl")
		levels := r.FormValue("levels")
		Vocabular := sentencegator_utils.LoadWaniKaniData(apik, levels)
		var includeB bool
		var Sentences []string
		if withB == "false" {
			Sentences = sentencegator_utils.LoadSentencesDB(assets.F_SENTENCES_DB)
			includeB = false
		} else {
			Sentences = sentencegator_utils.LoadSentencesDB(assets.F_BSENTENCES_DB)
			includeB = true
		}
		Result := sentencegator_utils.ProcessingSentences(Sentences, Vocabular, includeB)
		sentencegator_utils.WriteLines(Result, assets.F_RESULT)
		file = "result.txt"
	default:
		file = HOME_DIR + "index.tpl"
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("404\n")
		return
	}

	w.Write(data)
}

func CreateWebInterface(port string) {
	fmt.Printf(assets.O_WEB_RUNNING, port)
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":"+port, nil)
}
