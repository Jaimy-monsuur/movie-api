package dtos

type GetReviewDto struct {
	ReviewID  string  `json:"reviewId" binding:"required"`
	Review    string  `json:"review" binding:"required"`
	Rating    float32 `json:"rating" binding:"required"`
	MovieName string  `json:"movieName" binding:"required"`
	UserName  string  `json:"userName" binding:"required"`
}

type CreateReviewDto struct {
	Review    string  `json:"review" binding:"required"`
	Rating    float32 `json:"rating" binding:"required"`
	MovieName string  `json:"movieName" binding:"required"`
	UserName  string  `json:"userName" binding:"required"`
	MovieID   string  `json:"movieId" binding:"required"`
	UserID    string  `json:"userId" binding:"required"`
}

type UpdateReviewDto struct {
	Review string  `json:"review" binding:"required"`
	Rating float32 `json:"rating" binding:"required"`
}

type DeleteReviewDto struct {
	ReviewID string `json:"reviewId" binding:"required"`
	UserID   string `json:"userId" binding:"required"`
}
