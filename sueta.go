package pet

type Sueta struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
