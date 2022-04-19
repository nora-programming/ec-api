package domain

type Purchase struct {
	ID         int `json:"id" gorm:"primary_key"`
	Product_id int `json:"product_id"`
	Buyer_id   int `json:"buyer_id"`
}

type Sale struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"buyer_name"`
	Price int    `json:"price"`
	Title string `json:"title"`
}
