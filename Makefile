GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

build: terraform-provider-gitlabci

terraform-provider-gitlabci: *.go gitlabci/*.go
	go build .

clean:
	rm -f terraform-provider-gitlabci

fmt:
	gofmt -w $(GOFMT_FILES)

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi
