MIGRATION_NAME ?=  $(shell bash -c 'read -p "Name: " pwd; echo $$pwd')

db-init:
	@docker exec -it ec-api mysql -u root -ppassword -e 'DROP DATABASE IF EXISTS ec_api'
	@docker exec -it ec-api mysql -u root -ppassword -e 'CREATE DATABASE IF NOT EXISTS ec_api CHARACTER SET utf8mb4'

migrate-up:
	@migrate -database "mysql://root:password@tcp(localhost)/ec_api" -path db/migrations up

migrate-down:
	@migrate -database "mysql://root:password@tcp(localhost)/ec_api" -path db/migrations down

migrate-new:
	@migrate create -ext sql -dir db/migrations -seq $(MIGRATION_NAME)
