.DEFAULT: build

build:
	go build -o troll .

run: build
	./troll
