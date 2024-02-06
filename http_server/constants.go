package http_server

const VERSION = "v1"

// http url paths
const EMPTY_PATH = "/"
const DEFAULT_PATH = "/librarystats/" + VERSION + "/"
const BOOKCOUNT_PATH = "/librarystats/" + VERSION + "/bookcount/"
const READERSHIP_PATH = "/librarystats/" + VERSION + "/readership/"
const STATUS_PATH = "/librarystats/" + VERSION + "/status/"

// resource urls
const GUTENDEXAPI_URL = "http://129.241.150.113:8000/books/"
const LANGUAGEAPI_URL = "http://129.241.150.113:3000/language2countries/"
const COUNTRIESAPI_URL = "http://129.241.150.113:8080/v3.1"
