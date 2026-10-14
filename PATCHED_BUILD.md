cat > PATCHED_BUILD.md <<'EOF'
# Patched Ollama v0.40.4

Upstream base: ollama/ollama v0.40.4

Original local patch: 80cb827f

## Custom runner options

| API option | Type | llama.cpp flag |
|---|---|---|
| n_cpu_ffn | integer or omitted | --n-cpu-ffn |
| tensor_split | string or omitted | --tensor-split |
| kv_offload | boolean or omitted | --kv-offload / --no-kv-offload |

The options are applied when llama-server starts.

Changing runner options can trigger a model reload.

## Build requirements

Go 1.26.0 or newer.

The official Ollama v0.40.4 native runtime must be
installed separately. The Go build does not replace
llama-server or its GPU backend libraries.

## Test

    go test ./api ./llm ./server

## Build

    go build -trimpath \
      -ldflags="-X=github.com/ollama/ollama/version.Version=0.40.4 -X=github.com/ollama/ollama/server.mode=release" \
      -o ollama-ncpu-ffn .

## Install

    sudo install -o root -g root -m 0755 \
      ./ollama-ncpu-ffn \
      /usr/local/bin/ollama-ncpu-ffn

Both systemd services use this executable.

Do not replace llama-server with the Go executable.

## Upgrade policy

Create a new patched branch from each upstream stable
release tag. Cherry-pick the local patch, resolve any
conflicts, run tests, then build.

Preserve previous release branches for rollback.
EOF
