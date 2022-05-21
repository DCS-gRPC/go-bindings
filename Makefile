generate: init
	bin/generate

init: modules tools | tidy

modules:
	go mod download

tools:
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -P $$(nproc) -L 1 -tI % go install %

tidy:
	go mod tidy

.PHONY: generate init modules tools tidy
