ifndef GOPATH
$(warning You need to set up a GOPATH.  See the README file.)
endif

PROJECT := github.com/juju/schemagen

.PHONY: check-go check

check: check-go
	go test $(PROJECT)/...

check-go:
	$(eval GOFMT := $(strip $(shell gofmt -l .| sed -e "s/^/ /g")))
	@(if [ x$(GOFMT) != x"" ]; then \
		echo go fmt is sad: $(GOFMT); \
		exit 1; \
	fi )
	@(go tool vet -all -composites=false -copylocks=false .)

$(GOPATH)/bin/dep:
	go get -u github.com/golang/dep/cmd/dep

# populate vendor/ from Gopkg.lock without updating it first (lock file is the single source of truth for machine).
dep: $(GOPATH)/bin/dep
	$(GOPATH)/bin/dep ensure -vendor-only $(verbose)

rebuild-dependencies:
	$(GOPATH)/bin/dep ensure -v -no-vendor $(dep-update)

install-dependencies:
	@echo Installing go-1.11 snap
	@sudo snap install go --channel=1.11/stable --classic