package wanikani_datatypes

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

type Kanji struct {
	Character        string  `json:"character"`
	Meaning          string  `json:"meaning"`
	Onyomi           string  `json:"onyomi"`
	Kunyomi          string  `json:"kunyomi"`
	ImportantReading string  `json:"important_reading"`
	Level            float64 `json:"level"`
	Stats            struct {
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

type WKResponseKanji struct {
	UserInfo      User_info `json:"user_information"`
	RequestedInfo []Kanji   `json:"requested_information"`
}
