DB_URL := "postgres://postgres:+_+diyor2005+_+@localhost:5432/chat_service_user?sslmode=disable"

migrate-up:
	migrate -path db/migrations -database $(DB_URL) -verbose up

migrate-down:
	migrate -path db/migrations -database $(DB_URL) -verbose down

migrate-force:
	migrate -path db/migrations -database $(DB_URL) -verbose force 1