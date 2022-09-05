package go_url_saver

type Url struct {
	ID          int    `json:"id" db:"id"`
	Url         string `json:"url" binding:"required" db:"url"`
	Description string `json:"description" db:"description"`
	UserID      int    `json:"user_id" db:"user_id"`
}

type UrlResponse struct {
	Url         string `json:"url" binding:"required" db:"url"`
	Description string `json:"description" db:"description"`
}

type UpdateUrl struct {
	Url         *string `json:"url"`
	Description *string `json:"description"`
}
