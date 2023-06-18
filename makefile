build:
	@go build -o ./bin/ott

run:
	go run .

run-nb:
	go run . -b

clean:
	@rm -rf ./bin
