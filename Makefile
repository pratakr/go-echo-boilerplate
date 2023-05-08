dev:
	fresh
run:
	go run app/main.go

build:
	go build -o server app/main.go

# Generate model from database
gen: 
	go run ./genModel.go