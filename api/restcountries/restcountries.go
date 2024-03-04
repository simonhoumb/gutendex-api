package restcountries

type RestCountry struct {
	CCA2       string `json:"cca2"`
	CCA3       string `json:"cca3"`
	Population int    `json:"population"`
}
