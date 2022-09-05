package go_url_saver

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name"`
	Password string `json:"password" binding:"required"`
}
