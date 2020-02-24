# Usage:
# make test       	# test all modules
# make download  	# download dependencies
# make release  	# add git TAG and push
.PHONY: download, test, release

download:
	@for d in `find * -name 'go.mod'`; do \
		pushd `dirname $$d` >/dev/null; \
		rm -f go.sum; \
		go mod download; \
		popd >/dev/null; \
	done

#go test -race -v ./... || :

test: download
	@for d in `find * -name 'go.mod'`; do \
		pushd `dirname $$d` >/dev/null; \
		go test -v ./...; \
		popd >/dev/null; \
	done

release: download
	@if [ -z $(TAG) ]; then \
		echo "no  TAG. Usage: make release TAG=v0.1.1"; \
	else \
		for m in `find * -name 'go.mod' -exec dirname {} \;`; do \
			echo hub release create -m "\"$$m/${TAG} release\"" $$m/${TAG}; \
		done \
	fi
