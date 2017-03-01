package dao

type SearchResultDTO struct {
	Title      string  `json:"title" bson:"title"`
	Rating     float64 `json:"rating" bson:"rating"`
	NumReviews int     `json:"num_reviews" bson:"num_reviews"`
	Address    string  `json:"address" address"`
}
