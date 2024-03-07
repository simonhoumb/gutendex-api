package gutendex

type Books struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous string  `json:"previous"`
	Results  *[]Book `json:"results"`
}

type Book struct {
	Id        int      `json:"id"`
	Title     string   `json:"title"`
	Authors   []Person `json:"authors"`
	Languages []string `json:"languages"`
}

type Person struct {
	BirthYear int    `json:"birth_year"`
	DeathYear int    `json:"death_year"`
	Name      string `json:"name"`
}
