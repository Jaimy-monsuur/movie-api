package dtos

type GetMovie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Language    string  `json:"language"`
	Length      int     `json:"length"`
	Year        int     `json:"year"`
	Director    string  `json:"director"`
	Actors      string  `json:"actors"`
	Plot        string  `json:"plot"`
	AVGRating   float64 `json:"avg_rating"`
	NrOfRatings int     `json:"nr_of_ratings"`
	Url         string  `json:"url"`
}

type CreateMovie struct {
	Title    string `json:"title"`
	Language string `json:"language"`
	Length   int    `json:"length"`
	Year     int    `json:"year"`
	Director string `json:"director"`
	Actors   string `json:"actors"`
	Plot     string `json:"plot"`
	Url      string `json:"url"`
}

type UpdateMovie struct {
	Title    string `json:"title"`
	Language string `json:"language"`
	Length   int    `json:"length"`
	Year     int    `json:"year"`
	Director string `json:"director"`
	Actors   string `json:"actors"`
	Plot     string `json:"plot"`
	Url      string `json:"url"`
}

type DeleteMovie struct {
	ID int `json:"id"`
}
