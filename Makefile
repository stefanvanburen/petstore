clean:
	rm -rf gen
.PHONY: clean

generate: clean
	buf generate buf.build/acme/petapis
.PHONY: generate
