GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

sources = $(wildcard *.go go.mod internal/**/*.go) $(schema_tools) $(runner_structs_tools)
docs = $(wildcard docs/**/*)
doc_sources = $(wildcard templates/**/* examples/**/resource.tf examples/**/data-source.tf) \
		examples/provider/provider.tf \
		$(wildcard examples/registering-a-runner/**/*)

binary_name = terraform-provider-gitlabci

local_bin = terraform.d/plugins/registry.terraform.io/rsrchboy/gitlabci/0.0.1/linux_amd64/$(binary_name)

build: $(binary_name)

$(binary_name): $(sources)
	go build .
	rm -f .terraform.lock.hcl # ...or provider checksums mismatch

test:
	TF_ACC=1 go test -v `go list ./internal/...`

clean:
	rm -f $(binary_name) predefined_variables.md vars-data .terraform.lock.hcl
	rm -rf terraform.d/ .terraform/

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

$(docs): $(doc_sources)
	go generate

# generate docs if templates have been changed
docs: $(docs)

docs-commit: $(docs)
	git add -A docs/
	git commit -m 'Update generated docs' -- docs/

gen-third-party: gen-runner-structs

runner_structs_tools = $(wildcard tools/centrifuge/*.go)
runner_structs_files = $(wildcard third_party/**/*.go)

third-party: $(runner_structs_files)
	git diff -- third_party/

$(runner_structs_files): $(runner_structs_tools)
	go run $(runner_structs_tools)
	git diff -- third_party/

gen-runner-structs: $(runner_structs_files)
	# structs files regenerated if any changes to the tooling

schema_tools = $(wildcard tools/config-schema-gen/*.go)
schema_files = internal/provider/generated.go

$(schema_files): $(schema_tools)
	go run $(schema_tools)
	git diff -- $(schema_files)

gen-schema: $(schema_files)
	# schema files regenerated if any changes to the tooling

force-regen:
	touch $(runner_structs_tools) $(schema_tools)
	$(MAKE) gen-schema gen-runner-structs

.PHONY: gen-schema gen-runner-structs gen force-regen

# convenience targets for development

tfa: $(binary_name)
	terraform init && TF_LOG=TRACE terraform apply

tfp: $(binary_name)
	terraform init && TF_LOG=TRACE terraform plan

local: $(local_bin)

init: $(local_bin)
	terraform init

plan:
	terraform plan

$(local_bin): $(binary_name)
	mkdir -p $(dir $@)
	cp $(binary_name) $@

# generate the bits we need for the env data source
# FIXME turn this into a proper `go generate` bit

predefined_variables.md:
	curl https://gitlab.com/gitlab-org/gitlab/-/raw/master/doc/ci/variables/predefined_variables.md > $@

vars-data: predefined_variables.md
	cat predefined_variables.md | grep '^| `' | awk -F\| '{ print $$2 $$5 }' > $@

env-ds-struct: vars-data
	@cat $< | perl -nE '/`(\w+)`\s+(.*\S)\s+$$/; say q{"} . lc($$1) . qq{": {\nType: schema.TypeString,\nComputed: true,\nDescription: "$$2",\n},}'

env-ds-set: vars-data
	@cat $< | sed -e 's/^ `//; s/`.*//' | perl -nE 'chomp; say qq{d.Set("} . lc($$_) . qq{", os.Getenv("$$_"))}'

.PHONY: build clean fmt vet tfa tfp test
