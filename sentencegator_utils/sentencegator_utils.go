package sentencegator_utils

import (
	"../assets"
	"../wanikani_datatypes"
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

func FastProcessingSentences(sent []string, chunks int, vocab []string, includeB bool) []string {

	var result []string
	//var main_channel chan string = make(chan string)
	var quit_channel chan int = make(chan int)

	fmt.Printf("Fast processing\n")

	var nSentences = len(sent)
	var chunkSize = (nSentences / chunks) + 1

	if includeB && chunkSize%2 != 0 {
		chunkSize = chunkSize + 1
	}

	//go buildResult(&result, main_channel, chunks)

	var chunk []string

	for _, s := range sent {

		chunk = append(chunk, s)

		if len(chunk) == chunkSize {
			go processingSentences(chunk, vocab, includeB, &result, quit_channel)
			//fmt.Printf("Thread running.\n")
			var emptyChunk []string
			chunk = emptyChunk
			continue
		}
	}
	//fmt.Printf("Thread running.\n")
	go processingSentences(chunk, vocab, includeB, &result, quit_channel)

	var quits int = 0
	for {
		<-quit_channel
		//fmt.Printf("Thread done.\n")
		quits++
		if quits >= chunks {
			break
		}
	}

	return result

}

//func buildResult(res *[]string, ch chan string, chunks int) {

//var quits = 0

//for {
//fmt.Printf("Channel wake up.\n")
//	received := <-ch
//fmt.Printf("%s.\n", received)
//*res = append(*res, received)
//q := <-quit
//quits += q
//fmt.Printf("Thread running.\n")
//if quits >= chunks {
//	fmt.Printf("Quits.\n")
//	break
//}
//}
//}

func ProcessingSentences(sent []string, vocab []string, includeB bool) []string {

	var reslt []string

	//fmt.Printf("Sentences length: %d\n", len(sent))

	fmt.Printf(assets.O_DATA_PROCESSING)

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

		var isGoodItem = !ContainKanji(tempSentence)

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

func processingSentences(sent []string, vocab []string, includeB bool, res *[]string, quit chan int) {

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

		var isGoodItem = !ContainKanji(tempSentence)

		if isGoodItem && containVocab {
			*res = append(*res, sentence)
			//fmt.Printf("Sent.\n")
			//if includeB {
			//	ch <- sent[k+1]
			//	ch <- "\n"
			//}
		}
	}

	quit <- 1
}

func LoadWaniKaniData(apik string, levels string) []string {

	res, err := http.Get(assets.WK_API_URL + apik + assets.WK_API_REQUEST_VOCAB + levels)
	if err != nil {
		log.Fatal(err)
	}

	var inp = new(wanikani_datatypes.WKResponse)
	var inpLimited = new(wanikani_datatypes.WKResponseLimited)
	var encode_err error

	jsonResp, resp_err := ioutil.ReadAll(res.Body)

	if levels != "" {
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

	var json = new(wanikani_datatypes.WKResponse)

	if levels != "" {
		json.RequestedInfo.Items = inpLimited.RequestedInfo
		json.UserInfo = inpLimited.UserInfo
	} else {
		json = inp
	}

	result := make([]string, len(json.RequestedInfo.Items))

	for k, word := range json.RequestedInfo.Items {
		result[k] = word.Character
	}

	fmt.Printf(assets.O_GREETINGS, json.UserInfo.Username, json.UserInfo.Title)
	fmt.Printf(assets.O_VOCAB_PENDING)

	return result

}

func ReadLines(path string) ([]string, error) {
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

func WriteLines(lines []string, path string) error {
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

func LoadSentencesDB(dbfile string) []string {

	var lines []string

	lines, open_err := ReadLines(dbfile)
	if open_err != nil {
		log.Fatal(open_err)
	}

	return lines

}

func ContainKanji(arg string) bool {
	matched, merr := regexp.MatchString(assets.REGEXP_CONTAIN_KANJI, arg)
	if merr != nil {
		log.Fatal(merr)
	}
	return matched
}
