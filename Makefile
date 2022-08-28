clean:
	rm -rf gen
.PHONY: clean

generate: clean
	buf generate buf.build/acme/petapis --include-imports
.PHONY: generate

deploy:
	fly deploy
.PHONY: deploy

test:
	go test # only need to test the main package for now
.PHONY: test
