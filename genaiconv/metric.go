// Code generated from semantic convention specification. DO NOT EDIT.

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package genaiconv provides types and functionality for OpenTelemetry semantic
// conventions in the "gen_ai" namespace.
package genaiconv

import (
	"context"

	"github.com/humbornjo/genai-go/internal/metricpool"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

// ErrorTypeAttr is an attribute conforming to the error.type semantic
// conventions. It represents the describes a class of error the operation ended
// with.
type ErrorTypeAttr string

var (
	// ErrorTypeOther is a fallback error value to be used when the instrumentation
	// doesn't define a custom value.
	ErrorTypeOther ErrorTypeAttr = "_OTHER"
)

// OperationNameAttr is an attribute conforming to the gen_ai.operation.name
// semantic conventions. It represents the name of the operation being performed.
type OperationNameAttr string

var (
	// OperationNameChat is the chat completion operation such as [OpenAI Chat API]
	// .
	//
	// [OpenAI Chat API]: https://platform.openai.com/docs/api-reference/chat
	OperationNameChat OperationNameAttr = "chat"
	// OperationNameGenerateContent is the multimodal content generation operation
	// such as [Gemini Generate Content].
	//
	// [Gemini Generate Content]: https://ai.google.dev/api/generate-content
	OperationNameGenerateContent OperationNameAttr = "generate_content"
	// OperationNameTextCompletion is the text completions operation such as
	// [OpenAI Completions API (Legacy)].
	//
	// [OpenAI Completions API (Legacy)]: https://platform.openai.com/docs/api-reference/completions
	OperationNameTextCompletion OperationNameAttr = "text_completion"
	// OperationNameEmbeddings is the embeddings operation such as
	// [OpenAI Create embeddings API].
	//
	// [OpenAI Create embeddings API]: https://platform.openai.com/docs/api-reference/embeddings/create
	OperationNameEmbeddings OperationNameAttr = "embeddings"
	// OperationNameRetrieval is the retrieval operation such as
	// [OpenAI Search Vector Store API].
	//
	// [OpenAI Search Vector Store API]: https://platform.openai.com/docs/api-reference/vector-stores/search
	OperationNameRetrieval OperationNameAttr = "retrieval"
	// OperationNameFetchResponse is the fetch a previously generated model response
	// by its identifier, without performing inference, such as
	// [OpenAI Get a model response].
	//
	// [OpenAI Get a model response]: https://platform.openai.com/docs/api-reference/responses/get
	OperationNameFetchResponse OperationNameAttr = "fetch_response"
	// OperationNameCreateAgent is the create GenAI agent.
	OperationNameCreateAgent OperationNameAttr = "create_agent"
	// OperationNameInvokeAgent is the invoke GenAI agent.
	OperationNameInvokeAgent OperationNameAttr = "invoke_agent"
	// OperationNameExecuteTool is the execute a tool.
	OperationNameExecuteTool OperationNameAttr = "execute_tool"
	// OperationNameInvokeWorkflow is the invoke GenAI workflow.
	OperationNameInvokeWorkflow OperationNameAttr = "invoke_workflow"
	// OperationNamePlan is the agent planning or task decomposition phase.
	OperationNamePlan OperationNameAttr = "plan"
	// OperationNameSearchMemory is the search/query memories from a memory store.
	OperationNameSearchMemory OperationNameAttr = "search_memory"
	// OperationNameCreateMemory is the create new memory records.
	OperationNameCreateMemory OperationNameAttr = "create_memory"
	// OperationNameUpdateMemory is the update existing memory records.
	OperationNameUpdateMemory OperationNameAttr = "update_memory"
	// OperationNameUpsertMemory is the create or update memory records without the
	// caller choosing which.
	OperationNameUpsertMemory OperationNameAttr = "upsert_memory"
	// OperationNameDeleteMemory is the delete memory records.
	OperationNameDeleteMemory OperationNameAttr = "delete_memory"
	// OperationNameCreateMemoryStore is the create or initialize a memory store.
	OperationNameCreateMemoryStore OperationNameAttr = "create_memory_store"
	// OperationNameDeleteMemoryStore is the delete or deprovision a memory store.
	OperationNameDeleteMemoryStore OperationNameAttr = "delete_memory_store"
)

// ProviderNameAttr is an attribute conforming to the gen_ai.provider.name
// semantic conventions. It represents the Generative AI provider as identified
// by the client or server instrumentation.
type ProviderNameAttr string

var (
	// ProviderNameOpenAI is the [OpenAI].
	//
	// [OpenAI]: https://openai.com/
	ProviderNameOpenAI ProviderNameAttr = "openai"
	// ProviderNameGCPGenAI is the any Google generative AI endpoint.
	ProviderNameGCPGenAI ProviderNameAttr = "gcp.gen_ai"
	// ProviderNameGCPVertexAI is the [Vertex AI].
	//
	// [Vertex AI]: https://cloud.google.com/vertex-ai
	ProviderNameGCPVertexAI ProviderNameAttr = "gcp.vertex_ai"
	// ProviderNameGCPGemini is the [Gemini].
	//
	// [Gemini]: https://cloud.google.com/products/gemini
	ProviderNameGCPGemini ProviderNameAttr = "gcp.gemini"
	// ProviderNameAnthropic is the [Anthropic].
	//
	// [Anthropic]: https://www.anthropic.com/
	ProviderNameAnthropic ProviderNameAttr = "anthropic"
	// ProviderNameCohere is the [Cohere].
	//
	// [Cohere]: https://cohere.com/
	ProviderNameCohere ProviderNameAttr = "cohere"
	// ProviderNameAzureAIInference is the azure AI Inference.
	ProviderNameAzureAIInference ProviderNameAttr = "azure.ai.inference"
	// ProviderNameAzureAIOpenAI is the [Azure OpenAI].
	//
	// [Azure OpenAI]: https://learn.microsoft.com/en-us/azure/ai-services/openai/overview
	ProviderNameAzureAIOpenAI ProviderNameAttr = "azure.ai.openai"
	// ProviderNameIBMWatsonxAI is the [IBM Watsonx AI].
	//
	// [IBM Watsonx AI]: https://www.ibm.com/products/watsonx-ai
	ProviderNameIBMWatsonxAI ProviderNameAttr = "ibm.watsonx.ai"
	// ProviderNameAWSBedrock is the [AWS Bedrock].
	//
	// [AWS Bedrock]: https://aws.amazon.com/bedrock
	ProviderNameAWSBedrock ProviderNameAttr = "aws.bedrock"
	// ProviderNamePerplexity is the [Perplexity].
	//
	// [Perplexity]: https://www.perplexity.ai/
	ProviderNamePerplexity ProviderNameAttr = "perplexity"
	// ProviderNameXAI is the [xAI].
	//
	// [xAI]: https://x.ai/
	ProviderNameXAI ProviderNameAttr = "x_ai"
	// ProviderNameDeepseek is the [DeepSeek].
	//
	// [DeepSeek]: https://www.deepseek.com/
	ProviderNameDeepseek ProviderNameAttr = "deepseek"
	// ProviderNameGroq is the [Groq].
	//
	// [Groq]: https://groq.com/
	ProviderNameGroq ProviderNameAttr = "groq"
	// ProviderNameMistral is the [Mistral].
	//
	// [Mistral]: https://mistral.ai/
	ProviderNameMistral ProviderNameAttr = "mistral"
	// ProviderNameMoonshotAI is the [Moonshot AI].
	//
	// [Moonshot AI]: https://www.moonshot.ai/
	ProviderNameMoonshotAI ProviderNameAttr = "moonshot_ai"
)

// TokenModalityAttr is an attribute conforming to the gen_ai.token.modality
// semantic conventions. It represents the modality of the tokens being counted.
type TokenModalityAttr string

var (
	// TokenModalityText is the text tokens.
	TokenModalityText TokenModalityAttr = "text"
	// TokenModalityImage is the image tokens.
	TokenModalityImage TokenModalityAttr = "image"
	// TokenModalityAudio is the audio tokens.
	TokenModalityAudio TokenModalityAttr = "audio"
	// TokenModalityUnknown is the modality is not known.
	TokenModalityUnknown TokenModalityAttr = "unknown"
)

// ClientEmbeddingsDuration is an instrument used to record metric values
// conforming to the "gen_ai.client.embeddings.duration" semantic conventions. It
// represents the genAI client embeddings operation duration.
type ClientEmbeddingsDuration struct {
	metric.Float64Histogram
}

var newClientEmbeddingsDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("GenAI client embeddings operation duration."),
	metric.WithUnit("s"),
}

// NewClientEmbeddingsDuration returns a new ClientEmbeddingsDuration instrument.
func NewClientEmbeddingsDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientEmbeddingsDuration, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientEmbeddingsDuration{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientEmbeddingsDurationOpts
	} else {
		opt = append(opt, newClientEmbeddingsDurationOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.client.embeddings.duration",
		opt...,
	)
	if err != nil {
		return ClientEmbeddingsDuration{noop.Float64Histogram{}}, err
	}
	return ClientEmbeddingsDuration{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientEmbeddingsDuration) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ClientEmbeddingsDuration) Name() string {
	return "gen_ai.client.embeddings.duration"
}

// Unit returns the semantic convention unit of the instrument
func (ClientEmbeddingsDuration) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ClientEmbeddingsDuration) Description() string {
	return "GenAI client embeddings operation duration."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
//
// Represents the end-to-end duration of an embeddings operation, measured from
// when the client issues the embeddings request to when the response is fully
// received or the operation ends with an error or cancellation.
// When this metric is reported alongside a `gen_ai.client.embeddings` span, the
// metric value SHOULD be the same as the span duration.
func (m ClientEmbeddingsDuration) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// Represents the end-to-end duration of an embeddings operation, measured from
// when the client issues the embeddings request to when the response is fully
// received or the operation ends with an error or cancellation.
// When this metric is reported alongside a `gen_ai.client.embeddings` span, the
// metric value SHOULD be the same as the span duration.
func (m ClientEmbeddingsDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (ClientEmbeddingsDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientEmbeddingsDuration) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientEmbeddingsDuration) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientEmbeddingsDuration) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientEmbeddingsDuration) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientEmbeddingsOperationInputTokens is an instrument used to record metric
// values conforming to the "gen_ai.client.embeddings.operation.input_tokens"
// semantic conventions. It represents the number of input tokens used per
// embeddings operation.
type ClientEmbeddingsOperationInputTokens struct {
	metric.Int64Histogram
}

var newClientEmbeddingsOperationInputTokensOpts = []metric.Int64HistogramOption{
	metric.WithDescription("The number of input tokens used per embeddings operation."),
	metric.WithUnit("{token}"),
}

// NewClientEmbeddingsOperationInputTokens returns a new
// ClientEmbeddingsOperationInputTokens instrument.
func NewClientEmbeddingsOperationInputTokens(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientEmbeddingsOperationInputTokens, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientEmbeddingsOperationInputTokens{noop.Int64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientEmbeddingsOperationInputTokensOpts
	} else {
		opt = append(opt, newClientEmbeddingsOperationInputTokensOpts...)
	}

	i, err := m.Int64Histogram(
		"gen_ai.client.embeddings.operation.input_tokens",
		opt...,
	)
	if err != nil {
		return ClientEmbeddingsOperationInputTokens{noop.Int64Histogram{}}, err
	}
	return ClientEmbeddingsOperationInputTokens{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientEmbeddingsOperationInputTokens) Inst() metric.Int64Histogram {
	return m.Int64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ClientEmbeddingsOperationInputTokens) Name() string {
	return "gen_ai.client.embeddings.operation.input_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientEmbeddingsOperationInputTokens) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientEmbeddingsOperationInputTokens) Description() string {
	return "The number of input tokens used per embeddings operation."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
//
// This metric SHOULD be reported when the count is readily available.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentations SHOULD record this metric for each embeddings operation,
// including failed ones. The recorded value SHOULD match the token count
// returned by the provider, or `0` if the response was not received or did not
// contain token usage.
func (m ClientEmbeddingsOperationInputTokens) Record(
	ctx context.Context,
	val int64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Int64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// This metric SHOULD be reported when the count is readily available.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentations SHOULD record this metric for each embeddings operation,
// including failed ones. The recorded value SHOULD match the token count
// returned by the provider, or `0` if the response was not received or did not
// contain token usage.
func (m ClientEmbeddingsOperationInputTokens) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (ClientEmbeddingsOperationInputTokens) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientEmbeddingsOperationInputTokens) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientEmbeddingsOperationInputTokens) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientEmbeddingsOperationInputTokens) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientEmbeddingsOperationInputTokens) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceDuration is an instrument used to record metric values
// conforming to the "gen_ai.client.inference.duration" semantic conventions. It
// represents the genAI client inference operation duration.
type ClientInferenceDuration struct {
	metric.Float64Histogram
}

var newClientInferenceDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("GenAI client inference operation duration."),
	metric.WithUnit("s"),
}

// NewClientInferenceDuration returns a new ClientInferenceDuration instrument.
func NewClientInferenceDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientInferenceDuration, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceDuration{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceDurationOpts
	} else {
		opt = append(opt, newClientInferenceDurationOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.client.inference.duration",
		opt...,
	)
	if err != nil {
		return ClientInferenceDuration{noop.Float64Histogram{}}, err
	}
	return ClientInferenceDuration{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceDuration) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceDuration) Name() string {
	return "gen_ai.client.inference.duration"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceDuration) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceDuration) Description() string {
	return "GenAI client inference operation duration."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
//
// Represents the end-to-end duration of a model inference operation, measured
// from when the client issues the inference request to when the response is
// fully received or the operation ends with an error or cancellation.
// When this metric is reported alongside a `gen_ai.client.inference` span, the
// metric value SHOULD be the same as the span duration.
func (m ClientInferenceDuration) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// Represents the end-to-end duration of a model inference operation, measured
// from when the client issues the inference request to when the response is
// fully received or the operation ends with an error or cancellation.
// When this metric is reported alongside a `gen_ai.client.inference` span, the
// metric value SHOULD be the same as the span duration.
func (m ClientInferenceDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (ClientInferenceDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceDuration) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceDuration) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceDuration) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceDuration) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceOperationInputTokens is an instrument used to record metric
// values conforming to the "gen_ai.client.inference.operation.input_tokens"
// semantic conventions. It represents the number of input (prompt) tokens used
// per inference operation.
type ClientInferenceOperationInputTokens struct {
	metric.Int64Histogram
}

