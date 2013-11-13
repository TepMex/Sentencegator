package kanjistats

import (
	"../assets"
	"../sentencegator_utils"
	"strconv"
	"unicode/utf8"
)

func GetKanjiStats(apik string, levels string, inputFile []string) []string {

	var KanjiFromWK string
	var KanjiInText string
	var UniqKanjiInTexts string
	var TextLength int
	var KanjiPercentageInTexts string

	var UnknownKanji string
	var KanjiNotInWK string
	var UnknownKanjiInWK string
	var KnownKanji string

	var UnknownKanjiPercentage string
	var KanjiNotInWKPercentage string
	var UnknownKanjiInWKPercentage string
	var KnownKanjiPercentage string

	var slice []string
	slice = append(slice, assets.F_ALL_WK_KANJI)

	AllWKKanji, _ := sentencegator_utils.ReadInputFiles(slice)

	KanjiFromWK = sentencegator_utils.LoadWaniKaniKanjiData(apik, levels)

	KanjiInText, TextLength = sentencegator_utils.ReadInput(inputFile)
	UniqKanjiInTexts = sentencegator_utils.UniqueKanjiInString(KanjiInText)
	KanjiPercentageInTexts = strconv.FormatFloat((float64(utf8.RuneCountInString(KanjiInText)) / float64(TextLength) * 100), 'f', 1, 64)

	UnknownKanji = sentencegator_utils.KanjiDifference(UniqKanjiInTexts, KanjiFromWK)
	KanjiNotInWK = sentencegator_utils.KanjiDifference(UnknownKanji, AllWKKanji)
	UnknownKanjiInWK = sentencegator_utils.KanjiDifference(UnknownKanji, KanjiNotInWK)
	KnownKanji = sentencegator_utils.KanjiDifference(UniqKanjiInTexts, UnknownKanji)

	UnknownKanjiPercentage = strconv.FormatFloat(sentencegator_utils.KanjiPercent(UnknownKanji, KanjiInText), 'f', 1, 64)
	KanjiNotInWKPercentage = strconv.FormatFloat(sentencegator_utils.KanjiPercent(KanjiNotInWK, KanjiInText), 'f', 1, 64)
	UnknownKanjiInWKPercentage = strconv.FormatFloat(sentencegator_utils.KanjiPercent(UnknownKanjiInWK, KanjiInText), 'f', 1, 64)
	KnownKanjiPercentage = strconv.FormatFloat(sentencegator_utils.KanjiPercent(KnownKanji, KanjiInText), 'f', 1, 64)

	var result []string

	result = append(result,
		"ALL", "Count: "+strconv.FormatInt(int64(utf8.RuneCountInString(UniqKanjiInTexts)), 10), "% in text: 100", "===========", "Kanji: "+UniqKanjiInTexts, "===========",
		"UNKNOWN", "Count: "+strconv.FormatInt(int64(utf8.RuneCountInString(UnknownKanji)), 10), "% in text: "+UnknownKanjiPercentage, "===========", "Kanji: "+UnknownKanji, "===========",
		"UNKNOWN NOT FROM WK", "Count: "+strconv.FormatInt(int64(utf8.RuneCountInString(KanjiNotInWK)), 10), "% in text: "+KanjiNotInWKPercentage, "===========", "Kanji: "+KanjiNotInWK, "===========",
		"UNKNOWN FROM WK", "Count: "+strconv.FormatInt(int64(utf8.RuneCountInString(UnknownKanjiInWK)), 10), "% in text: "+UnknownKanjiInWKPercentage, "===========", "Kanji: "+UnknownKanjiInWK, "===========",
		"KNOWN", "Count: "+strconv.FormatInt(int64(utf8.RuneCountInString(KnownKanji)), 10), "% in text: "+KnownKanjiPercentage, "===========", "Kanji: "+KnownKanji, "===========",
		"Kanji % in text: "+KanjiPercentageInTexts)

	return result

}
