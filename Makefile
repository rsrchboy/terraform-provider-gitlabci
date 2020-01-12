GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

build: terraform-provider-gitlabci

terraform-provider-gitlabci: *.go gitlabci/*.go go.mod
	gofmt -w $(GOFMT_FILES)
	go build .

test: terraform-provider-gitlabci
	go test `go list ./...`

clean:
	rm -f terraform-provider-gitlabci

ci-datasource: terraform-provider-gitlabci
	cd examples/data-source-config \
	    && ln -sf ../../terraform-provider-gitlabci . \
	    && terraform init \
	    && terraform apply

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

# convenience targets for development

tfa: terraform-provider-gitlabci
	terraform init && TF_LOG=TRACE terraform apply

tfp: terraform-provider-gitlabci
	terraform init && TF_LOG=TRACE terraform plan

.PHONY: build clean ci-datasource fmt vet tfa tfp test
