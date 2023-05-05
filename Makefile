dev:
	fresh
run:
	go run app/main.go

genModel:
	go run github.com/99designs/gqlgen generate