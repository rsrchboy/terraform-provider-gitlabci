GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

build: terraform-provider-gitlabci

terraform-provider-gitlabci: *.go gitlabci/*.go go.mod helper/**/*.go internal/**/*.go
	gofmt -w $(GOFMT_FILES)
	go build .

test: terraform-provider-gitlabci
	go test `go list ./...`

clean:
	rm -f terraform-provider-gitlabci mkdoc/schema.json

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

schema.json: terraform-provider-gitlabci
	terraform init && terraform providers schema --json > schema.json

README.md: README.md.gotmpl schema.json README.yml mkdoc/*
	gomplate --file README.md.gotmpl > README.md.tmp
	doctoc --gitlab --notitle README.md.tmp
	mv README.md.tmp README.md

.PHONY: build clean ci-datasource fmt vet tfa tfp test


