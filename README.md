# genai-go

Go semantic conventions for GenAI, generated from [open-telemetry/semantic-conventions-genai](https://github.com/open-telemetry/semantic-conventions-genai) with [Weaver](https://github.com/open-telemetry/weaver).

This module exists because the upstream repository has no official Go artifact yet (tracked in [semantic-conventions-genai#247](https://github.com/open-telemetry/semantic-conventions-genai/issues/247)). It will be retired once the dedicated GenAI conventions publish an official Go package.

## Packages

| Package | Contents |
| --- | --- |
| `github.com/humbornjo/genai-go` | Attribute keys, value helpers, and enum values for all GenAI namespaces (`gen_ai`, `mcp`, `aws.bedrock`, `openai`) |
| `github.com/humbornjo/genai-go/genaiconv` | Typed metric instruments for the `gen_ai` namespace |
| `github.com/humbornjo/genai-go/mcpconv` | Typed metric instruments for the `mcp` namespace |

## Generated and maintained code

`make vendor` generates these files from the official schema; do not edit them directly:

- `attribute_group.go`: attribute keys, typed enum values, and attribute helper functions.
- `genaiconv/metric.go` and `mcpconv/metric.go`: typed metric instruments and their recording helpers.

The schema submodule and generator templates have different upstreams: `semantic-conventions-genai` provides schema YAML but no Go templates, while the templates forked here come from `opentelemetry-go/semconv/templates`. Their compatibility changes are maintained in the generation inputs, not as post-generation patches:

- `templates/registry/go/` contains the opentelemetry-go semconv templates with a small patch set for Weaver's v2 registry representation: `attr.key` instead of `attr.name`, `metric.name` instead of `metric.metric_name`, v2 filters, and template-typed attribute helpers.
- `internal/metricpool/pool.go` is maintained support code used by the generated metric methods. The official metric template imports OpenTelemetry's own `semconv/internal/metricpool`, which Go's `internal` import rules make inaccessible to this module, so the template imports this local equivalent.
- `doc.go` is maintained package documentation and declares `SchemaURL`.

Keeping the changes in templates is necessary because regeneration overwrites the generated files. Editing those outputs directly would either lose the changes on the next run or make manual CI generation non-reproducible.

Adding `opentelemetry-go` as another submodule would not eliminate the local templates: Weaver still needs the patched template directory, so generation would require an additional copy-and-patch step.

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

The official schema is a Git submodule at `upstream/`. Initialize it once after cloning:

```sh
git submodule update --init upstream
make vendor  # generate from the exact pinned schema commit
make update  # advance the pin to upstream main and regenerate
make test    # test, vet, and check module tidiness
```

`make vendor` uses the checked-out submodule without fetching or moving its pin. `make update` uses `git submodule update --remote` to advance the pin from the configured `main` branch before regenerating. Both targets download the Weaver release pinned by the submodule's `versions.env`, verify the archive checksum, regenerate the Go sources, and build and vet them. The Weaver binary stays ignored under `bin/`.

Normal module use and `make test` do not require the submodule; only regeneration does.

The manually triggered GitHub Actions workflow runs `make update`. When the schema pin or generated output changes, it commits and pushes both directly to `main`; the commit records the exact upstream schema commit. Pushes and pull requests run the tests.

## Stability

All conventions upstream are development-stability. Expect breaking renames between regenerations; pin the module version in consumers.
