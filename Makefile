GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

sources = $(wildcard *.go gitlabci/*.go go.mod helper/**/*.go internal/**/*.go)

binary_name = terraform-provider-gitlabci

local_bin = terraform.d/plugins/registry.terraform.io/rsrchboy/gitlabci/1.0.0/linux_amd64/$(binary_name)

build: $(binary_name)

$(binary_name): $(sources)
	go build .

test: $(binary_name)
	go test `go list ./...`

clean:
	rm -f $(binary_name)
	rm -rf terraform.d/

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

gen:
	go generate

# convenience targets for development

tfa: $(binary_name)
	terraform init && TF_LOG=TRACE terraform apply

tfp: $(binary_name)
	terraform init && TF_LOG=TRACE terraform plan

local: $(local_bin)
	echo

init: $(local_bin)
	terraform init

$(local_bin): $(binary_name)
	mkdir -p $(dir $@)
	cp $(binary_name) $@

.PHONY: build clean ci-datasource fmt vet tfa tfp test