var newClientInferenceOperationInputTokensOpts = []metric.Int64HistogramOption{
	metric.WithDescription("The number of input (prompt) tokens used per inference operation."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceOperationInputTokens returns a new
// ClientInferenceOperationInputTokens instrument.
func NewClientInferenceOperationInputTokens(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientInferenceOperationInputTokens, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceOperationInputTokens{noop.Int64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceOperationInputTokensOpts
	} else {
		opt = append(opt, newClientInferenceOperationInputTokensOpts...)
	}

	i, err := m.Int64Histogram(
		"gen_ai.client.inference.operation.input_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceOperationInputTokens{noop.Int64Histogram{}}, err
	}
	return ClientInferenceOperationInputTokens{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceOperationInputTokens) Inst() metric.Int64Histogram {
	return m.Int64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceOperationInputTokens) Name() string {
	return "gen_ai.client.inference.operation.input_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceOperationInputTokens) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceOperationInputTokens) Description() string {
	return "The number of input (prompt) tokens used per inference operation."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
//
// This metric SHOULD be reported when an operation involves the usage of tokens
// and the count is readily available.
// For example, if GenAI system returns usage information in the streaming
// response, it SHOULD be used. Or if GenAI system returns each token
// independently, instrumentation SHOULD count number of output tokens and record
// the result.
// If instrumentation cannot efficiently obtain number of input and/or output
// tokens, it MAY allow users to enable offline token counting. Otherwise it MUST
// NOT report usage metrics.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentations SHOULD record this metric for each inference operation,
// including failed ones. The recorded value SHOULD match the token count
// returned by the provider, or `0` if the response was not received or did not
// contain token usage.
func (m ClientInferenceOperationInputTokens) Record(
	ctx context.Context,
	val int64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Int64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// This metric SHOULD be reported when an operation involves the usage of tokens
// and the count is readily available.
// For example, if GenAI system returns usage information in the streaming
// response, it SHOULD be used. Or if GenAI system returns each token
// independently, instrumentation SHOULD count number of output tokens and record
// the result.
// If instrumentation cannot efficiently obtain number of input and/or output
// tokens, it MAY allow users to enable offline token counting. Otherwise it MUST
// NOT report usage metrics.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentations SHOULD record this metric for each inference operation,
// including failed ones. The recorded value SHOULD match the token count
// returned by the provider, or `0` if the response was not received or did not
// contain token usage.
func (m ClientInferenceOperationInputTokens) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (ClientInferenceOperationInputTokens) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceOperationInputTokens) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceOperationInputTokens) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceOperationInputTokens) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceOperationInputTokens) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceOperationOutputTokens is an instrument used to record metric
// values conforming to the "gen_ai.client.inference.operation.output_tokens"
// semantic conventions. It represents the number of output (completion) tokens
// used per inference operation.
type ClientInferenceOperationOutputTokens struct {
	metric.Int64Histogram
}

var newClientInferenceOperationOutputTokensOpts = []metric.Int64HistogramOption{
	metric.WithDescription("The number of output (completion) tokens used per inference operation."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceOperationOutputTokens returns a new
// ClientInferenceOperationOutputTokens instrument.
func NewClientInferenceOperationOutputTokens(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientInferenceOperationOutputTokens, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceOperationOutputTokens{noop.Int64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceOperationOutputTokensOpts
	} else {
		opt = append(opt, newClientInferenceOperationOutputTokensOpts...)
	}

	i, err := m.Int64Histogram(
		"gen_ai.client.inference.operation.output_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceOperationOutputTokens{noop.Int64Histogram{}}, err
	}
	return ClientInferenceOperationOutputTokens{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceOperationOutputTokens) Inst() metric.Int64Histogram {
	return m.Int64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceOperationOutputTokens) Name() string {
	return "gen_ai.client.inference.operation.output_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceOperationOutputTokens) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceOperationOutputTokens) Description() string {
	return "The number of output (completion) tokens used per inference operation."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
//
// This metric SHOULD be reported when an operation involves the usage of tokens
// and the count is readily available.
// For example, if GenAI system returns usage information in the streaming
// response, it SHOULD be used. Or if GenAI system returns each token
// independently, instrumentation SHOULD count number of output tokens and record
// the result.
// If instrumentation cannot efficiently obtain number of input and/or output
// tokens, it MAY allow users to enable offline token counting. Otherwise it MUST
// NOT report usage metrics.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentations SHOULD record this metric for each inference operation,
// including failed ones. The recorded value SHOULD match the token count
// returned by the provider, or `0` if the response was not received or did not
// contain token usage.
func (m ClientInferenceOperationOutputTokens) Record(
	ctx context.Context,
	val int64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Int64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// This metric SHOULD be reported when an operation involves the usage of tokens
// and the count is readily available.
// For example, if GenAI system returns usage information in the streaming
// response, it SHOULD be used. Or if GenAI system returns each token
// independently, instrumentation SHOULD count number of output tokens and record
// the result.
// If instrumentation cannot efficiently obtain number of input and/or output
// tokens, it MAY allow users to enable offline token counting. Otherwise it MUST
// NOT report usage metrics.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentations SHOULD record this metric for each inference operation,
// including failed ones. The recorded value SHOULD match the token count
// returned by the provider, or `0` if the response was not received or did not
// contain token usage.
func (m ClientInferenceOperationOutputTokens) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (ClientInferenceOperationOutputTokens) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceOperationOutputTokens) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceOperationOutputTokens) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceOperationOutputTokens) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceOperationOutputTokens) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceTimePerOutputChunk is an instrument used to record metric
// values conforming to the "gen_ai.client.inference.time_per_output_chunk"
// semantic conventions. It represents the time per output chunk, recorded for
// each chunk received after the first one, measured as the time elapsed from the
// end of the previous chunk to the end of the current chunk.
type ClientInferenceTimePerOutputChunk struct {
	metric.Float64Histogram
}

var newClientInferenceTimePerOutputChunkOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Time per output chunk, recorded for each chunk received after the first one, measured as the time elapsed from the end of the previous chunk to the end of the current chunk."),
	metric.WithUnit("s"),
}

// NewClientInferenceTimePerOutputChunk returns a new
// ClientInferenceTimePerOutputChunk instrument.
func NewClientInferenceTimePerOutputChunk(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientInferenceTimePerOutputChunk, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceTimePerOutputChunk{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceTimePerOutputChunkOpts
	} else {
		opt = append(opt, newClientInferenceTimePerOutputChunkOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.client.inference.time_per_output_chunk",
		opt...,
	)
	if err != nil {
		return ClientInferenceTimePerOutputChunk{noop.Float64Histogram{}}, err
	}
	return ClientInferenceTimePerOutputChunk{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceTimePerOutputChunk) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceTimePerOutputChunk) Name() string {
	return "gen_ai.client.inference.time_per_output_chunk"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceTimePerOutputChunk) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceTimePerOutputChunk) Description() string {
	return "Time per output chunk, recorded for each chunk received after the first one, measured as the time elapsed from the end of the previous chunk to the end of the current chunk."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
//
// This metric SHOULD be reported for streaming calls and SHOULD NOT be reported
// otherwise.
func (m ClientInferenceTimePerOutputChunk) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// This metric SHOULD be reported for streaming calls and SHOULD NOT be reported
// otherwise.
func (m ClientInferenceTimePerOutputChunk) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceTimePerOutputChunk) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceTimePerOutputChunk) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceTimePerOutputChunk) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceTimePerOutputChunk) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceTimeToFirstChunk is an instrument used to record metric values
// conforming to the "gen_ai.client.inference.time_to_first_chunk" semantic
// conventions. It represents the time to receive the first chunk, measured from
// when the client issues the generation request to when the first chunk is
// received in the response stream.
type ClientInferenceTimeToFirstChunk struct {
	metric.Float64Histogram
}

var newClientInferenceTimeToFirstChunkOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Time to receive the first chunk, measured from when the client issues the generation request to when the first chunk is received in the response stream."),
	metric.WithUnit("s"),
}

// NewClientInferenceTimeToFirstChunk returns a new
// ClientInferenceTimeToFirstChunk instrument.
func NewClientInferenceTimeToFirstChunk(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientInferenceTimeToFirstChunk, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceTimeToFirstChunk{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceTimeToFirstChunkOpts
	} else {
		opt = append(opt, newClientInferenceTimeToFirstChunkOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.client.inference.time_to_first_chunk",
		opt...,
	)
	if err != nil {
		return ClientInferenceTimeToFirstChunk{noop.Float64Histogram{}}, err
	}
	return ClientInferenceTimeToFirstChunk{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceTimeToFirstChunk) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceTimeToFirstChunk) Name() string {
	return "gen_ai.client.inference.time_to_first_chunk"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceTimeToFirstChunk) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceTimeToFirstChunk) Description() string {
	return "Time to receive the first chunk, measured from when the client issues the generation request to when the first chunk is received in the response stream."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
//
// This metric SHOULD be reported for streaming calls and SHOULD NOT be reported
// otherwise.
func (m ClientInferenceTimeToFirstChunk) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// This metric SHOULD be reported for streaming calls and SHOULD NOT be reported
// otherwise.
func (m ClientInferenceTimeToFirstChunk) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceTimeToFirstChunk) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceTimeToFirstChunk) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceTimeToFirstChunk) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceTimeToFirstChunk) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageCacheReadInputTokens is an instrument used to record
// metric values conforming to the
// "gen_ai.client.inference.usage.cache_read.input_tokens" semantic conventions.
// It represents the number of input tokens served from a provider-managed cache.
type ClientInferenceUsageCacheReadInputTokens struct {
	metric.Int64Counter
}

var newClientInferenceUsageCacheReadInputTokensOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of input tokens served from a provider-managed cache."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageCacheReadInputTokens returns a new
// ClientInferenceUsageCacheReadInputTokens instrument.
func NewClientInferenceUsageCacheReadInputTokens(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientInferenceUsageCacheReadInputTokens, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageCacheReadInputTokens{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageCacheReadInputTokensOpts
	} else {
		opt = append(opt, newClientInferenceUsageCacheReadInputTokensOpts...)
	}

	i, err := m.Int64Counter(
		"gen_ai.client.inference.usage.cache_read.input_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageCacheReadInputTokens{noop.Int64Counter{}}, err
	}
	return ClientInferenceUsageCacheReadInputTokens{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageCacheReadInputTokens) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageCacheReadInputTokens) Name() string {
	return "gen_ai.client.inference.usage.cache_read.input_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageCacheReadInputTokens) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageCacheReadInputTokens) Description() string {
	return "The number of input tokens served from a provider-managed cache."
}

