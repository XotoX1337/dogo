# govulncheck IDs that are knowingly accepted. All three are daemon-/server-side
# Moby vulnerabilities that are not exploitable through dogo's docker client usage:
#   GO-2026-4887 - AuthZ plugin bypass on oversized request bodies (no fix released)
#   GO-2026-4883 - off-by-one in plugin privilege validation (no fix released)
#   GO-2025-3829 - firewalld reload drops bridge isolation (fixed only in docker v25,
#                  which breaks the SDK API we depend on)
# Any other vulnerability reported by govulncheck fails `make check`.
VULN_ALLOWLIST := GO-2026-4887 GO-2026-4883 GO-2025-3829

icon:
	rsrc -ico assets/dogo_cli.ico

build:
	go mod tidy
	go build -ldflags "-s -w" -o bin/ .

run:
	go run .

clean:
	go clean

compile:
	go mod tidy
	GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/dogo-linux-arm .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/dogo-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/dogo-windows-amd64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/dogo-macos-amd64 .

check:
	go install github.com/client9/misspell/cmd/misspell@latest
	misspell -error .
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	gocyclo -over 10 .
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec -quiet --severity high ./...
	go install golang.org/x/vuln/cmd/govulncheck@latest
	@echo "govulncheck (accepted/unfixable IDs: $(VULN_ALLOWLIST))"; \
	out=$$(govulncheck ./... 2>&1); \
	echo "$$out"; \
	found=$$(echo "$$out" | grep -E '^Vulnerability #' | grep -oE 'GO-[0-9]{4}-[0-9]+' | sort -u); \
	status=0; \
	for id in $$found; do \
		case " $(VULN_ALLOWLIST) " in \
			*" $$id "*) echo "  accepted (allowlisted): $$id" ;; \
			*) echo "  NEW vulnerability not in allowlist: $$id"; status=1 ;; \
		esac; \
	done; \
	exit $$status

loc:
	go install github.com/boyter/scc/v3@latest
	scc --exclude-dir vendor --exclude-dir bin .

all: clean compile