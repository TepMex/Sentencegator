package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const INCORRECT_API_KEY = "Incorrect"

const WK_API_URL = "http://www.wanikani.com/api/user/"
const WK_API_REQUEST_VOCAB = "/vocabulary/"
const WK_API_REQUEST_KANJI = "/kanji/"
const WK_API_REQUEST_USER_INFO = "/user-information"

const REGEXP_CONTAIN_KANJI = "[\u4E00-\u9FAF].*"

const SENTENCES_DB_FILENAME = "sentences.db"

var WaniKaniApiKey string
var Levels = ""
var Vocabular []string
var Sentences []string
var Result []string

var includeB = false

func main() {

	for _, arg := range os.Args {
		pair := strings.Split(arg, "=")
		if len(pair) > 2 {
			continue
		}

		switch pair[0] {
		case "--apik":
			WaniKaniApiKey = pair[1]
			if len(WaniKaniApiKey) != 32 {
				WaniKaniApiKey = INCORRECT_API_KEY
				fmt.Printf("Incorrect API key. Please check your input and try again.\n")
			}
		case "--levels":
			Levels = pair[1]
			fmt.Printf("Request vocab only for levels: %s\n", Levels)
		case "-b":
			includeB = true
			fmt.Printf("B-lines will be included in your result.\n")
		}
	}

	if WaniKaniApiKey != INCORRECT_API_KEY {
		fmt.Print("API key is valid.\n")
		Vocabular = loadWaniKaniData(WaniKaniApiKey)
		fmt.Printf("Your vocabular loaded.\n")
	} else {
		return
	}

	fmt.Printf("Loading sentences database.\n")

	if !includeB {
		Sentences = loadSentencesDB(SENTENCES_DB_FILENAME)
	} else {
		Sentences = loadSentencesDB("b." + SENTENCES_DB_FILENAME)
	}

	fmt.Printf("Sentences loaded.\n")

	Result = processingSentences(Sentences, Vocabular)

	fmt.Printf("Done. See your sentence list in file result.txt.\n")

	writeLines(Result, "result.txt")

}

func processingSentences(sent []string, vocab []string) []string {

	var reslt []string

	fmt.Printf("Data processing(it may take few minutes)")

	for k, sentence := range sent {

		var tempSentence = sentence

		var containVocab = false

		if includeB && k%2 == 1 {
			continue
		}

		for _, word := range vocab {

			if strings.Contains(tempSentence, word) {
				containVocab = true
				tempSentence = strings.Replace(tempSentence, word, "", -1)
			}
		}

		var isGoodItem = !containKanji(tempSentence)

		if isGoodItem && containVocab {
			reslt = append(reslt, sentence)
			if includeB {
				reslt = append(reslt, sent[k+1])
				reslt = append(reslt, "\n")
			}
		}

		if k%15000 == 0 {
			fmt.Printf(".")
		}
	}

	fmt.Printf("\n")

	return reslt

}

func loadWaniKaniData(apik string) []string {

	fmt.Printf("Load your vocabular data.\n")
	res, err := http.Get(WK_API_URL + apik + WK_API_REQUEST_VOCAB + Levels)
	if err != nil {
		log.Fatal(err)
	}

	type Item_stats struct {
		Srs                    string  `json:"srs"`
		Unlocked_date          float64 `json:"unlocked_date"`
		Available_date         float64 `json:"available_date"`
		Burned                 bool    `json:"burned"`
		Burned_date            float64 `json:"burned_date"`
		Meaning_correct        float64 `json:"meaning_correct"`
		Meaning_incorrect      float64 `json:"meaning_incorrect"`
		Meaning_max_streak     float64 `json:"meaning_max_streak"`
		Meaning_current_streak float64 `json:"meaning_current_streak"`
		Reading_correct        float64 `json:"reading_correct"`
		Reading_incorrect      float64 `json:"reading_incorrect"`
		Reading_max_streak     float64 `json:"reading_max_streak"`
		Reading_current_streak float64 `json:"reading_current_streak"`
	}

	type Vocab_item struct {
		Character string  `json:"character"`
		Kana      string  `json:"kana"`
		Meaning   string  `json:"meaning"`
		Level     float64 `json:"level"`
		Stats     struct {
			Srs                    string  `json:"srs"`
			Unlocked_date          float64 `json:"unlocked_date"`
			Available_date         float64 `json:"available_date"`
			Burned                 bool    `json:"burned"`
			Burned_date            float64 `json:"burned_date"`
			Meaning_correct        float64 `json:"meaning_correct"`
			Meaning_incorrect      float64 `json:"meaning_incorrect"`
			Meaning_max_streak     float64 `json:"meaning_max_streak"`
			Meaning_current_streak float64 `json:"meaning_current_streak"`
			Reading_correct        float64 `json:"reading_correct"`
			Reading_incorrect      float64 `json:"reading_incorrect"`
			Reading_max_streak     float64 `json:"reading_max_streak"`
			Reading_current_streak float64 `json:"reading_current_streak"`
		} `json:"stats"`
	}

	type Requested_info struct {
		Items []Vocab_item `json:"general"`
	}

	type User_info struct {
		Username      string  `json:"username"`
		Gravatar      string  `json:"gravatar"`
		Level         float64 `json:"level"`
		Title         string  `json:"title"`
		About         string  `json:"about"`
		Website       string  `json:"website"`
		Twitter       string  `json:"twitter"`
		Topics_count  float64 `json:"topics_count"`
		Posts_count   float64 `json:"posts_count"`
		Creation_date float64 `json:"creation_date"`
	}

	type WKResponse struct {
		UserInfo      User_info      `json:"user_information"`
		RequestedInfo Requested_info `json:"requested_information"`
	}

	type WKResponseLimited struct {
		UserInfo      User_info    `json:"user_information"`
		RequestedInfo []Vocab_item `json:"requested_information"`
	}

	var inp = new(WKResponse)
	var inpLimited = new(WKResponseLimited)
	var encode_err error

	jsonResp, resp_err := ioutil.ReadAll(res.Body)

	if Levels != "" {
		encode_err = json.Unmarshal(jsonResp, &inpLimited)
	} else {
		encode_err = json.Unmarshal(jsonResp, &inp)
	}

	if resp_err != nil {
		log.Fatal(resp_err)
		fmt.Printf("resperr")
	}
	if encode_err != nil {
		log.Fatal(encode_err)
		fmt.Printf("encerr")
	}

	res.Body.Close()

	var json = new(WKResponse)

	if Levels != "" {
		json.RequestedInfo.Items = inpLimited.RequestedInfo
		json.UserInfo = inpLimited.UserInfo
	} else {
		json = inp
	}

	result := make([]string, len(json.RequestedInfo.Items))

	for k, word := range json.RequestedInfo.Items {
		result[k] = word.Character
	}

	fmt.Printf("Hello, %s of sect %s! ^_^\n", json.UserInfo.Username, json.UserInfo.Title)

	return result

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func loadSentencesDB(dbfile string) []string {

	var lines []string

	lines, open_err := readLines(dbfile)
	if open_err != nil {
		log.Fatal(open_err)
	}

	return lines

}

func containKanji(arg string) bool {
	matched, merr := regexp.MatchString(REGEXP_CONTAIN_KANJI, arg)
	if merr != nil {
		log.Fatal(merr)
	}
	return matched
}
