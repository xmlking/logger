

.PHONY: download, test

download:
	@for d in `find * -name 'go.mod'`; do \
		pushd `dirname $$d` >/dev/null; \
		rm -f go.sum; \
		go mod download; \
		popd >/dev/null; \
	done

#go test -race -v ./... || :

test:
	@for d in `find * -name 'go.mod'`; do \
		pushd `dirname $$d` >/dev/null; \
		go mod download; \
		go test -v ./...; \
		popd >/dev/null; \
	done

