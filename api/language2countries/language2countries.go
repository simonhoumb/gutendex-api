package language2countries

type Country struct {
	Iso31661Alpha3 string `json:"ISO3166_1_Alpha_3"`
	Iso31661Alpha2 string `json:"ISO3166_1_Alpha_2"`
	OfficialName   string `json:"Official_Name"`
	RegionName     string `json:"Region_Name"`
	SubRegionName  string `json:"Sub_Region_Name"`
	Language       string `json:"Language"`
}