// Add adds incr to the existing count for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// The tokenModality is the the modality of the tokens being counted.
//
// All additional attrs passed are included in the recorded value.
//
// This value is a subset of `gen_ai.client.inference.usage.input_tokens`.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageCacheReadInputTokens) Add(
	ctx context.Context,
	incr int64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	tokenModality TokenModalityAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
			attribute.String("gen_ai.token.modality", string(tokenModality)),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
				attribute.String("gen_ai.token.modality", string(tokenModality)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// This value is a subset of `gen_ai.client.inference.usage.input_tokens`.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageCacheReadInputTokens) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageCacheReadInputTokens) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageCacheReadInputTokens) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageCacheReadInputTokens) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageCacheReadInputTokens) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageCacheReadInputTokensObservable is an instrument used to
// record metric values conforming to the
// "gen_ai.client.inference.usage.cache_read.input_tokens" semantic conventions.
// It represents the number of input tokens served from a provider-managed cache.
type ClientInferenceUsageCacheReadInputTokensObservable struct {
	metric.Int64ObservableCounter
}

var newClientInferenceUsageCacheReadInputTokensObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of input tokens served from a provider-managed cache."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageCacheReadInputTokensObservable returns a new
// ClientInferenceUsageCacheReadInputTokensObservable instrument.
func NewClientInferenceUsageCacheReadInputTokensObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ClientInferenceUsageCacheReadInputTokensObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageCacheReadInputTokensObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageCacheReadInputTokensObservableOpts
	} else {
		opt = append(opt, newClientInferenceUsageCacheReadInputTokensObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"gen_ai.client.inference.usage.cache_read.input_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageCacheReadInputTokensObservable{noop.Int64ObservableCounter{}}, err
	}
	return ClientInferenceUsageCacheReadInputTokensObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageCacheReadInputTokensObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageCacheReadInputTokensObservable) Name() string {
	return "gen_ai.client.inference.usage.cache_read.input_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageCacheReadInputTokensObservable) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageCacheReadInputTokensObservable) Description() string {
	return "The number of input tokens served from a provider-managed cache."
}

// AttrOperationName returns a required attribute for the "gen_ai.operation.name"
// semantic convention. It represents the name of the operation being performed.
func (ClientInferenceUsageCacheReadInputTokensObservable) AttrOperationName(val OperationNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.operation.name", string(val))
}

// AttrProviderName returns a required attribute for the "gen_ai.provider.name"
// semantic convention. It represents the Generative AI provider as identified by
// the client or server instrumentation.
func (ClientInferenceUsageCacheReadInputTokensObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.provider.name", string(val))
}

// AttrTokenModality returns a required attribute for the "gen_ai.token.modality"
// semantic convention. It represents the modality of the tokens being counted.
func (ClientInferenceUsageCacheReadInputTokensObservable) AttrTokenModality(val TokenModalityAttr) attribute.KeyValue {
	return attribute.String("gen_ai.token.modality", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageCacheReadInputTokensObservable) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageCacheReadInputTokensObservable) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageCacheReadInputTokensObservable) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageCacheReadInputTokensObservable) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageCacheWriteInputTokens is an instrument used to record
// metric values conforming to the
// "gen_ai.client.inference.usage.cache_write.input_tokens" semantic conventions.
// It represents the number of input tokens written to a provider-managed cache.
type ClientInferenceUsageCacheWriteInputTokens struct {
	metric.Int64Counter
}

var newClientInferenceUsageCacheWriteInputTokensOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of input tokens written to a provider-managed cache."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageCacheWriteInputTokens returns a new
// ClientInferenceUsageCacheWriteInputTokens instrument.
func NewClientInferenceUsageCacheWriteInputTokens(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientInferenceUsageCacheWriteInputTokens, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageCacheWriteInputTokens{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageCacheWriteInputTokensOpts
	} else {
		opt = append(opt, newClientInferenceUsageCacheWriteInputTokensOpts...)
	}

	i, err := m.Int64Counter(
		"gen_ai.client.inference.usage.cache_write.input_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageCacheWriteInputTokens{noop.Int64Counter{}}, err
	}
	return ClientInferenceUsageCacheWriteInputTokens{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageCacheWriteInputTokens) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageCacheWriteInputTokens) Name() string {
	return "gen_ai.client.inference.usage.cache_write.input_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageCacheWriteInputTokens) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageCacheWriteInputTokens) Description() string {
	return "The number of input tokens written to a provider-managed cache."
}

// Add adds incr to the existing count for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// The tokenModality is the the modality of the tokens being counted.
//
// All additional attrs passed are included in the recorded value.
//
// This value is a subset of `gen_ai.client.inference.usage.input_tokens`.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageCacheWriteInputTokens) Add(
	ctx context.Context,
	incr int64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	tokenModality TokenModalityAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
			attribute.String("gen_ai.token.modality", string(tokenModality)),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
				attribute.String("gen_ai.token.modality", string(tokenModality)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// This value is a subset of `gen_ai.client.inference.usage.input_tokens`.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageCacheWriteInputTokens) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageCacheWriteInputTokens) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageCacheWriteInputTokens) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageCacheWriteInputTokens) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageCacheWriteInputTokens) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageCacheWriteInputTokensObservable is an instrument used to
// record metric values conforming to the
// "gen_ai.client.inference.usage.cache_write.input_tokens" semantic conventions.
// It represents the number of input tokens written to a provider-managed cache.
type ClientInferenceUsageCacheWriteInputTokensObservable struct {
	metric.Int64ObservableCounter
}

var newClientInferenceUsageCacheWriteInputTokensObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of input tokens written to a provider-managed cache."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageCacheWriteInputTokensObservable returns a new
// ClientInferenceUsageCacheWriteInputTokensObservable instrument.
func NewClientInferenceUsageCacheWriteInputTokensObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ClientInferenceUsageCacheWriteInputTokensObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageCacheWriteInputTokensObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageCacheWriteInputTokensObservableOpts
	} else {
		opt = append(opt, newClientInferenceUsageCacheWriteInputTokensObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"gen_ai.client.inference.usage.cache_write.input_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageCacheWriteInputTokensObservable{noop.Int64ObservableCounter{}}, err
	}
	return ClientInferenceUsageCacheWriteInputTokensObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageCacheWriteInputTokensObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageCacheWriteInputTokensObservable) Name() string {
	return "gen_ai.client.inference.usage.cache_write.input_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageCacheWriteInputTokensObservable) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageCacheWriteInputTokensObservable) Description() string {
	return "The number of input tokens written to a provider-managed cache."
}

// AttrOperationName returns a required attribute for the "gen_ai.operation.name"
// semantic convention. It represents the name of the operation being performed.
func (ClientInferenceUsageCacheWriteInputTokensObservable) AttrOperationName(val OperationNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.operation.name", string(val))
}

// AttrProviderName returns a required attribute for the "gen_ai.provider.name"
// semantic convention. It represents the Generative AI provider as identified by
// the client or server instrumentation.
func (ClientInferenceUsageCacheWriteInputTokensObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.provider.name", string(val))
}

// AttrTokenModality returns a required attribute for the "gen_ai.token.modality"
// semantic convention. It represents the modality of the tokens being counted.
func (ClientInferenceUsageCacheWriteInputTokensObservable) AttrTokenModality(val TokenModalityAttr) attribute.KeyValue {
	return attribute.String("gen_ai.token.modality", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageCacheWriteInputTokensObservable) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageCacheWriteInputTokensObservable) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageCacheWriteInputTokensObservable) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageCacheWriteInputTokensObservable) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageInputTokens is an instrument used to record metric values
// conforming to the "gen_ai.client.inference.usage.input_tokens" semantic
// conventions. It represents the number of input (prompt) tokens used, including
// cached tokens.
type ClientInferenceUsageInputTokens struct {
	metric.Int64Counter
}

var newClientInferenceUsageInputTokensOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of input (prompt) tokens used, including cached tokens."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageInputTokens returns a new
// ClientInferenceUsageInputTokens instrument.
func NewClientInferenceUsageInputTokens(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientInferenceUsageInputTokens, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageInputTokens{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageInputTokensOpts
	} else {
		opt = append(opt, newClientInferenceUsageInputTokensOpts...)
	}

	i, err := m.Int64Counter(
		"gen_ai.client.inference.usage.input_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageInputTokens{noop.Int64Counter{}}, err
	}
	return ClientInferenceUsageInputTokens{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageInputTokens) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageInputTokens) Name() string {
	return "gen_ai.client.inference.usage.input_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageInputTokens) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageInputTokens) Description() string {
	return "The number of input (prompt) tokens used, including cached tokens."
}

