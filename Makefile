APP?=app
RELEASE?=0.0.1
GOOS?=linux

.PHONY: check
check: prepare_metalinter
	gometalinter --fast --vendor ./...

.PHONY: build
build: clean
	CGO_ENABLED=1 GOOS=${GOOS} go build \
	--tags "libsqlite3 linux" \
	-ldflags "-w -s" -o bin/${GOOS}/${APP}

.PHONY: clean
clean:
	@rm -f bin/${GOOS}/${APP}

.PHONY: vendor
vendor: prepare_dep
	dep ensure

HAS_DEP := $(shell command -v dep;)
HAS_METALINTER := $(shell command -v gometalinter;)

.PHONY: prepare_dep
prepare_dep:
ifndef HAS_DEP
	go get -u -v -d github.com/golang/dep/cmd/dep && \
	go install -v github.com/golang/dep/cmd/dep
endif

.PHONY: prepare_metalinter
prepare_metalinter:
ifndef HAS_METALINTER
	go get -u -v -d github.com/alecthomas/gometalinter && \
	go install -v github.com/alecthomas/gometalinter && \
	gometalinter --install --update
endif
