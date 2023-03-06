build:
	go build -ldflags "-s -w -X 'main.$(shell grep -v '^#' .env | xargs)'" -o naming -trimpath

build-linux:
	go build -ldflags "-s -w -X 'main.$(shell grep -v '^#' .env | xargs)'" -o naming -trimpath
	mv naming $(HOME)/bin

remove-apikey:
	rm -rf $(HOME)/.config/gpt-naming/