// Add adds incr to the existing count for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// The tokenModality is the the modality of the tokens being counted.
//
// All additional attrs passed are included in the recorded value.
//
// This metric SHOULD be reported when an operation involves the usage of tokens
// and the count is readily available.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageInputTokens) Add(
	ctx context.Context,
	incr int64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	tokenModality TokenModalityAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
			attribute.String("gen_ai.token.modality", string(tokenModality)),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
				attribute.String("gen_ai.token.modality", string(tokenModality)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// This metric SHOULD be reported when an operation involves the usage of tokens
// and the count is readily available.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageInputTokens) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageInputTokens) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageInputTokens) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageInputTokens) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageInputTokens) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageInputTokensObservable is an instrument used to record
// metric values conforming to the "gen_ai.client.inference.usage.input_tokens"
// semantic conventions. It represents the number of input (prompt) tokens used,
// including cached tokens.
type ClientInferenceUsageInputTokensObservable struct {
	metric.Int64ObservableCounter
}

var newClientInferenceUsageInputTokensObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of input (prompt) tokens used, including cached tokens."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageInputTokensObservable returns a new
// ClientInferenceUsageInputTokensObservable instrument.
func NewClientInferenceUsageInputTokensObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ClientInferenceUsageInputTokensObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageInputTokensObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageInputTokensObservableOpts
	} else {
		opt = append(opt, newClientInferenceUsageInputTokensObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"gen_ai.client.inference.usage.input_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageInputTokensObservable{noop.Int64ObservableCounter{}}, err
	}
	return ClientInferenceUsageInputTokensObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageInputTokensObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageInputTokensObservable) Name() string {
	return "gen_ai.client.inference.usage.input_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageInputTokensObservable) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageInputTokensObservable) Description() string {
	return "The number of input (prompt) tokens used, including cached tokens."
}

// AttrOperationName returns a required attribute for the "gen_ai.operation.name"
// semantic convention. It represents the name of the operation being performed.
func (ClientInferenceUsageInputTokensObservable) AttrOperationName(val OperationNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.operation.name", string(val))
}

// AttrProviderName returns a required attribute for the "gen_ai.provider.name"
// semantic convention. It represents the Generative AI provider as identified by
// the client or server instrumentation.
func (ClientInferenceUsageInputTokensObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.provider.name", string(val))
}

// AttrTokenModality returns a required attribute for the "gen_ai.token.modality"
// semantic convention. It represents the modality of the tokens being counted.
func (ClientInferenceUsageInputTokensObservable) AttrTokenModality(val TokenModalityAttr) attribute.KeyValue {
	return attribute.String("gen_ai.token.modality", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageInputTokensObservable) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageInputTokensObservable) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageInputTokensObservable) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageInputTokensObservable) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageOutputTokens is an instrument used to record metric values
// conforming to the "gen_ai.client.inference.usage.output_tokens" semantic
// conventions. It represents the number of output (completion) tokens used,
// including reasoning tokens.
type ClientInferenceUsageOutputTokens struct {
	metric.Int64Counter
}

var newClientInferenceUsageOutputTokensOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of output (completion) tokens used, including reasoning tokens."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageOutputTokens returns a new
// ClientInferenceUsageOutputTokens instrument.
func NewClientInferenceUsageOutputTokens(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientInferenceUsageOutputTokens, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageOutputTokens{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageOutputTokensOpts
	} else {
		opt = append(opt, newClientInferenceUsageOutputTokensOpts...)
	}

	i, err := m.Int64Counter(
		"gen_ai.client.inference.usage.output_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageOutputTokens{noop.Int64Counter{}}, err
	}
	return ClientInferenceUsageOutputTokens{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageOutputTokens) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageOutputTokens) Name() string {
	return "gen_ai.client.inference.usage.output_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageOutputTokens) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageOutputTokens) Description() string {
	return "The number of output (completion) tokens used, including reasoning tokens."
}

// Add adds incr to the existing count for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// The tokenModality is the the modality of the tokens being counted.
//
// All additional attrs passed are included in the recorded value.
//
// This metric SHOULD be reported when an operation involves the usage of tokens
// and the count is readily available.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageOutputTokens) Add(
	ctx context.Context,
	incr int64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	tokenModality TokenModalityAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
			attribute.String("gen_ai.token.modality", string(tokenModality)),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
				attribute.String("gen_ai.token.modality", string(tokenModality)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// This metric SHOULD be reported when an operation involves the usage of tokens
// and the count is readily available.
// When systems report both used tokens and billable tokens, instrumentation MUST
// report billable tokens.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageOutputTokens) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageOutputTokens) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageOutputTokens) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageOutputTokens) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageOutputTokens) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageOutputTokensObservable is an instrument used to record
// metric values conforming to the "gen_ai.client.inference.usage.output_tokens"
// semantic conventions. It represents the number of output (completion) tokens
// used, including reasoning tokens.
type ClientInferenceUsageOutputTokensObservable struct {
	metric.Int64ObservableCounter
}

var newClientInferenceUsageOutputTokensObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of output (completion) tokens used, including reasoning tokens."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageOutputTokensObservable returns a new
// ClientInferenceUsageOutputTokensObservable instrument.
func NewClientInferenceUsageOutputTokensObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ClientInferenceUsageOutputTokensObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageOutputTokensObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageOutputTokensObservableOpts
	} else {
		opt = append(opt, newClientInferenceUsageOutputTokensObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"gen_ai.client.inference.usage.output_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageOutputTokensObservable{noop.Int64ObservableCounter{}}, err
	}
	return ClientInferenceUsageOutputTokensObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageOutputTokensObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageOutputTokensObservable) Name() string {
	return "gen_ai.client.inference.usage.output_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageOutputTokensObservable) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageOutputTokensObservable) Description() string {
	return "The number of output (completion) tokens used, including reasoning tokens."
}

// AttrOperationName returns a required attribute for the "gen_ai.operation.name"
// semantic convention. It represents the name of the operation being performed.
func (ClientInferenceUsageOutputTokensObservable) AttrOperationName(val OperationNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.operation.name", string(val))
}

// AttrProviderName returns a required attribute for the "gen_ai.provider.name"
// semantic convention. It represents the Generative AI provider as identified by
// the client or server instrumentation.
func (ClientInferenceUsageOutputTokensObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.provider.name", string(val))
}

// AttrTokenModality returns a required attribute for the "gen_ai.token.modality"
// semantic convention. It represents the modality of the tokens being counted.
func (ClientInferenceUsageOutputTokensObservable) AttrTokenModality(val TokenModalityAttr) attribute.KeyValue {
	return attribute.String("gen_ai.token.modality", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageOutputTokensObservable) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageOutputTokensObservable) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageOutputTokensObservable) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageOutputTokensObservable) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageReasoningOutputTokens is an instrument used to record
// metric values conforming to the
// "gen_ai.client.inference.usage.reasoning.output_tokens" semantic conventions.
// It represents the number of output tokens used for reasoning (e.g.
// chain-of-thought, extended thinking).
type ClientInferenceUsageReasoningOutputTokens struct {
	metric.Int64Counter
}

var newClientInferenceUsageReasoningOutputTokensOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of output tokens used for reasoning (e.g. chain-of-thought, extended thinking)."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageReasoningOutputTokens returns a new
// ClientInferenceUsageReasoningOutputTokens instrument.
func NewClientInferenceUsageReasoningOutputTokens(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientInferenceUsageReasoningOutputTokens, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageReasoningOutputTokens{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageReasoningOutputTokensOpts
	} else {
		opt = append(opt, newClientInferenceUsageReasoningOutputTokensOpts...)
	}

	i, err := m.Int64Counter(
		"gen_ai.client.inference.usage.reasoning.output_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageReasoningOutputTokens{noop.Int64Counter{}}, err
	}
	return ClientInferenceUsageReasoningOutputTokens{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageReasoningOutputTokens) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageReasoningOutputTokens) Name() string {
	return "gen_ai.client.inference.usage.reasoning.output_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageReasoningOutputTokens) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageReasoningOutputTokens) Description() string {
	return "The number of output tokens used for reasoning (e.g. chain-of-thought, extended thinking)."
}

// Add adds incr to the existing count for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// The tokenModality is the the modality of the tokens being counted.
//
// All additional attrs passed are included in the recorded value.
//
// This value is a subset of `gen_ai.client.inference.usage.output_tokens`.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageReasoningOutputTokens) Add(
	ctx context.Context,
	incr int64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	tokenModality TokenModalityAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
			attribute.String("gen_ai.token.modality", string(tokenModality)),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
				attribute.String("gen_ai.token.modality", string(tokenModality)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// This value is a subset of `gen_ai.client.inference.usage.output_tokens`.
// Instrumentation SHOULD record contributions from successful and failed
// inference operations. Each measurement MUST be a known, strictly positive
// modality contribution.
// When a contribution's modality cannot reliably be determined, instrumentation
// MUST set `gen_ai.token.modality` to `unknown`.
func (m ClientInferenceUsageReasoningOutputTokens) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageReasoningOutputTokens) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageReasoningOutputTokens) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageReasoningOutputTokens) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageReasoningOutputTokens) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientInferenceUsageReasoningOutputTokensObservable is an instrument used to
// record metric values conforming to the
// "gen_ai.client.inference.usage.reasoning.output_tokens" semantic conventions.
// It represents the number of output tokens used for reasoning (e.g.
// chain-of-thought, extended thinking).
type ClientInferenceUsageReasoningOutputTokensObservable struct {
	metric.Int64ObservableCounter
}

