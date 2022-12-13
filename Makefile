deploy:
	fly deploy
.PHONY: deploy

test:
	go test -race ./...
.PHONY: test
