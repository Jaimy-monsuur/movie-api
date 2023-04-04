package models

type Movie struct {
	Base
	Title       string `gorm:"unique"`
	Language    string
	Length      int
	Year        int
	Director    string
	Actors      string
	Plot        string
	AVGRating   float64  `gorm:"default:0"`
	NrOfRatings int      `gorm:"default:0"`
	Reviews     []Review `gorm:"foreignKey:MovieID"`
}
