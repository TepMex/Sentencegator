package web_interface

import (
	"../assets"
	"../kanjistats"
	"../sentencegator_utils"
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
)

const DEFAULT_PORT = "80"

const HOME_DIR = "webcontent/"

const FAVICON_REQUEST = "/favicon.ico"
const BACKGROUND_REQUEST = "/background.jpg"
const SENTENCES_REQUEST = "/sentences"
const SENTENCEGATOR_INTERFACE_REQUEST = "/sentencegator"
const KANJI_INTERFACE_REQUEST = "/kanjistats"
const STATS_REQUEST = "/stats"

func requestHandler(w http.ResponseWriter, r *http.Request) {

	var file string
	switch r.RequestURI {
	case FAVICON_REQUEST:
		file = HOME_DIR + "img/favicon.ico"
	case BACKGROUND_REQUEST:
		file = HOME_DIR + "img/background.jpg"
	case SENTENCES_REQUEST:
		apik := r.FormValue("apik")
		if len(apik) != 32 {
			file = HOME_DIR + "index.tpl"
		}
		withB := r.FormValue("bl")
		levels := r.FormValue("levels")
		Vocabular := sentencegator_utils.LoadWaniKaniVocabData(apik, levels)
		var includeB bool
		var Sentences []string
		if withB != "true" {
			Sentences = sentencegator_utils.LoadSentencesDB(assets.F_SENTENCES_DB)
			includeB = false
		} else {
			Sentences = sentencegator_utils.LoadSentencesDB(assets.F_BSENTENCES_DB)
			includeB = true
		}
		Result := sentencegator_utils.ProcessingSentences(Sentences, Vocabular, includeB)
		sentencegator_utils.WriteLines(Result, assets.F_RESULT)
		file = "result.txt"
	case SENTENCEGATOR_INTERFACE_REQUEST:
		file = HOME_DIR + "sentencegator/sentencegator.tpl"
	case KANJI_INTERFACE_REQUEST:
		file = HOME_DIR + "kanjistats/kanjistats.tpl"
	case STATS_REQUEST:
		inputFile, _, _ := r.FormFile("japfile")
		apik := r.FormValue("apik")
		levels := r.FormValue("levels")
		var lines []string
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		var statsRes []string
		statsRes = kanjistats.GetKanjiStats(apik, levels, lines)
		sentencegator_utils.WriteLines(statsRes, "kanjistats.txt")
		file = "kanjistats.txt"
	default:
		file = HOME_DIR + "index.tpl"
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error\n")
	}

	w.Write(data)
}

func CreateWebInterface(port string) {
	fmt.Printf(assets.O_WEB_RUNNING, port)
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":"+port, nil)
}
