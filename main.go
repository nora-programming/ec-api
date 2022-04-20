package main

import (
	"github.com/joho/godotenv"
	"github.com/nora-programming/ec-api/infrastructure"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}
