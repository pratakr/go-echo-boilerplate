dev:
	fresh
run:
	go run app/main.go

# Generate model from database
gen: 
	go run ./genModel.go