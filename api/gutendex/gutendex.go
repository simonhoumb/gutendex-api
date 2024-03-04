package gutendex

type Books struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous string  `json:"previous"`
	Results  []Book  `json:"results"`
}

type Book struct {
	Id            int      `json:"id"`
	Title         string   `json:"title"`
	Subjects      []string `json:"subjects"`
	Authors       []Person `json:"authors"`
	Translators   []Person `json:"translators"`
	Bookshelves   []string `json:"bookshelves"`
	Languages     []string `json:"languages"`
	Copyright     bool     `json:"copyright"`
	MediaType     string   `json:"media_type"`
	DownloadCount int      `json:"download_count"`
}

type Person struct {
	BirthYear int    `json:"birth_year"`
	DeathYear int    `json:"death_year"`
	Name      string `json:"name"`
}
