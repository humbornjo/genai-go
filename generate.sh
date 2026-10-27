#!/usr/bin/env bash
# Regenerates the Go semantic convention code from
# https://github.com/open-telemetry/semantic-conventions-genai
#
# Usage: ./generate.sh [<genai-semconv-ref>]   # default: main
#
# Requirements: git, curl, tar, go, and gofmt. The Weaver release pinned by
# the upstream schema is downloaded and checksum-verified under ./bin.
set -euo pipefail

REF="${1:-main}"
ROOT="$(cd "$(dirname "$0")" && pwd)"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

if [ -n "${UPSTREAM_DIR:-}" ]; then
  echo ">> using pre-cloned upstream at $UPSTREAM_DIR"
  ln -s "$UPSTREAM_DIR" "$WORK/upstream"
else
  echo ">> cloning open-telemetry/semantic-conventions-genai @ $REF"
  git clone --depth 1 --branch "$REF" \
    https://github.com/open-telemetry/semantic-conventions-genai.git "$WORK/upstream"
fi

WEAVER_VERSION="$(sed -n 's/^WEAVER_VERSION=//p' "$WORK/upstream/versions.env")"
if [ -z "$WEAVER_VERSION" ]; then
  echo "upstream versions.env does not pin WEAVER_VERSION" >&2
  exit 1
fi

if [ ! -x "$ROOT/bin/weaver" ] || [ ! -f "$ROOT/bin/.weaver-version" ] || [ "$(<"$ROOT/bin/.weaver-version")" != "$WEAVER_VERSION" ]; then
  case "$(uname -s)/$(uname -m)" in
    Darwin/arm64|Darwin/aarch64) TARGET="aarch64-apple-darwin" ;;
    Darwin/x86_64) TARGET="x86_64-apple-darwin" ;;
    Linux/arm64|Linux/aarch64) TARGET="aarch64-unknown-linux-gnu" ;;
    Linux/x86_64) TARGET="x86_64-unknown-linux-gnu" ;;
    *) echo "unsupported Weaver platform: $(uname -s)/$(uname -m)" >&2; exit 1 ;;
  esac

  ARCHIVE="$WORK/weaver.tar.xz"
  URL="https://github.com/open-telemetry/weaver/releases/download/${WEAVER_VERSION}/weaver-${TARGET}.tar.xz"
  echo ">> downloading Weaver $WEAVER_VERSION for $TARGET"
  curl -fsSL -o "$ARCHIVE" "$URL"
  curl -fsSL -o "$ARCHIVE.sha256" "$URL.sha256"

  EXPECTED="$(awk '{print $1}' "$ARCHIVE.sha256")"
  if command -v sha256sum >/dev/null 2>&1; then
    ACTUAL="$(sha256sum "$ARCHIVE" | awk '{print $1}')"
  else
    ACTUAL="$(shasum -a 256 "$ARCHIVE" | awk '{print $1}')"
  fi
  if [ "$ACTUAL" != "$EXPECTED" ]; then
    echo "Weaver archive checksum mismatch" >&2
    exit 1
  fi

  mkdir -p "$WORK/weaver" "$ROOT/bin"
  tar -xJf "$ARCHIVE" -C "$WORK/weaver" --strip-components=1 "weaver-${TARGET}/weaver"
  install -m 0755 "$WORK/weaver/weaver" "$ROOT/bin/weaver"
  printf '%s\n' "$WEAVER_VERSION" > "$ROOT/bin/.weaver-version"
fi

echo ">> weaver registry generate"
"$ROOT/bin/weaver" registry generate \
  -r "$WORK/upstream/model" \
  --v2 \
  -t "$ROOT/templates" \
  --param tag="$REF" \
  go "$ROOT"

gofmt -w "$ROOT/attribute_group.go" "$ROOT/genaiconv" "$ROOT/mcpconv"

echo ">> go build ./..."
(cd "$ROOT" && go build ./... && go vet ./...)

git -C "$WORK/upstream" rev-parse HEAD > "$ROOT/.upstream-sha"
echo ">> done"
