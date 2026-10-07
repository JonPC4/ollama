# Maintaining the Patched Ollama Build

This branch carries local Ollama modifications exposing:

- CPU FFN offload
- tensor split
- KV offload

The patched branch should be based on the latest stable Ollama release tag, not `upstream/main`.

## Upgrade to a New Stable Ollama Release

```bash
cd /home/steam/.build/ollama-ncpu-ffn

git fetch upstream --tags
```

Install the matching official Ollama runtime first. This refreshes the official `llama-server` and supporting runtime libraries:

```bash
curl -fsSL https://ollama.com/install.sh |
  OLLAMA_VERSION=0.<new-stable> sh
```

For example:

```bash
curl -fsSL https://ollama.com/install.sh |
  OLLAMA_VERSION=0.35.1 sh
```

Switch to the maintained patched branch:

```bash
git switch patched-stable
```

If the branch contains only the previous stable Ollama release plus the local patch commits, rebase it onto the new stable tag:

```bash
git rebase v0.<new-stable>
```

For example:

```bash
git rebase v0.35.1
```

If the branch ancestry has become contaminated with upstream development commits, do not attempt to rebase those commits backward onto a stable release.

Instead, create a clean branch from the stable tag and cherry-pick only the local patch commit(s):

```bash
git switch -c patched-stable-new v0.<new-stable>
git cherry-pick <patch-commit>
```

## Test and Build

Run the relevant Go tests:

```bash
go test ./api ./llm ./server
```

Build the patched Ollama binary:

```bash
go build -o ollama-ncpu-ffn .
```

Install it as the custom Ollama server binary:

```bash
sudo install -o root -g root -m 0755 \
  ./ollama-ncpu-ffn \
  /usr/local/bin/ollama-ncpu-ffn
```

Do **not** install this binary as:

```text
/usr/local/lib/ollama/llama-server
```

`llama-server` is Ollama's internal inference runner and is supplied by the official Ollama runtime installation.

## Push the Updated Patched Branch

Pushing is not required to build or run the patched version locally.

To keep the GitHub fork synchronized after a successful upgrade:

```bash
git push --force-with-lease origin patched-stable
```

`--force-with-lease` is required after a rebase because the patch commit hashes will have changed.

## Workflow Summary

Local build only:

```text
fetch tags
→ install matching official runtime
→ rebase/cherry-pick patch
→ test
→ build
→ install patched binary
```

Keep GitHub synchronized:

```text
fetch tags
→ install matching official runtime
→ rebase/cherry-pick patch
→ test
→ build
→ install patched binary
→ push
```
