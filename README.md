# genai-go

Go semantic conventions for GenAI, generated from
[open-telemetry/semantic-conventions-genai](https://github.com/open-telemetry/semantic-conventions-genai)
with [Weaver](https://github.com/open-telemetry/weaver).

This module exists because the upstream repository has no official Go
artifact yet (tracked in
[semantic-conventions-genai#247](https://github.com/open-telemetry/semantic-conventions-genai/issues/247)).
It will be retired once the dedicated GenAI conventions publish an official
Go package.

## Packages

| Package | Contents |
| --- | --- |
| `github.com/humbornjo/genai-go` | Attribute keys, value helpers, and enum values for all GenAI namespaces (`gen_ai`, `mcp`, `aws.bedrock`, `openai`) |
| `github.com/humbornjo/genai-go/genaiconv` | Typed metric instruments for the `gen_ai` namespace |
| `github.com/humbornjo/genai-go/mcpconv` | Typed metric instruments for the `mcp` namespace |

## Generated and maintained code

`make vendor` generates these files from the official schema; do not edit them
directly:

- `attribute_group.go`: attribute keys, typed enum values, and attribute helper
  functions.
- `genaiconv/metric.go` and `mcpconv/metric.go`: typed metric instruments and
  their recording helpers.

The compatibility changes are maintained in the generation inputs, not as
post-generation patches:

- `templates/registry/go/` contains the opentelemetry-go semconv templates with
  a small patch set for Weaver's v2 registry representation: `attr.key` instead
  of `attr.name`, `metric.name` instead of `metric.metric_name`, v2 filters, and
  template-typed attribute helpers.
- `internal/metricpool/pool.go` is maintained support code used by the generated
  metric methods. The official metric template imports OpenTelemetry's own
  `semconv/internal/metricpool`, which Go's `internal` import rules make
  inaccessible to this module, so the template imports this local equivalent.
- `doc.go` is maintained package documentation and declares `SchemaURL`.

Keeping the changes in templates is necessary because regeneration overwrites
the generated files. Editing those outputs directly would either lose the
changes on the next run or make manual CI generation non-reproducible.

## Example

```go
import (
	semconv "github.com/humbornjo/genai-go"
	"github.com/humbornjo/genai-go/genaiconv"
)

span.SetAttributes(
	semconv.GenAIOperationNameChat,
	semconv.GenAIProviderNameOpenAI,
	semconv.GenAIRequestModel("gpt-4o"),
)

dur, _ := genaiconv.NewClientOperationDuration(meter)
dur.Record(ctx, 1.24, genaiconv.OperationNameChat)
```

## Vendoring

```sh
make vendor               # generate from upstream main
make vendor REF=some-tag  # generate from another upstream ref
make test                 # test, vet, and check module tidiness
```

The vendor target clones the official schema, downloads the Weaver release
pinned by its `versions.env`, verifies the archive checksum, regenerates the
Go sources, and builds and vets them. The binary and upstream revision scratch
file stay ignored under `bin/` and `.upstream-sha`.

A GitHub Actions workflow runs the vendor target on demand. When the generated
output changes, it commits and pushes directly to `main`; the commit records
the exact upstream schema commit. Pushes and pull requests run the tests.

## Stability

All conventions upstream are development-stability. Expect breaking renames
between regenerations; pin the module version in consumers.
