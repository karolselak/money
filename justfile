@build:
	mkdir -p .bin
	go build ./cmd/money
	mv ./money ./.bin/
	rm -f log.json
	./.bin/money

install:
	go install ./cmd/money

clean:
	rm -rf .bin
	rm -f log.json

@r +args='':
	rm -f log.json
	./.bin/money {{args}}