var newClientInferenceUsageReasoningOutputTokensObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of output tokens used for reasoning (e.g. chain-of-thought, extended thinking)."),
	metric.WithUnit("{token}"),
}

// NewClientInferenceUsageReasoningOutputTokensObservable returns a new
// ClientInferenceUsageReasoningOutputTokensObservable instrument.
func NewClientInferenceUsageReasoningOutputTokensObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ClientInferenceUsageReasoningOutputTokensObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientInferenceUsageReasoningOutputTokensObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientInferenceUsageReasoningOutputTokensObservableOpts
	} else {
		opt = append(opt, newClientInferenceUsageReasoningOutputTokensObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"gen_ai.client.inference.usage.reasoning.output_tokens",
		opt...,
	)
	if err != nil {
		return ClientInferenceUsageReasoningOutputTokensObservable{noop.Int64ObservableCounter{}}, err
	}
	return ClientInferenceUsageReasoningOutputTokensObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientInferenceUsageReasoningOutputTokensObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (ClientInferenceUsageReasoningOutputTokensObservable) Name() string {
	return "gen_ai.client.inference.usage.reasoning.output_tokens"
}

// Unit returns the semantic convention unit of the instrument
func (ClientInferenceUsageReasoningOutputTokensObservable) Unit() string {
	return "{token}"
}

// Description returns the semantic convention description of the instrument
func (ClientInferenceUsageReasoningOutputTokensObservable) Description() string {
	return "The number of output tokens used for reasoning (e.g. chain-of-thought, extended thinking)."
}

// AttrOperationName returns a required attribute for the "gen_ai.operation.name"
// semantic convention. It represents the name of the operation being performed.
func (ClientInferenceUsageReasoningOutputTokensObservable) AttrOperationName(val OperationNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.operation.name", string(val))
}

// AttrProviderName returns a required attribute for the "gen_ai.provider.name"
// semantic convention. It represents the Generative AI provider as identified by
// the client or server instrumentation.
func (ClientInferenceUsageReasoningOutputTokensObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.provider.name", string(val))
}

// AttrTokenModality returns a required attribute for the "gen_ai.token.modality"
// semantic convention. It represents the modality of the tokens being counted.
func (ClientInferenceUsageReasoningOutputTokensObservable) AttrTokenModality(val TokenModalityAttr) attribute.KeyValue {
	return attribute.String("gen_ai.token.modality", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientInferenceUsageReasoningOutputTokensObservable) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientInferenceUsageReasoningOutputTokensObservable) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientInferenceUsageReasoningOutputTokensObservable) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientInferenceUsageReasoningOutputTokensObservable) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ClientOperationDuration is an instrument used to record metric values
// conforming to the "gen_ai.client.operation.duration" semantic conventions. It
// represents the genAI operation duration.
type ClientOperationDuration struct {
	metric.Float64Histogram
}

var newClientOperationDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("GenAI operation duration."),
	metric.WithUnit("s"),
}

// NewClientOperationDuration returns a new ClientOperationDuration instrument.
func NewClientOperationDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientOperationDuration, error) {
	// Check if the meter is nil.
	if m == nil {
		return ClientOperationDuration{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newClientOperationDurationOpts
	} else {
		opt = append(opt, newClientOperationDurationOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.client.operation.duration",
		opt...,
	)
	if err != nil {
		return ClientOperationDuration{noop.Float64Histogram{}}, err
	}
	return ClientOperationDuration{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ClientOperationDuration) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ClientOperationDuration) Name() string {
	return "gen_ai.client.operation.duration"
}

// Unit returns the semantic convention unit of the instrument
func (ClientOperationDuration) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ClientOperationDuration) Description() string {
	return "GenAI operation duration."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// All additional attrs passed are included in the recorded value.
//
// This metric SHOULD NOT be reported for inference or embeddings operations; use
// `gen_ai.client.inference.duration` or `gen_ai.client.embeddings.duration`
// instead.
func (m ClientOperationDuration) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// This metric SHOULD NOT be reported for inference or embeddings operations; use
// `gen_ai.client.inference.duration` or `gen_ai.client.embeddings.duration`
// instead.
func (m ClientOperationDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (ClientOperationDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrProviderName returns an optional attribute for the "gen_ai.provider.name"
// semantic convention. It represents the Generative AI provider as identified by
// the client or server instrumentation.
func (ClientOperationDuration) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	return attribute.String("gen_ai.provider.name", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ClientOperationDuration) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ClientOperationDuration) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ClientOperationDuration) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ClientOperationDuration) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ExecuteToolDuration is an instrument used to record metric values conforming
// to the "gen_ai.execute_tool.duration" semantic conventions. It represents the
// duration of a single tool execution.
type ExecuteToolDuration struct {
	metric.Float64Histogram
}

var newExecuteToolDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The duration of a single tool execution."),
	metric.WithUnit("s"),
}

