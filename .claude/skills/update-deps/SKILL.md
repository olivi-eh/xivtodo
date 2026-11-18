---
name: update-deps
description: Update dependencies across every surface of the repo (backend, godestone, frontend, scripts, GitHub Actions, Dockerfile), regenerate lock files, and verify everything builds. Use when asked to bump, upgrade, or refresh dependencies.
---

# Updating dependencies

Work through each surface below, verify it, then open one PR covering all of them. Do it on a branch, not `main`.

## Surfaces

| Directory | Manifest / lock | Tooling |
|---|---|---|
| `backend/` | `go.mod`, `go.sum` | Go (version from `go.mod` / Dockerfile) |
| `godestone/` | `go.mod` only, **no `go.sum`** | Go |
| `frontend/` | `package.json`, `package-lock.json` | Node from `frontend/.nvmrc`, npm |
| `scripts/` | `pyproject.toml`, `uv.lock` | uv, Python from `scripts/.python-version` |
| `.github/workflows/` | action version tags | `gh` |
| `backend/Dockerfile` | `golang:<ver>-trixie`, `distroless/base-debian13` | — |

## Backend

```sh
cd backend
go list -m -u all | grep '\['        # preview
go get -u ./... && go mod tidy
go build ./... && go vet ./...
DISCORD_CLIENT_ID=x DISCORD_CLIENT_SECRET=x DISCORD_REDIRECT_URI=http://localhost/auth go test ./...
```

`handlers` calls `log.Fatalln` at init if the Discord env vars are unset, so tests need placeholder values.

## Godestone

Do **not** use a blanket `go get -u ./...` here: it upgrades `github.com/gobwas/glob` to v1.x, which breaks `gocolly/colly/v2` v2.3.0 (`undefined: glob.Glob`). Instead upgrade the entries in `go.mod` selectively:

```sh
cd godestone
go get github.com/karashiiro/bingode@latest github.com/xivapi/godestone/v2@latest \
  github.com/gocolly/colly/v2@latest github.com/PuerkitoBio/goquery@latest \
  github.com/bits-and-blooms/bitset@latest golang.org/x/net@latest golang.org/x/text@latest
go mod tidy && go build ./... && go vet ./...
rm -f go.sum   # deployed from source; go.sum is intentionally not committed
```

Known pins (as of 2026-11):
- `gobwas/glob` stays at v0.2.3 until colly supports v1.
- `nlnwa/whatwg-url` stays at v0.6.2: newer versions renamed the module to `github.com/nationallibraryofnorway/whatwg-url`, so it can only move when colly switches.

## Frontend

Node is installed via nvm and isn't on the default PATH. Prefix commands with:

```sh
export PATH="$HOME/.nvm/versions/node/$(ls ~/.nvm/versions/node | grep "^v$(cat frontend/.nvmrc)\." | sort -V | tail -1)/bin:$PATH"
```

Run it from the repo root, then check `node -v` matches `.nvmrc`.

```sh
cd frontend
npm outdated
npm install <deps>@latest && npm install -D <devDeps>@latest   # bumps package.json floors too (repo convention)
npm update
npm audit
npm run lint
VITE_BACKEND_API_URI=https://api.xivtodo.com npm run build
```

- Major-version bumps: check the changelog and flag them to the user rather than taking them silently.
- Lint has ~33 pre-existing warnings (prettier / max-len); only errors are a regression. Compare against `main` if the count changes.
- Never run `npm audit fix --force`; it proposes downgrading vite. As of 2026-11, remaining audit findings come from `picomatch` (ReDoS, build-time only, no upstream fix).
- Confirm `lockfileVersion` didn't change in the lock diff.

## Scripts

```sh
cd scripts
uv lock --upgrade
uv tree --depth 1 --outdated
uv sync
uv run python -m py_compile *.py
```

If `pyproject.toml` floors are behind, bump them and re-lock. Watch transitive majors (e.g. `soupsieve` behind bs4) and smoke-test anything the scripts actually use.

## GitHub Actions and Docker

```sh
for a in actions/checkout actions/setup-node actions/configure-pages actions/upload-pages-artifact actions/deploy-pages; do
  echo "$a $(gh release view -R $a --json tagName -q .tagName)"
done
```

Workflows pin major tags (`@v7`), so only update on a new major. For the Dockerfile, bump `golang:<ver>-trixie` when a new Go minor ships (and the `go` directive in `go.mod` only if intentionally adopting it). Runtime upgrades (Node LTS in `.nvmrc`/`engines`, Python in `.python-version`/`requires-python`) are a separate decision: flag them, don't do them silently.

## Wrap-up

- Branch name like `update-dependencies-YYYY-MM`; one commit per run is fine.
- PR body: a table per surface (what changed, how verified), plus pins, audit findings, and anything skipped.
- Update the "Known pins" and audit notes above if they've changed.
