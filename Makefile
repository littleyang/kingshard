all: build

build: piratesharding

piratesharding:
	go build -o bin/piratesharding ./cmd/piratesharding

clean:
	@rm -rf bin

test:
	go test ./go/... -race
