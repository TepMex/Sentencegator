package assets

//Service constants
const INCORRECT_API_KEY = "Incorrect"

//WaniKani API constants
const WK_API_URL = "http://www.wanikani.com/api/user/"
const WK_API_REQUEST_VOCAB = "/vocabulary/"
const WK_API_REQUEST_KANJI = "/kanji/"
const WK_API_REQUEST_USER_INFO = "/user-information"

//Regular expressions
const REGEXP_CONTAIN_KANJI = "[\u4E00-\u9FAF].*"

//Files (F_ - prefix)
const F_BSENTENCES_DB = "b.sentences.db"
const F_SENTENCES_DB = "sentences.db"
const F_RESULT = "result.txt"

//Output strings(O_ - prefix)
const O_INCORRECT_API = "Incorrect API key. Please check your input and try again.\n"
const O_VOCAB_ONLY = "Request vocab only for levels: %s\n"
const O_BLINES_ADDED = "B-lines will be included in your result.\n"
const O_VALID_API = "API key is valid.\n"
const O_VOCAB_LOADED = "Your vocabular loaded.\n"
const O_SENTENCES_PENDING = "Loading sentences database.\n"
const O_SENTENCES_LOADED = "Sentences loaded.\n"
const O_DONE = "Done. See your sentence list in file %s.\n"
const O_DATA_PROCESSING = "Data processing(it may take few minutes)"
const O_VOCAB_PENDING = "Loading your vocabular.\n"
const O_GREETINGS = "######################\nHello, %s of sect %s! ^_^\n######################\n"
