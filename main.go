package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const INCORRECT_API_KEY = "Incorrect"

const WK_API_URL = "http://www.wanikani.com/api/user/"
const WK_API_REQUEST_VOCAB = "/vocabulary"
const WK_API_REQUEST_USER_INFO = "/user-information"

var WaniKaniApiKey string

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
		}
	}

	if WaniKaniApiKey != INCORRECT_API_KEY {
		fmt.Print("API key is correct.\n")
		loadWaniKaniData(WaniKaniApiKey)
	} else {
		return
	}

}

func loadWaniKaniData(apik string) {

	fmt.Printf("Load your vocabular.\n")
	res, err := http.Get(WK_API_URL + apik + WK_API_REQUEST_VOCAB)
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
		Items []interface{} `json:"requested_information"`
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
		UserInfo      User_info     `json:"user_information"`
		RequestedInfo []interface{} `json:"requested_information"`
	}
	//var r Requested_info

	inp := new(WKResponse)

	jsonResp, resp_err := ioutil.ReadAll(res.Body)
	encode_err := json.Unmarshal(jsonResp, &inp)
	//ee := json.Unmarshal([]byte(inp.RequestedInfo.([]byte)), &r)
	if resp_err != nil {
		log.Fatal(resp_err)
		fmt.Printf("resperr")
	}
	if encode_err != nil {
		log.Fatal(encode_err)
		fmt.Printf("encerr")
	}
	//jsonDecoder := json.NewDecoder(res.Body)
	res.Body.Close()
	//nigga := inp.RequestedInfo.([]Vocab_item)

	fmt.Println(inp.RequestedInfo)

}

//func parseJsonResponse
