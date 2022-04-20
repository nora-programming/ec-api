package domain

type Product struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Creater_id  int    `json:"creater_id"`
}

type PurchasedProducts struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Name        string `json:"creater_name"`
}

type ProductWithImg struct {
	Product
	ImgUrl string `json:"img_url"`
}
