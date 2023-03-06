build:
	go build -o naming -trimpath

build-linux:
	go build -o naming -trimpath
	mv naming $(HOME)/bin

build-macos:
	go build -o naming -trimpath
	mv naming $(HOME)/bin

remove-apikey:
	rm -rf $(HOME)/.config/gpt-naming/