// NewExecuteToolDuration returns a new ExecuteToolDuration instrument.
func NewExecuteToolDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ExecuteToolDuration, error) {
	// Check if the meter is nil.
	if m == nil {
		return ExecuteToolDuration{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newExecuteToolDurationOpts
	} else {
		opt = append(opt, newExecuteToolDurationOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.execute_tool.duration",
		opt...,
	)
	if err != nil {
		return ExecuteToolDuration{noop.Float64Histogram{}}, err
	}
	return ExecuteToolDuration{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ExecuteToolDuration) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ExecuteToolDuration) Name() string {
	return "gen_ai.execute_tool.duration"
}

// Unit returns the semantic convention unit of the instrument
func (ExecuteToolDuration) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ExecuteToolDuration) Description() string {
	return "The duration of a single tool execution."
}

// Record records val to the current distribution for attrs.
//
// The toolName is the name of the tool utilized by the agent.
//
// All additional attrs passed are included in the recorded value.
//
// Instrumentation that can reliably bound a single tool call SHOULD
// record this metric for every tool execution they can observe.
//
// When this metric is reported alongside a `gen_ai.execute_tool` span,
// the metric value SHOULD be the same as the span duration.
func (m ExecuteToolDuration) Record(
	ctx context.Context,
	val float64,
	toolName string,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.tool.name", toolName),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.tool.name", toolName),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// Instrumentation that can reliably bound a single tool call SHOULD
// record this metric for every tool execution they can observe.
//
// When this metric is reported alongside a `gen_ai.execute_tool` span,
// the metric value SHOULD be the same as the span duration.
func (m ExecuteToolDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (ExecuteToolDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrAgentName returns an optional attribute for the "gen_ai.agent.name"
// semantic convention. It represents the human-readable name of the agent
// executing the tool.
func (ExecuteToolDuration) AttrAgentName(val string) attribute.KeyValue {
	return attribute.String("gen_ai.agent.name", val)
}

// AttrToolType returns an optional attribute for the "gen_ai.tool.type" semantic
// convention. It represents the type of the tool utilized by the agent.
func (ExecuteToolDuration) AttrToolType(val string) attribute.KeyValue {
	return attribute.String("gen_ai.tool.type", val)
}

// InvokeAgentDuration is an instrument used to record metric values conforming
// to the "gen_ai.invoke_agent.duration" semantic conventions. It represents the
// end-to-end duration of a single in-process agent invocation, from the moment
// the invocation starts until the agent emits the last chunk of its final
// response or terminates with an error.
type InvokeAgentDuration struct {
	metric.Float64Histogram
}

var newInvokeAgentDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The end-to-end duration of a single in-process agent invocation, from the moment the invocation starts until the agent emits the last chunk of its final response or terminates with an error."),
	metric.WithUnit("s"),
}

// NewInvokeAgentDuration returns a new InvokeAgentDuration instrument.
func NewInvokeAgentDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (InvokeAgentDuration, error) {
	// Check if the meter is nil.
	if m == nil {
		return InvokeAgentDuration{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newInvokeAgentDurationOpts
	} else {
		opt = append(opt, newInvokeAgentDurationOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.invoke_agent.duration",
		opt...,
	)
	if err != nil {
		return InvokeAgentDuration{noop.Float64Histogram{}}, err
	}
	return InvokeAgentDuration{i}, nil
}

// Inst returns the underlying metric instrument.
func (m InvokeAgentDuration) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (InvokeAgentDuration) Name() string {
	return "gen_ai.invoke_agent.duration"
}

// Unit returns the semantic convention unit of the instrument
func (InvokeAgentDuration) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (InvokeAgentDuration) Description() string {
	return "The end-to-end duration of a single in-process agent invocation, from the moment the invocation starts until the agent emits the last chunk of its final response or terminates with an error."
}

// Record records val to the current distribution for attrs.
//
// All additional attrs passed are included in the recorded value.
//
// Intended for instrumentations of agent frameworks (for example, ADK,
// LangChain agents, CrewAI agents) that can reliably bound a single
// agent invocation.
//
// If instrumentation can only measure a single provider-facing client
// operation (for example, one model API call),
// `gen_ai.client.operation.duration` SHOULD be used instead. If
// instrumentation can reliably bound a higher-level workflow that
// coordinates multiple agents, `gen_ai.invoke_workflow.duration` SHOULD be
// used for that workflow. Instrumentation MAY emit several of these
// metrics for the same request path when more than one boundary is
// available.
//
// When this metric is reported alongside a `gen_ai.invoke_agent.internal`
// span, the metric value SHOULD be the same as the span duration.
func (m InvokeAgentDuration) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			attrs...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// Intended for instrumentations of agent frameworks (for example, ADK,
// LangChain agents, CrewAI agents) that can reliably bound a single
// agent invocation.
//
// If instrumentation can only measure a single provider-facing client
// operation (for example, one model API call),
// `gen_ai.client.operation.duration` SHOULD be used instead. If
// instrumentation can reliably bound a higher-level workflow that
// coordinates multiple agents, `gen_ai.invoke_workflow.duration` SHOULD be
// used for that workflow. Instrumentation MAY emit several of these
// metrics for the same request path when more than one boundary is
// available.
//
// When this metric is reported alongside a `gen_ai.invoke_agent.internal`
// span, the metric value SHOULD be the same as the span duration.
func (m InvokeAgentDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (InvokeAgentDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrAgentName returns an optional attribute for the "gen_ai.agent.name"
// semantic convention. It represents the human-readable name of the invoked
// GenAI agent.
func (InvokeAgentDuration) AttrAgentName(val string) attribute.KeyValue {
	return attribute.String("gen_ai.agent.name", val)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model configured for
// the agent.
func (InvokeAgentDuration) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// InvokeAgentInferenceCalls is an instrument used to record metric values
// conforming to the "gen_ai.invoke_agent.inference_calls" semantic conventions.
// It represents the number of inference (model) calls a GenAI agent makes during
// a single invocation.
type InvokeAgentInferenceCalls struct {
	metric.Int64Histogram
}

var newInvokeAgentInferenceCallsOpts = []metric.Int64HistogramOption{
	metric.WithDescription("The number of inference (model) calls a GenAI agent makes during a single invocation."),
	metric.WithUnit("{inference_call}"),
}

// NewInvokeAgentInferenceCalls returns a new InvokeAgentInferenceCalls
// instrument.
func NewInvokeAgentInferenceCalls(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (InvokeAgentInferenceCalls, error) {
	// Check if the meter is nil.
	if m == nil {
		return InvokeAgentInferenceCalls{noop.Int64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newInvokeAgentInferenceCallsOpts
	} else {
		opt = append(opt, newInvokeAgentInferenceCallsOpts...)
	}

	i, err := m.Int64Histogram(
		"gen_ai.invoke_agent.inference_calls",
		opt...,
	)
	if err != nil {
		return InvokeAgentInferenceCalls{noop.Int64Histogram{}}, err
	}
	return InvokeAgentInferenceCalls{i}, nil
}

// Inst returns the underlying metric instrument.
func (m InvokeAgentInferenceCalls) Inst() metric.Int64Histogram {
	return m.Int64Histogram
}

// Name returns the semantic convention name of the instrument.
func (InvokeAgentInferenceCalls) Name() string {
	return "gen_ai.invoke_agent.inference_calls"
}

// Unit returns the semantic convention unit of the instrument
func (InvokeAgentInferenceCalls) Unit() string {
	return "{inference_call}"
}

// Description returns the semantic convention description of the instrument
func (InvokeAgentInferenceCalls) Description() string {
	return "The number of inference (model) calls a GenAI agent makes during a single invocation."
}

// Record records val to the current distribution for attrs.
//
// All additional attrs passed are included in the recorded value.
//
// The distribution is scoped to a single agent invocation and SHOULD include
// only the inference calls the agent itself issued including failed ones;
// calls made by sub-agents or transferred-to agents are recorded against
// those agents' own invocations so that each inference call is counted
// exactly once across the call tree.
//
// This metric SHOULD be emitted together with the
// `gen_ai.invoke_agent.internal` span for the same invocation.
func (m InvokeAgentInferenceCalls) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			attrs...,
		),
	)

	m.Int64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// The distribution is scoped to a single agent invocation and SHOULD include
// only the inference calls the agent itself issued including failed ones;
// calls made by sub-agents or transferred-to agents are recorded against
// those agents' own invocations so that each inference call is counted
// exactly once across the call tree.
//
// This metric SHOULD be emitted together with the
// `gen_ai.invoke_agent.internal` span for the same invocation.
func (m InvokeAgentInferenceCalls) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Histogram.Record(ctx, val, *o...)
}

// AttrAgentName returns an optional attribute for the "gen_ai.agent.name"
// semantic convention. It represents the human-readable name of the GenAI agent
// provided by the application.
func (InvokeAgentInferenceCalls) AttrAgentName(val string) attribute.KeyValue {
	return attribute.String("gen_ai.agent.name", val)
}

// InvokeAgentToolCalls is an instrument used to record metric values conforming
// to the "gen_ai.invoke_agent.tool_calls" semantic conventions. It represents
// the number of tool calls a GenAI agent makes during a single invocation.
type InvokeAgentToolCalls struct {
	metric.Int64Histogram
}

var newInvokeAgentToolCallsOpts = []metric.Int64HistogramOption{
	metric.WithDescription("The number of tool calls a GenAI agent makes during a single invocation."),
	metric.WithUnit("{tool_call}"),
}

// NewInvokeAgentToolCalls returns a new InvokeAgentToolCalls instrument.
func NewInvokeAgentToolCalls(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (InvokeAgentToolCalls, error) {
	// Check if the meter is nil.
	if m == nil {
		return InvokeAgentToolCalls{noop.Int64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newInvokeAgentToolCallsOpts
	} else {
		opt = append(opt, newInvokeAgentToolCallsOpts...)
	}

	i, err := m.Int64Histogram(
		"gen_ai.invoke_agent.tool_calls",
		opt...,
	)
	if err != nil {
		return InvokeAgentToolCalls{noop.Int64Histogram{}}, err
	}
	return InvokeAgentToolCalls{i}, nil
}

// Inst returns the underlying metric instrument.
func (m InvokeAgentToolCalls) Inst() metric.Int64Histogram {
	return m.Int64Histogram
}

// Name returns the semantic convention name of the instrument.
func (InvokeAgentToolCalls) Name() string {
	return "gen_ai.invoke_agent.tool_calls"
}

// Unit returns the semantic convention unit of the instrument
func (InvokeAgentToolCalls) Unit() string {
	return "{tool_call}"
}

// Description returns the semantic convention description of the instrument
func (InvokeAgentToolCalls) Description() string {
	return "The number of tool calls a GenAI agent makes during a single invocation."
}

// Record records val to the current distribution for attrs.
//
// All additional attrs passed are included in the recorded value.
//
// The distribution is scoped to a single agent invocation and SHOULD include
// only the tool calls the agent itself triggers including failed ones;
// calls made by sub-agents or transferred-to agents are recorded against
// those agents' own invocations so that each tool call is counted exactly
// once across the call tree.
//
// Only client-side tool calls (tools executed by the agent or framework)
// are counted. Tools executed server-side by the model provider (for
// example, provider built-in web search or code execution) are not
// counted here.
//
// This metric SHOULD be emitted together with the
// `gen_ai.invoke_agent.internal` span for the same invocation.
func (m InvokeAgentToolCalls) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			attrs...,
		),
	)

	m.Int64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// The distribution is scoped to a single agent invocation and SHOULD include
// only the tool calls the agent itself triggers including failed ones;
// calls made by sub-agents or transferred-to agents are recorded against
// those agents' own invocations so that each tool call is counted exactly
// once across the call tree.
//
// Only client-side tool calls (tools executed by the agent or framework)
// are counted. Tools executed server-side by the model provider (for
// example, provider built-in web search or code execution) are not
// counted here.
//
// This metric SHOULD be emitted together with the
// `gen_ai.invoke_agent.internal` span for the same invocation.
func (m InvokeAgentToolCalls) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	if !m.Int64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Histogram.Record(ctx, val, *o...)
}

// AttrAgentName returns an optional attribute for the "gen_ai.agent.name"
// semantic convention. It represents the human-readable name of the GenAI agent
// provided by the application.
func (InvokeAgentToolCalls) AttrAgentName(val string) attribute.KeyValue {
	return attribute.String("gen_ai.agent.name", val)
}

// InvokeWorkflowDuration is an instrument used to record metric values
// conforming to the "gen_ai.invoke_workflow.duration" semantic conventions. It
// represents the records duration of GenAI workflow.
type InvokeWorkflowDuration struct {
	metric.Float64Histogram
}

var newInvokeWorkflowDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Records duration of GenAI workflow."),
	metric.WithUnit("s"),
}

// NewInvokeWorkflowDuration returns a new InvokeWorkflowDuration instrument.
func NewInvokeWorkflowDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (InvokeWorkflowDuration, error) {
	// Check if the meter is nil.
	if m == nil {
		return InvokeWorkflowDuration{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newInvokeWorkflowDurationOpts
	} else {
		opt = append(opt, newInvokeWorkflowDurationOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.invoke_workflow.duration",
		opt...,
	)
	if err != nil {
		return InvokeWorkflowDuration{noop.Float64Histogram{}}, err
	}
	return InvokeWorkflowDuration{i}, nil
}

// Inst returns the underlying metric instrument.
func (m InvokeWorkflowDuration) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (InvokeWorkflowDuration) Name() string {
	return "gen_ai.invoke_workflow.duration"
}

// Unit returns the semantic convention unit of the instrument
func (InvokeWorkflowDuration) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (InvokeWorkflowDuration) Description() string {
	return "Records duration of GenAI workflow."
}

// Record records val to the current distribution for attrs.
//
// All additional attrs passed are included in the recorded value.
//
// Represents the end-to-end duration of a workflow execution, measured from the
// point where application code initiates the workflow to the point where the
// workflow completes, independent of workflow complexity.
// A workflow is a coordinated process composed of multiple agents or other
// operations involving generative AI.
// If instrumentation measures only a single provider-facing client operation
// (for example, one model API call), `gen_ai.client.operation.duration` SHOULD
// be used instead. Instrumentation MAY emit both metrics for the same request
// path when both boundaries are available.
// When this metric is reported alongside a `gen_ai.invoke_workflow.internal`
// span, the metric value SHOULD be the same as the span duration.
// Individual systems may include additional system-specific attributes. Refer to
// system-specific documentation, if available.
func (m InvokeWorkflowDuration) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			attrs...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// Represents the end-to-end duration of a workflow execution, measured from the
// point where application code initiates the workflow to the point where the
// workflow completes, independent of workflow complexity.
// A workflow is a coordinated process composed of multiple agents or other
// operations involving generative AI.
// If instrumentation measures only a single provider-facing client operation
// (for example, one model API call), `gen_ai.client.operation.duration` SHOULD
// be used instead. Instrumentation MAY emit both metrics for the same request
// path when both boundaries are available.
// When this metric is reported alongside a `gen_ai.invoke_workflow.internal`
// span, the metric value SHOULD be the same as the span duration.
// Individual systems may include additional system-specific attributes. Refer to
// system-specific documentation, if available.
func (m InvokeWorkflowDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (InvokeWorkflowDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrWorkflowName returns an optional attribute for the "gen_ai.workflow.name"
// semantic convention. It represents the human-readable name of the GenAI
// workflow provided by the application.
func (InvokeWorkflowDuration) AttrWorkflowName(val string) attribute.KeyValue {
	return attribute.String("gen_ai.workflow.name", val)
}

// ServerRequestDuration is an instrument used to record metric values conforming
// to the "gen_ai.server.request.duration" semantic conventions. It represents
// the generative AI server request duration such as time-to-last byte or last
// output token.
type ServerRequestDuration struct {
	metric.Float64Histogram
}

var newServerRequestDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Generative AI server request duration such as time-to-last byte or last output token."),
	metric.WithUnit("s"),
}

// NewServerRequestDuration returns a new ServerRequestDuration instrument.
func NewServerRequestDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerRequestDuration, error) {
	// Check if the meter is nil.
	if m == nil {
		return ServerRequestDuration{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newServerRequestDurationOpts
	} else {
		opt = append(opt, newServerRequestDurationOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.server.request.duration",
		opt...,
	)
	if err != nil {
		return ServerRequestDuration{noop.Float64Histogram{}}, err
	}
	return ServerRequestDuration{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ServerRequestDuration) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ServerRequestDuration) Name() string {
	return "gen_ai.server.request.duration"
}

// Unit returns the semantic convention unit of the instrument
func (ServerRequestDuration) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ServerRequestDuration) Description() string {
	return "Generative AI server request duration such as time-to-last byte or last output token."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
func (m ServerRequestDuration) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
func (m ServerRequestDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (ServerRequestDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ServerRequestDuration) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ServerRequestDuration) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ServerRequestDuration) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ServerRequestDuration) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ServerTimePerOutputToken is an instrument used to record metric values
// conforming to the "gen_ai.server.time_per_output_token" semantic conventions.
// It represents the time per output token generated after the first token for
// successful responses.
type ServerTimePerOutputToken struct {
	metric.Float64Histogram
}

var newServerTimePerOutputTokenOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Time per output token generated after the first token for successful responses."),
	metric.WithUnit("s"),
}

// NewServerTimePerOutputToken returns a new ServerTimePerOutputToken instrument.
func NewServerTimePerOutputToken(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerTimePerOutputToken, error) {
	// Check if the meter is nil.
	if m == nil {
		return ServerTimePerOutputToken{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newServerTimePerOutputTokenOpts
	} else {
		opt = append(opt, newServerTimePerOutputTokenOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.server.time_per_output_token",
		opt...,
	)
	if err != nil {
		return ServerTimePerOutputToken{noop.Float64Histogram{}}, err
	}
	return ServerTimePerOutputToken{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ServerTimePerOutputToken) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ServerTimePerOutputToken) Name() string {
	return "gen_ai.server.time_per_output_token"
}

// Unit returns the semantic convention unit of the instrument
func (ServerTimePerOutputToken) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ServerTimePerOutputToken) Description() string {
	return "Time per output token generated after the first token for successful responses."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
func (m ServerTimePerOutputToken) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
func (m ServerTimePerOutputToken) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ServerTimePerOutputToken) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ServerTimePerOutputToken) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ServerTimePerOutputToken) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ServerTimePerOutputToken) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// ServerTimeToFirstToken is an instrument used to record metric values
// conforming to the "gen_ai.server.time_to_first_token" semantic conventions. It
// represents the time to generate first token for successful responses.
type ServerTimeToFirstToken struct {
	metric.Float64Histogram
}

var newServerTimeToFirstTokenOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Time to generate first token for successful responses."),
	metric.WithUnit("s"),
}

// NewServerTimeToFirstToken returns a new ServerTimeToFirstToken instrument.
func NewServerTimeToFirstToken(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerTimeToFirstToken, error) {
	// Check if the meter is nil.
	if m == nil {
		return ServerTimeToFirstToken{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newServerTimeToFirstTokenOpts
	} else {
		opt = append(opt, newServerTimeToFirstTokenOpts...)
	}

	i, err := m.Float64Histogram(
		"gen_ai.server.time_to_first_token",
		opt...,
	)
	if err != nil {
		return ServerTimeToFirstToken{noop.Float64Histogram{}}, err
	}
	return ServerTimeToFirstToken{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ServerTimeToFirstToken) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ServerTimeToFirstToken) Name() string {
	return "gen_ai.server.time_to_first_token"
}

// Unit returns the semantic convention unit of the instrument
func (ServerTimeToFirstToken) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ServerTimeToFirstToken) Description() string {
	return "Time to generate first token for successful responses."
}

// Record records val to the current distribution for attrs.
//
// The operationName is the the name of the operation being performed.
//
// The providerName is the the Generative AI provider as identified by the client
// or server instrumentation.
//
// All additional attrs passed are included in the recorded value.
func (m ServerTimeToFirstToken) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	providerName ProviderNameAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("gen_ai.operation.name", string(operationName)),
			attribute.String("gen_ai.provider.name", string(providerName)),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("gen_ai.operation.name", string(operationName)),
				attribute.String("gen_ai.provider.name", string(providerName)),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
func (m ServerTimeToFirstToken) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrRequestModel returns an optional attribute for the "gen_ai.request.model"
// semantic convention. It represents the name of the GenAI model a request is
// being made to.
func (ServerTimeToFirstToken) AttrRequestModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.request.model", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the genAI server port.
func (ServerTimeToFirstToken) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// AttrResponseModel returns an optional attribute for the
// "gen_ai.response.model" semantic convention. It represents the name of the
// model that generated the response.
func (ServerTimeToFirstToken) AttrResponseModel(val string) attribute.KeyValue {
	return attribute.String("gen_ai.response.model", val)
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the genAI server address.
func (ServerTimeToFirstToken) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}
