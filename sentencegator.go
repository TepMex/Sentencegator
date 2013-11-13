package main

import (
	"./assets"
	"./sentencegator_utils"
	"./web_interface"
	"fmt"
	"os"
	"strings"
)

var WaniKaniApiKey string
var Levels = ""
var Vocabular []string
var Sentences []string
var Result []string

var includeB = false

func main() {

	if len(os.Args) <= 1 {
		web_interface.CreateWebInterface(web_interface.DEFAULT_PORT)
	} else {
		for _, arg := range os.Args {
			pair := strings.Split(arg, "=")
			if len(pair) > 2 {
				continue
			}

			switch pair[0] {
			case "--apik":
				WaniKaniApiKey = pair[1]
				if len(WaniKaniApiKey) != 32 {
					WaniKaniApiKey = assets.INCORRECT_API_KEY
					fmt.Printf(assets.O_INCORRECT_API)
				}
			case "--levels":
				Levels = pair[1]
				fmt.Printf(assets.O_VOCAB_ONLY, Levels)
			case "-b":
				includeB = true
				fmt.Printf(assets.O_BLINES_ADDED)
			}
		}

		if WaniKaniApiKey != assets.INCORRECT_API_KEY {
			fmt.Print(assets.O_VALID_API)
			Vocabular = sentencegator_utils.LoadWaniKaniData(WaniKaniApiKey, Levels)
			fmt.Printf(assets.O_VOCAB_LOADED)
		} else {
			return
		}

		fmt.Printf(assets.O_SENTENCES_PENDING)

		if !includeB {
			Sentences = sentencegator_utils.LoadSentencesDB(assets.F_SENTENCES_DB)
		} else {
			Sentences = sentencegator_utils.LoadSentencesDB(assets.F_BSENTENCES_DB)
		}

		fmt.Printf(assets.O_SENTENCES_LOADED)

		Result = sentencegator_utils.ProcessingSentences(Sentences, Vocabular, includeB)

		fmt.Printf(assets.O_DONE, assets.F_RESULT)

		sentencegator_utils.WriteLines(Result, assets.F_RESULT)
	}

}
