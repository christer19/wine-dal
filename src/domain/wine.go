package domain

// Wine struct contains information about the bottle
type Wine struct {
	Country      string `json:"country"`
	Region       string `json:"region"`
	Lage         string `json:"lage"`
	Sweetness    string `json:"sweetness"`
	SugarLevel   string `json:"sugarLevel"`
	WineType     string `json:"wineType"`
	WineColor    string `json:"wineColor"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	AlcoholLevel string `json:"alcoholLevel"`
	Vintage      string `json:"vintage"`
	ValidEAN     bool   `json:"validEAN"`
	Acidity      string `json:"acidity"`
	Winery       string `json:"winery"`
	Grape        string `json:"grape"`
	Appellation  string `json:"appellation"`
}

// WineList contains a slice of different bottles
type WineList struct {
	WineSlice []Wine `json:"wines"`
}
