build_app:
	go build -o ./build/go-encode .

test:
	go test -v ./tests/
