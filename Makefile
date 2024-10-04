#
# Makefile for aitokenizer
#

#
# Automatically extract version and build information from git and the evironment.
# This will be injected into the binary at build time via linker flags.
#
PROGRAM=aitokenizer
VERSION=$(shell git describe --tags --always --dirty || echo "v0.0.0")
REVISION=$(shell git rev-parse --short HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILDDATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)
OSARCH=$(shell echo `go env GOHOSTOS`/`go env GOHOSTARCH`)

# Git repository information 
GIT_USER_ID=nulladmin
GIT_REPO_ID=${PROGRAM}

#
# Linker flags to inject version information
#
VERSION_FLAGS = -X github.com/${GIT_USER_ID}/${GIT_REPO_ID}/go.ServiceName=${PROGRAM} \
				-X github.com/${GIT_USER_ID}/${GIT_REPO_ID}/go.Version=${VERSION} \
				-X github.com/${GIT_USER_ID}/${GIT_REPO_ID}/go.BuildDate=${BUILDDATE} \
				-X github.com/${GIT_USER_ID}/${GIT_REPO_ID}/go.Revision=${REVISION} \
				-X github.com/${GIT_USER_ID}/${GIT_REPO_ID}/go.Branch=${BRANCH} \
				-X github.com/${GIT_USER_ID}/${GIT_REPO_ID}/go.OSArch=${OSARCH}

#
# Linker flags for smaller binaries
#
#   -w Omit the DWARF symbol table.
#   -s Omit the symbol table and debug information.
#
LDFLAGS = -w -s

build:
	@echo "PROGRAM   = '${PROGRAM}'"
	@echo "VERSiON   = '${VERSION}'"
	@echo "REVISION  = '${REVISION}'"
	@echo "BRANCH    = '${BRANCH}'"
	@echo "BUILDDATE = '${BUILDDATE}'"
	@echo "OSARCH    = '${OSARCH}'"

	go mod tidy
	go fmt ./...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS} ${VERSION_FLAGS}" -o ${PROGRAM}

# Run program
run:
	./${PROGRAM}

# Run with config file
runc:
	./${PROGRAM} -c ${PROGRAM}_config.yaml

clean:
	rm ${PROGRAM}

.PHONY: build clean run runc
