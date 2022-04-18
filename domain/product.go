package domain

type Product struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"password"`
	Creater_id  int    `json:"creater_id"`
}
