# Usage:
# make test       	# test all modules
# make download  	# download dependencies
# make release  	# add git TAG and push
GOPATH					:= $(shell go env GOPATH)

.PHONY: download, lint, format, test, release

download:
	@for d in `find * -name 'go.mod'`; do \
		pushd `dirname $$d` >/dev/null; \
		rm -f go.sum; \
		go mod download; \
		popd >/dev/null; \
	done

lint:
	@${GOPATH}/bin/golangci-lint run ./... --deadline=5m;
	@goup -v -m  ./...

format:
	@gofmt -l -w . ;

#go test -race -v ./... || :

test: download
	@for d in `find * -name 'go.mod'`; do \
		pushd `dirname $$d` >/dev/null; \
		go test -mod=readonly  -v ./...; \
		popd >/dev/null; \
	done

release: download
	@if [ -z $(TAG) ]; then \
		echo "no  TAG. Usage: make release TAG=v0.1.1"; \
	else \
		for m in `find * -name 'go.mod' -mindepth 1 -exec dirname {} \;`; do \
			hub release create -m "$$m/${TAG} release" $$m/${TAG}; \
		done \
	fi

