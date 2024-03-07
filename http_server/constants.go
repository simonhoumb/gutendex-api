package http_server

const DEFAULT_PORT = "8080"
const VERSION = "v1"

// http url paths
const EMPTY_PATH = "/"
const ROOT_PATH = "/librarystats/" + VERSION + "/"
const BOOKCOUNT_PATH = "/librarystats/" + VERSION + "/bookcount/"
const READERSHIP_PATH = "/librarystats/" + VERSION + "/readership/"
const STATUS_PATH = "/librarystats/" + VERSION + "/status/"

// resource urls
const GUTENDEXAPI_URL = "http://129.241.150.113:8000/books/"
const FALLBACK_GUTENDEXAPI_URL = "https://gutendex.com/books"
const LANGUAGEAPI_URL = "http://129.241.150.113:3000/language2countries/"
const FALLBACK_LANGUAGEAPI_URL = "https://language2countries.onrender.com/"
const COUNTRIESAPI_URL = "http://129.241.150.113:8080/v3.1"
const FALLBACK_COUNTRIESAPI_URL = "https://restcountries.com/"
