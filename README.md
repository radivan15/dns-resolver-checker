# DNS Resolver Checker

A Go CLI tool to query multiple DNS resolvers in parallel, compare results, and detect inconsistencies across DNS servers.

## What It Does

This tool reads a list of DNS resolvers (e.g. Google, Cloudflare, internal servers) and a list of domains from config files, then queries every combination concurrently using goroutines. Results are collected and printed to the terminal.

## Use Cases

- **DNS Migration** — Verify that new records have propagated to all resolvers before cutover
- **Split-horizon DNS Debugging** — Check whether internal and external resolvers return different answers
- **Consistency Audit** — Detect if any resolver returns a different answer than the rest

### Real-world Example

You just moved a server from `1.2.3.4` to `5.6.7.8`. This tool instantly shows which resolvers are still serving the old IP and which have updated — without manually querying each one.

> **Note on Consistency Audit:** This check is most reliable for domains with a static/single IP (e.g. internal services). For CDN-backed or anycast domains (like `google.com`, `cloudflare.com`), different resolvers returning different IPs is normal behavior (GeoDNS/anycast), not an inconsistency. The example output below shows this — don't treat it as a bug.

## Installation

```bash
git clone https://github.com/radivan15/dns-resolver-checker.git
cd dns-resolver-checker
go build -o dns-checker .
```

## Usage

### 1. Add DNS resolvers to `config/resolvers.txt`

```
8.8.8.8:53
8.8.4.4:53
1.1.1.1:53
1.0.0.1:53
```

### 2. Add domains to `config/domains.txt`

```
google.com
github.com
cloudflare.com
```

### 3. Run

```bash
go run main.go
```

### Example Output

```
domain=github.com           resolver=1.1.1.1:53      result=[20.205.243.166]
domain=google.com           resolver=1.1.1.1:53      result=[74.125.24.138 74.125.24.100 ...]
domain=google.com           resolver=8.8.8.8:53      result=[172.217.194.102 172.217.194.139 ...]
domain=cloudflare.com       resolver=8.8.8.8:53      result=[104.16.132.229 104.16.133.229 ...]
```

> Output order is non-deterministic because all queries run in parallel using goroutines.

## Project Structure

```
dns-resolver-checker/
├── main.go
├── go.mod
├── config/
│   ├── resolvers.txt    # list of DNS servers to query
│   └── domains.txt      # list of domains to check
├── .gitignore
├── LICENSE
└── README.md
```

## Roadmap

- [ ] Stage 1 — Query DNS against custom resolvers
- [ ] Stage 2 — Read resolvers and domains from config files
- [ ] Stage 3 — Concurrent queries with goroutines, channels, and `context` timeout/cancellation (so a dead resolver can't hang the whole run)
- [ ] Stage 4 — Detect inconsistencies across resolvers, exit non-zero when found (so it's usable in CI/cron monitoring)
- [ ] Stage 5 — Table and JSON output format

## Tech Stack

- Go standard library only (`net`, `context`, `sync`, `bufio` — no external dependencies)

## Apa yang saya pelajari

> TODO: isi setelah selesai mengerjakan tiap stage — jangan lupa update sebelum push final ke GitHub.
