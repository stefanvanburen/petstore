deploy:
	fly deploy
.PHONY: deploy

test:
	go test # only need to test the main package for now
.PHONY: test
