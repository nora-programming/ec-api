package domain

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserWithImg struct {
	User
	ImgUrl string `json:"img_url"`
}
