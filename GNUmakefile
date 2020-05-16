TEST?=$$(go list ./...)
PKG_NAME=drp

default: build

build: fmtcheck
	go install

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w -s ./$(PKG_NAME)

fmtcheck:
	@echo "==> Checking source code against gofmt..."
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

test: fmtcheck
	go test $(TESTARGS) -timeout=30s -parallel=4 $(TEST)
