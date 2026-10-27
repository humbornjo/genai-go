// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated from semantic convention specification. DO NOT EDIT.

package semconv

import "go.opentelemetry.io/otel/attribute"

// Namespace: a2a
const (
	// A2aMessageIDKey is the attribute Key conforming to the "a2a.message.id"
	// semantic conventions. It represents the unique identifier of the A2A message.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "msg-user-1234"
	// Note: See the A2A specification for the [Message] object.
	//
	// [Message]: https://a2a-protocol.org/latest/specification/#414-message
	A2aMessageIDKey = attribute.Key("a2a.message.id")

	// A2aMessageReferenceTaskIdsKey is the attribute Key conforming to the
	// "a2a.message.reference_task_ids" semantic conventions. It represents the task
	// IDs referenced by the A2A request message.
	//
	// Type: string[]
	// Stability: Development
	//
	// Examples: "task-abc-5678"
	// Note: See the A2A specification for the [Message] object.
	//
	// [Message]: https://a2a-protocol.org/latest/specification/#414-message
	A2aMessageReferenceTaskIdsKey = attribute.Key("a2a.message.reference_task_ids")

	// A2aMethodNameKey is the attribute Key conforming to the "a2a.method.name"
	// semantic conventions. It represents the name of the A2A request method.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples:
	// Note: This attribute records the logical A2A operation name defined by the
	// protocol. For example, `SendMessage` is the A2A operation name, while
	// HTTP+JSON maps it to `POST /message:send`. See the A2A specification for
	// [core operations] and [method mappings].
	//
	// [core operations]: https://a2a-protocol.org/latest/specification/#31-core-operations
	// [method mappings]: https://a2a-protocol.org/latest/specification/#53-method-mapping-reference
	A2aMethodNameKey = attribute.Key("a2a.method.name")

	// A2aProtocolVersionKey is the attribute Key conforming to the
	// "a2a.protocol.version" semantic conventions. It represents the version of the
	// A2A protocol used.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "1.0"
	// Note: See the A2A specification for [versioning].
	//
	// [versioning]: https://a2a-protocol.org/latest/specification/#36-versioning
	A2aProtocolVersionKey = attribute.Key("a2a.protocol.version")

	// A2aTaskIDKey is the attribute Key conforming to the "a2a.task.id" semantic
	// conventions. It represents the unique identifier of an A2A task.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "task-5f2a7c7d9f8e4d2a"
	// Note: See the A2A specification for the [Task] object and
	// [task identifier semantics].
	//
	// [Task]: https://a2a-protocol.org/latest/specification/#411-task
	// [task identifier semantics]: https://a2a-protocol.org/latest/specification/#342-task-identifier-semantics
	A2aTaskIDKey = attribute.Key("a2a.task.id")

	// A2aTaskStateKey is the attribute Key conforming to the "a2a.task.state"
	// semantic conventions. It represents the state of an A2A task.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples:
	// Note: See the A2A specification for [TaskState].
	//
	// [TaskState]: https://a2a-protocol.org/latest/specification/#413-taskstate
	A2aTaskStateKey = attribute.Key("a2a.task.state")

	// A2aTenantKey is the attribute Key conforming to the "a2a.tenant" semantic
	// conventions. It represents the opaque routing identifier of the target
	// [A2A tenant].
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "billing"
	//
	// [A2A tenant]: https://github.com/a2aproject/A2A/blob/main/docs/topics/multi-tenancy.md
	A2aTenantKey = attribute.Key("a2a.tenant")
)

// A2aMessageID returns an attribute KeyValue conforming to the "a2a.message.id"
// semantic conventions. It represents the unique identifier of the A2A message.
func A2aMessageID(val string) attribute.KeyValue {
	return A2aMessageIDKey.String(val)
}

// A2aMessageReferenceTaskIds returns an attribute KeyValue conforming to the
// "a2a.message.reference_task_ids" semantic conventions. It represents the task
// IDs referenced by the A2A request message.
func A2aMessageReferenceTaskIds(val ...string) attribute.KeyValue {
	return A2aMessageReferenceTaskIdsKey.StringSlice(val)
}

// A2aProtocolVersion returns an attribute KeyValue conforming to the
// "a2a.protocol.version" semantic conventions. It represents the version of the
// A2A protocol used.
func A2aProtocolVersion(val string) attribute.KeyValue {
	return A2aProtocolVersionKey.String(val)
}

// A2aTaskID returns an attribute KeyValue conforming to the "a2a.task.id"
// semantic conventions. It represents the unique identifier of an A2A task.
func A2aTaskID(val string) attribute.KeyValue {
	return A2aTaskIDKey.String(val)
}

// A2aTenant returns an attribute KeyValue conforming to the "a2a.tenant"
// semantic conventions. It represents the opaque routing identifier of the
// target [A2A tenant].
//
// [A2A tenant]: https://github.com/a2aproject/A2A/blob/main/docs/topics/multi-tenancy.md
func A2aTenant(val string) attribute.KeyValue {
	return A2aTenantKey.String(val)
}

// Enum values for a2a.method.name
var (
	// Request to send a message to an agent and receive a non-streaming response.
	//
	// Stability: development
	A2aMethodNameSendMessage = A2aMethodNameKey.String("SendMessage")
	// Request to send a message to an agent and receive a streaming response.
	//
	// Stability: development
	A2aMethodNameSendStreamingMessage = A2aMethodNameKey.String("SendStreamingMessage")
	// Request to get the current state of a task.
	//
	// Stability: development
	A2aMethodNameGetTask = A2aMethodNameKey.String("GetTask")
	// Request to cancel a task.
	//
	// Stability: development
	A2aMethodNameCancelTask = A2aMethodNameKey.String("CancelTask")
	// Request to subscribe to task updates.
	//
	// Stability: development
	A2aMethodNameSubscribeToTask = A2aMethodNameKey.String("SubscribeToTask")
	// Request to list tasks.
	//
	// Stability: development
	A2aMethodNameListTasks = A2aMethodNameKey.String("ListTasks")
	// Request to create or update a task push notification configuration.
	//
	// Stability: development
	A2aMethodNameCreateTaskPushNotificationConfig = A2aMethodNameKey.String("CreateTaskPushNotificationConfig")
	// Request to get a task push notification configuration.
	//
	// Stability: development
	A2aMethodNameGetTaskPushNotificationConfig = A2aMethodNameKey.String("GetTaskPushNotificationConfig")
	// Request to list task push notification configurations.
	//
	// Stability: development
	A2aMethodNameListTaskPushNotificationConfigs = A2aMethodNameKey.String("ListTaskPushNotificationConfigs")
	// Request to delete a task push notification configuration.
	//
	// Stability: development
	A2aMethodNameDeleteTaskPushNotificationConfig = A2aMethodNameKey.String("DeleteTaskPushNotificationConfig")
	// Request to get the extended agent card.
	//
	// Stability: development
	A2aMethodNameGetExtendedAgentCard = A2aMethodNameKey.String("GetExtendedAgentCard")
)

// Enum values for a2a.task.state
var (
	// The task has been submitted.
	//
	// Stability: development
	A2aTaskStateTaskStateSubmitted = A2aTaskStateKey.String("TASK_STATE_SUBMITTED")
	// The task is being processed.
	//
	// Stability: development
	A2aTaskStateTaskStateWorking = A2aTaskStateKey.String("TASK_STATE_WORKING")
	// The task requires additional input.
	//
	// Stability: development
	A2aTaskStateTaskStateInputRequired = A2aTaskStateKey.String("TASK_STATE_INPUT_REQUIRED")
	// The task completed successfully.
	//
	// Stability: development
	A2aTaskStateTaskStateCompleted = A2aTaskStateKey.String("TASK_STATE_COMPLETED")
	// The task was canceled.
	//
	// Stability: development
	A2aTaskStateTaskStateCanceled = A2aTaskStateKey.String("TASK_STATE_CANCELED")
	// The task failed.
	//
	// Stability: development
	A2aTaskStateTaskStateFailed = A2aTaskStateKey.String("TASK_STATE_FAILED")
	// The task was rejected.
	//
	// Stability: development
	A2aTaskStateTaskStateRejected = A2aTaskStateKey.String("TASK_STATE_REJECTED")
	// The task requires authentication.
	//
	// Stability: development
	A2aTaskStateTaskStateAuthRequired = A2aTaskStateKey.String("TASK_STATE_AUTH_REQUIRED")
	// The task state is unspecified.
	//
	// Stability: development
	A2aTaskStateTaskStateUnspecified = A2aTaskStateKey.String("TASK_STATE_UNSPECIFIED")
)

// Namespace: aws
const (
	// AWSBedrockGuardrailIDKey is the attribute Key conforming to the
	// "aws.bedrock.guardrail.id" semantic conventions. It represents the unique
	// identifier of the AWS Bedrock Guardrail. A [guardrail] helps safeguard and
	// prevent unwanted behavior from model responses or user messages.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "sgi5gkybzqak"
	//
	// [guardrail]: https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails.html
	AWSBedrockGuardrailIDKey = attribute.Key("aws.bedrock.guardrail.id")

	// AWSBedrockKnowledgeBaseIDKey is the attribute Key conforming to the
	// "aws.bedrock.knowledge_base.id" semantic conventions. It represents the
	// unique identifier of the AWS Bedrock Knowledge base. A [knowledge base] is a
	// bank of information that can be queried by models to generate more relevant
	// responses and augment prompts.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "XFWUPB9PAW"
	//
	// [knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base.html
	AWSBedrockKnowledgeBaseIDKey = attribute.Key("aws.bedrock.knowledge_base.id")
)

// AWSBedrockGuardrailID returns an attribute KeyValue conforming to the
// "aws.bedrock.guardrail.id" semantic conventions. It represents the unique
// identifier of the AWS Bedrock Guardrail. A [guardrail] helps safeguard and
// prevent unwanted behavior from model responses or user messages.
//
// [guardrail]: https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails.html
func AWSBedrockGuardrailID(val string) attribute.KeyValue {
	return AWSBedrockGuardrailIDKey.String(val)
}

// AWSBedrockKnowledgeBaseID returns an attribute KeyValue conforming to the
// "aws.bedrock.knowledge_base.id" semantic conventions. It represents the unique
// identifier of the AWS Bedrock Knowledge base. A [knowledge base] is a bank of
// information that can be queried by models to generate more relevant responses
// and augment prompts.
//
// [knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base.html
func AWSBedrockKnowledgeBaseID(val string) attribute.KeyValue {
	return AWSBedrockKnowledgeBaseIDKey.String(val)
}

// Namespace: gen_ai
const (
	// GenAIAgentDescriptionKey is the attribute Key conforming to the
	// "gen_ai.agent.description" semantic conventions. It represents the free-form
	// description of the GenAI agent provided by the application.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "Helps with math problems", "Generates fiction stories"
	GenAIAgentDescriptionKey = attribute.Key("gen_ai.agent.description")

	// GenAIAgentIDKey is the attribute Key conforming to the "gen_ai.agent.id"
	// semantic conventions. It represents the unique and stable identifier of the
	// GenAI hosted agent resource.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "asst_5j66UpCpwteGg4YSxUnt7lPY",
	// "arn:aws:bedrock:us-east-1:123:agent/42",
	// "urn:agent:projects-123:projects:123:locations:us-east1:aiplatform:reasoningEngines:456"
	// Note: For hosted agents, this SHOULD be the provider-assigned stable
	// identifier of the agent resource such as [AWS Bedrock agent ARN] or
	// [GCP Agent Registry identifier].
	// It's NOT RECOMMENDED to record in-memory agent instance ids on this attribute
	// due to their transient nature.
	//
	// [AWS Bedrock agent ARN]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_Agent.html
	// [GCP Agent Registry identifier]: https://docs.cloud.google.com/agent-registry/concepts#agent-identifier
	GenAIAgentIDKey = attribute.Key("gen_ai.agent.id")

	// GenAIAgentNameKey is the attribute Key conforming to the "gen_ai.agent.name"
	// semantic conventions. It represents the human-readable name of the GenAI
	// agent provided by the application.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "Math Tutor", "Fiction Writer"
	GenAIAgentNameKey = attribute.Key("gen_ai.agent.name")

	// GenAIAgentVersionKey is the attribute Key conforming to the
	// "gen_ai.agent.version" semantic conventions. It represents the version of the
	// GenAI agent.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "1.0.0", "2025-05-01"
	GenAIAgentVersionKey = attribute.Key("gen_ai.agent.version")

	// GenAIConversationCompactedKey is the attribute Key conforming to the
	// "gen_ai.conversation.compacted" semantic conventions. It represents the
	// indicates whether the effective conversation context used for this operation
	// is a compacted view of a prior conversation.
	//
	// Type: boolean
	// Stability: Development
	//
	// Examples: true
	// Note: This attribute is a positive indicator of context compaction.
	// Instrumentations
	// SHOULD set it to `true` only when they can reliably determine that context
	// compaction was applied. Instrumentations SHOULD NOT set it to `false`; they
	// SHOULD leave it unset otherwise.
	GenAIConversationCompactedKey = attribute.Key("gen_ai.conversation.compacted")

	// GenAIConversationIDKey is the attribute Key conforming to the
	// "gen_ai.conversation.id" semantic conventions. It represents the unique
	// identifier for a conversation (session, thread), used to store and correlate
	// messages within this conversation.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "conv_5j66UpCpwteGg4YSxUnt7lPY"
	// Note: Instrumentations SHOULD populate conversation id when they have an
	// identifier
	// for the conversation readily available for a given operation, for example:
	//
	//   - when the client framework being instrumented manages conversation history
	//     (see [LlamaIndex chat store],
	//     [LangChain `session_id`],
	//     and [Google ADK sessions])
	//   - when instrumenting GenAI client libraries that maintain a conversation on
	//     the backend
	//     (see [AWS Bedrock agent sessions],
	//     [OpenAI Assistant threads])
	//
	// When no identifier for the conversation is available, instrumentations SHOULD
	// NOT
	// populate conversation id. For example, a new UUID, a trace identifier, or a
	// hash
	// of request content SHOULD NOT be used as a fallback value.
	//
	// Application developers that manage conversation history MAY add conversation
	// id to GenAI and other
	// spans or logs using custom span or log record processors or hooks provided by
	// instrumentation
	// libraries.
	//
	// [LlamaIndex chat store]: https://docs.llamaindex.ai/en/stable/module_guides/storing/chat_stores/
	// [LangChain `session_id`]: https://reference.langchain.com/python/langchain-core/runnables/history/RunnableWithMessageHistory
	// [Google ADK sessions]: https://adk.dev/sessions/session
	// [AWS Bedrock agent sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-session-state.html
	// [OpenAI Assistant threads]: https://platform.openai.com/docs/api-reference/threads
	GenAIConversationIDKey = attribute.Key("gen_ai.conversation.id")

	// GenAIDataSourceIDKey is the attribute Key conforming to the
	// "gen_ai.data_source.id" semantic conventions. It represents the data source
	// identifier.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "H7STPQYOND"
	// Note: Data sources are used by AI agents and RAG applications to store
	// grounding data. A data source may be an external database, object store,
	// document collection, website, or any other storage system used by the GenAI
	// agent or application. The `gen_ai.data_source.id` SHOULD match the identifier
	// used by the GenAI system rather than a name specific to the external storage,
	// such as a database or object store. Semantic conventions referencing
	// `gen_ai.data_source.id` MAY also leverage additional attributes, such as
	// `db.*`, to further identify and describe the data source.
	GenAIDataSourceIDKey = attribute.Key("gen_ai.data_source.id")

	// GenAIEmbeddingsDimensionCountKey is the attribute Key conforming to the
	// "gen_ai.embeddings.dimension.count" semantic conventions. It represents the
	// number of dimensions the resulting output embeddings should have.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 512, 1024
	GenAIEmbeddingsDimensionCountKey = attribute.Key("gen_ai.embeddings.dimension.count")

	// GenAIEvaluationExplanationKey is the attribute Key conforming to the
	// "gen_ai.evaluation.explanation" semantic conventions. It represents a
	// free-form explanation for the assigned score provided by the evaluator.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "The response is factually accurate but lacks sufficient detail to
	// fully address the question."
	GenAIEvaluationExplanationKey = attribute.Key("gen_ai.evaluation.explanation")

	// GenAIEvaluationNameKey is the attribute Key conforming to the
	// "gen_ai.evaluation.name" semantic conventions. It represents the name of the
	// evaluation metric used for the GenAI response.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "Relevance", "IntentResolution"
	GenAIEvaluationNameKey = attribute.Key("gen_ai.evaluation.name")

	// GenAIEvaluationScoreLabelKey is the attribute Key conforming to the
	// "gen_ai.evaluation.score.label" semantic conventions. It represents the human
	// readable label for evaluation.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "relevant", "not_relevant", "correct", "incorrect", "pass", "fail"
	// Note: This attribute provides a human-readable interpretation of the
	// evaluation score produced by an evaluator. For example, a score value of 1
	// could mean "relevant" in one evaluation system and "not relevant" in another,
	// depending on the scoring range and evaluator. The label SHOULD have low
	// cardinality. Possible values depend on the evaluation metric and evaluator
	// used; implementations SHOULD document the possible values.
	GenAIEvaluationScoreLabelKey = attribute.Key("gen_ai.evaluation.score.label")

	// GenAIEvaluationScoreValueKey is the attribute Key conforming to the
	// "gen_ai.evaluation.score.value" semantic conventions. It represents the
	// evaluation score returned by the evaluator.
	//
	// Type: double
	// Stability: Development
	//
	// Examples: 4.0
	GenAIEvaluationScoreValueKey = attribute.Key("gen_ai.evaluation.score.value")

	// GenAIInputMessagesKey is the attribute Key conforming to the
	// "gen_ai.input.messages" semantic conventions. It represents the chat history
	// provided to the model as an input.
	//
	// Type: any
	// Stability: Development
	//
	// Examples: "[\n {\n "role": "user",\n "parts": [\n {\n "type": "text",\n
	// "content": "Weather in Paris?"\n }\n ]\n },\n {\n "role": "assistant",\n
	// "parts": [\n {\n "type": "tool_call",\n "id":
	// "call_VSPygqKTWdrhaFErNvMV18Yl",\n "name": "get_weather",\n "arguments": {\n
	// "location": "Paris"\n }\n }\n ]\n },\n {\n "role": "tool",\n "parts": [\n {\n
	// "type": "tool_call_response",\n "id": "call_VSPygqKTWdrhaFErNvMV18Yl",\n
	// "response": "rainy, 57°F"\n }\n ]\n }\n]\n"
	// Note: Messages MUST be provided in the order they were sent to the model.
	// Instrumentations MAY provide a way for users to filter or truncate
	// input messages.
	//
	// > [!Warning]
	// > This attribute is likely to contain sensitive information including
	// > user/PII data.
	//
	// See [Recording content on attributes]
	// section for more details.
	//
	// [Recording content on attributes]: /docs/gen-ai/gen-ai-spans.md#recording-content-on-attributes
	GenAIInputMessagesKey = attribute.Key("gen_ai.input.messages")

	// GenAIMainAgentDescriptionKey is the attribute Key conforming to the
	// "gen_ai.main_agent.description" semantic conventions. It represents the
	// free-form description of the top-level Generative AI agent in the process.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "Helps with math problems", "Generates fiction stories"
	// Note: In Agent-to-Agent (A2A) communication, this maps to the `description`
	// property declared in the Agent Card.
	GenAIMainAgentDescriptionKey = attribute.Key("gen_ai.main_agent.description")

	// GenAIMainAgentIDKey is the attribute Key conforming to the
	// "gen_ai.main_agent.id" semantic conventions. It represents the unique and
	// stable identifier of the top-level Generative AI agent in the process.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "asst_5j66UpCpwteGg4YSxUnt7lPY",
	// "arn:aws:bedrock:us-east-1:123:agent/42",
	// "urn:agent:projects-123:projects:123:locations:us-east1:aiplatform:reasoningEngines:456"
	// Note: For main agents, this SHOULD be the provider-assigned stable identifier
	// of the agent resource such as [AWS Bedrock agent ARN] or
	// [GCP Agent Registry identifier].
	// Instrumentation MUST NOT generate `gen_ai.main_agent.id` from an in-memory or
	// process-local identifier. If no stable identifier is available, the agent
	// entity is not emitted, and instrumentation continues to report
	// `gen_ai.agent.name` at the span level when relevant.
	//
	// [AWS Bedrock agent ARN]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_Agent.html
	// [GCP Agent Registry identifier]: https://docs.cloud.google.com/agent-registry/concepts#agent-identifier
	GenAIMainAgentIDKey = attribute.Key("gen_ai.main_agent.id")

	// GenAIMainAgentNameKey is the attribute Key conforming to the
	// "gen_ai.main_agent.name" semantic conventions. It represents the
	// human-readable name of the top-level Generative AI agent in the process.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "Math Tutor", "Fiction Writer"
	// Note: In Agent-to-Agent (A2A) communication, this maps to the `name` property
	// declared in the Agent Card.
	GenAIMainAgentNameKey = attribute.Key("gen_ai.main_agent.name")

	// GenAIMemoryQueryTextKey is the attribute Key conforming to the
	// "gen_ai.memory.query.text" semantic conventions. It represents the search
	// query used to retrieve memories.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "user dietary preferences", "past flight bookings"
	// Note: Instrumentations SHOULD NOT capture this attribute by default. Capture
	// SHOULD be gated
	// by an explicit user opt-in via [content capture configuration].
	//
	// > [!Warning]
	// > This attribute may contain sensitive information.
	//
	// [content capture configuration]: /docs/gen-ai/gen-ai-spans.md#content-capture-configuration
	GenAIMemoryQueryTextKey = attribute.Key("gen_ai.memory.query.text")

	// GenAIMemoryRecordCountKey is the attribute Key conforming to the
	// "gen_ai.memory.record.count" semantic conventions. It represents the number
	// of memory records relevant to the operation.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 3
	// Note: For `search_memory` operations, this is the number of memory records
	// returned by the operation. For `create_memory` operations, this is the number
	// of memory records the operation attempted to create. For `update_memory`
	// operations, this is the number of memory records the operation attempted to
	// modify. For `upsert_memory` operations, this is the number of memory records
	// the operation attempted to create or update. For `delete_memory` operations,
	// this is the number of memory records the operation attempted to delete.
	GenAIMemoryRecordCountKey = attribute.Key("gen_ai.memory.record.count")

	// GenAIMemoryRecordIDKey is the attribute Key conforming to the
	// "gen_ai.memory.record.id" semantic conventions. It represents the unique
	// identifier of the memory record.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "mem_5j66UpCpwteGg4YSxUnt7lPY"
	GenAIMemoryRecordIDKey = attribute.Key("gen_ai.memory.record.id")

	// GenAIMemoryRecordsKey is the attribute Key conforming to the
	// "gen_ai.memory.records" semantic conventions. It represents the memory
	// records stored or retrieved in a memory operation.
	//
	// Type: any
	// Stability: Development
	//
	// Examples: "[\n {\n "content": "User prefers dark mode",\n "id": "mem_123",\n
	// "score": 0.95\n },\n {\n "content": {\n "preference": "vegetarian meals",\n
	// "confidence": 0.9\n },\n "metadata": {\n "source": "profile"\n }\n }\n]\n"
	// Note: Instrumentations SHOULD NOT capture this attribute by default. Capture
	// SHOULD be gated
	// by an explicit user opt-in via [content capture configuration].
	//
	// > [!Warning]
	// > This attribute may contain sensitive information including user/PII data.
	//
	// [content capture configuration]: /docs/gen-ai/gen-ai-spans.md#content-capture-configuration
	GenAIMemoryRecordsKey = attribute.Key("gen_ai.memory.records")

	// GenAIMemoryStoreIDKey is the attribute Key conforming to the
	// "gen_ai.memory.store.id" semantic conventions. It represents the unique
	// identifier of the memory store.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "ms_abc123", "user-preferences-store"
	// Note: Semantic conventions for individual components SHOULD document what
	// `gen_ai.memory.store.id` maps to within the implementation.
	GenAIMemoryStoreIDKey = attribute.Key("gen_ai.memory.store.id")

	// GenAIOperationNameKey is the attribute Key conforming to the
	// "gen_ai.operation.name" semantic conventions. It represents the name of the
	// operation being performed.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples:
	// Note: If one of the predefined values applies, but specific system uses a
	// different name it's RECOMMENDED to document it in the semantic conventions
	// for specific GenAI system and use system-specific name in the
	// instrumentation. If a different name is not documented, instrumentation
	// libraries SHOULD use applicable predefined value.
	GenAIOperationNameKey = attribute.Key("gen_ai.operation.name")

	// GenAIOutputMessagesKey is the attribute Key conforming to the
	// "gen_ai.output.messages" semantic conventions. It represents the messages
	// returned by the model where each message represents a specific model response
	// (choice, candidate).
	//
	// Type: any
	// Stability: Development
	//
	// Examples: "[\n {\n "role": "assistant",\n "parts": [\n {\n "type": "text",\n
	// "content": "The weather in Paris is currently rainy with a temperature of
	// 57°F."\n }\n ]\n }\n]\n"
	// Note: Each message represents a single output choice/candidate generated by
	// the model. Each message corresponds to exactly one generation
	// (choice/candidate) and vice versa - one choice cannot be split across
	// multiple messages or one message cannot contain parts from multiple choices.
	//
	// Instrumentations MAY provide a way for users to filter or truncate
	// output messages. `gen_ai.response.finish_reasons` remains aligned with the
	// generations returned by the provider, not with a filtered or truncated
	// `gen_ai.output.messages` value.
	//
	// When an operation ends in an error, this attribute SHOULD only be set
	// when the operation produced output, including partial streamed output.
	// Instrumentation SHOULD NOT synthesize output content for a failed
	// operation.
	//
	// > [!Warning]
	// > This attribute is likely to contain sensitive information including
	// > user/PII data.
	//
	// See [Recording content on attributes]
	// section for more details.
	//
	// [Recording content on attributes]: /docs/gen-ai/gen-ai-spans.md#recording-content-on-attributes
	GenAIOutputMessagesKey = attribute.Key("gen_ai.output.messages")

	// GenAIOutputTypeKey is the attribute Key conforming to the
	// "gen_ai.output.type" semantic conventions. It represents the represents the
	// content type requested by the client.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples:
	// Note: This attribute SHOULD be used when the client requests output of a
	// specific type. The model may return zero or more outputs of this type.
	// This attribute specifies the output modality and not the actual output
	// format. For example, if an image is requested, the actual output could be a
	// URL pointing to an image file.
	// Additional output format details may be recorded in the future in the
	// `gen_ai.output.{type}.*` attributes.
	GenAIOutputTypeKey = attribute.Key("gen_ai.output.type")

	// GenAIPromptNameKey is the attribute Key conforming to the
	// "gen_ai.prompt.name" semantic conventions. It represents the name of the
	// prompt that uniquely identifies it.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "analyze-code"
	GenAIPromptNameKey = attribute.Key("gen_ai.prompt.name")

	// GenAIPromptVersionKey is the attribute Key conforming to the
	// "gen_ai.prompt.version" semantic conventions. It represents the version of
	// the prompt template used.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "1.0.0", "2025-05-01", "prod", "v2"
	// Note: The version string can follow any versioning scheme chosen by the
	// application (e.g., SemVer, date-based, or platform-specific tags).
	// When a prompt management system is in use, this SHOULD match the
	// version identifier used by that system.
	GenAIPromptVersionKey = attribute.Key("gen_ai.prompt.version")

	// GenAIProviderNameKey is the attribute Key conforming to the
	// "gen_ai.provider.name" semantic conventions. It represents the Generative AI
	// provider as identified by the client or server instrumentation.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples:
	// Note: Semantic conventions for individual GenAI operations SHOULD clarify
	// which
	// kinds of providers (e.g. inference, embeddings, retrieval, memory, hosted
	// agent providers) apply when it is not clear from context.
	//
	// The attribute SHOULD be set based on the instrumentation's best knowledge
	// and may differ from the actual upstream provider. For example, a client SDK
	// may be configured against a proxy or hosting platform that transparently
	// relays requests to a different provider.
	//
	// The `gen_ai.provider.name` attribute acts as a discriminator that
	// identifies the GenAI telemetry format flavor specific to that provider
	// within GenAI semantic conventions.
	// It SHOULD be set consistently with provider-specific attributes and signals.
	// For example, GenAI spans, metrics, and events related to AWS Bedrock
	// should have the `gen_ai.provider.name` set to `aws.bedrock` and include
	// applicable `aws.bedrock.*` attributes and are not expected to include
	// `openai.*` attributes.
	GenAIProviderNameKey = attribute.Key("gen_ai.provider.name")

	// GenAIRequestChoiceCountKey is the attribute Key conforming to the
	// "gen_ai.request.choice.count" semantic conventions. It represents the target
	// number of candidate completions to return.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 3
	GenAIRequestChoiceCountKey = attribute.Key("gen_ai.request.choice.count")

	// GenAIRequestEncodingFormatsKey is the attribute Key conforming to the
	// "gen_ai.request.encoding_formats" semantic conventions. It represents the
	// encoding formats requested in an embeddings operation, if specified.
	//
	// Type: string[]
	// Stability: Development
	//
	// Examples: "base64"], ["float", "binary"
	// Note: In some GenAI systems the encoding formats are called embedding types.
	// Also, some GenAI systems only accept a single format per request.
	GenAIRequestEncodingFormatsKey = attribute.Key("gen_ai.request.encoding_formats")

	// GenAIRequestFrequencyPenaltyKey is the attribute Key conforming to the
	// "gen_ai.request.frequency_penalty" semantic conventions. It represents the
	// frequency penalty setting for the GenAI request.
	//
	// Type: double
	// Stability: Development
	//
	// Examples: 0.1
	GenAIRequestFrequencyPenaltyKey = attribute.Key("gen_ai.request.frequency_penalty")

	// GenAIRequestMaxTokensKey is the attribute Key conforming to the
	// "gen_ai.request.max_tokens" semantic conventions. It represents the maximum
	// number of tokens the model generates for a request.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 100
	GenAIRequestMaxTokensKey = attribute.Key("gen_ai.request.max_tokens")

	// GenAIRequestModelKey is the attribute Key conforming to the
	// "gen_ai.request.model" semantic conventions. It represents the name of the
	// GenAI model a request is being made to.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: gpt-4
	GenAIRequestModelKey = attribute.Key("gen_ai.request.model")

	// GenAIRequestPresencePenaltyKey is the attribute Key conforming to the
	// "gen_ai.request.presence_penalty" semantic conventions. It represents the
	// presence penalty setting for the GenAI request.
	//
	// Type: double
	// Stability: Development
	//
	// Examples: 0.1
	GenAIRequestPresencePenaltyKey = attribute.Key("gen_ai.request.presence_penalty")

	// GenAIRequestPreviousResponseIDKey is the attribute Key conforming to the
	// "gen_ai.request.previous_response.id" semantic conventions. It represents the
	// unique identifier of a previous response or interaction used to provide
	// context for the current operation.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "resp_0123456789aBCdef", "interaction-123"
	// Note: Instrumentations SHOULD populate this attribute when the request
	// references a previous response or interaction identifier to continue a
	// conversation or pass prior context.
	// For example, `previous_response_id` in [OpenAI Responses API]
	// or `previous_interaction_id` in [Google GenAI Interactions API].
	//
	// [OpenAI Responses API]: https://developers.openai.com/api/docs/guides/conversation-state#passing-context-from-the-previous-response
	// [Google GenAI Interactions API]: https://ai.google.dev/gemini-api/docs/interactions-overview
	GenAIRequestPreviousResponseIDKey = attribute.Key("gen_ai.request.previous_response.id")

	// GenAIRequestReasoningLevelKey is the attribute Key conforming to the
	// "gen_ai.request.reasoning.level" semantic conventions. It represents the
	// reasoning or thinking effort level requested for a GenAI model.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "low", "medium", "high"
	// Note: The value SHOULD be the exact string value sent to the provider.
	// Semantic conventions for individual providers SHOULD document which input
	// parameter maps to this attribute.
	GenAIRequestReasoningLevelKey = attribute.Key("gen_ai.request.reasoning.level")

	// GenAIRequestSeedKey is the attribute Key conforming to the
	// "gen_ai.request.seed" semantic conventions. It represents the requests with
	// same seed value more likely to return same result.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 100
	GenAIRequestSeedKey = attribute.Key("gen_ai.request.seed")

	// GenAIRequestStopSequencesKey is the attribute Key conforming to the
	// "gen_ai.request.stop_sequences" semantic conventions. It represents the list
	// of sequences that the model will use to stop generating further tokens.
	//
	// Type: string[]
	// Stability: Development
	//
	// Examples: "forest", "lived"
	GenAIRequestStopSequencesKey = attribute.Key("gen_ai.request.stop_sequences")

	// GenAIRequestStreamKey is the attribute Key conforming to the
	// "gen_ai.request.stream" semantic conventions. It represents the indicates
	// whether the GenAI request was made in streaming mode.
	//
	// Type: boolean
	// Stability: Development
	//
	// Examples:
	GenAIRequestStreamKey = attribute.Key("gen_ai.request.stream")

	// GenAIRequestStreamCursorKey is the attribute Key conforming to the
	// "gen_ai.request.stream_cursor" semantic conventions. It represents the cursor
	// identifying the last streamed event already received, used to resume a
	// streamed response from that position.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "42", "event-abc123"
	// Note: Instrumentations SHOULD populate this attribute when a request resumes
	// a streamed response from a prior position, for example when fetching a stored
	// response with streaming enabled.
	// For example, `starting_after` in the [OpenAI Responses API]
	// or `last_event_id` in the [Google GenAI Interactions API].
	//
	// [OpenAI Responses API]: https://developers.openai.com/api/docs/guides/background#streaming-a-background-response
	// [Google GenAI Interactions API]: https://ai.google.dev/api/interactions-api
	GenAIRequestStreamCursorKey = attribute.Key("gen_ai.request.stream_cursor")

	// GenAIRequestTemperatureKey is the attribute Key conforming to the
	// "gen_ai.request.temperature" semantic conventions. It represents the
	// temperature setting for the GenAI request.
	//
	// Type: double
	// Stability: Development
	//
	// Examples: 0.0
	GenAIRequestTemperatureKey = attribute.Key("gen_ai.request.temperature")

	// GenAIRequestTopKKey is the attribute Key conforming to the
	// "gen_ai.request.top_k" semantic conventions. It represents the top-K sampling
	// setting for the GenAI request: restricts token generation at each step to the
	// K most likely next tokens.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 40
	// Note: This is a decoding/sampling parameter (e.g., Anthropic `top_k`, Cohere
	// `k`, Google `topK`), not an output-shaping parameter. In particular, OpenAI's
	// `top_logprobs` controls how many per-token log-probabilities are returned in
	// the response and does not change generation; it MUST NOT be reported as
	// `gen_ai.request.top_k`.
	GenAIRequestTopKKey = attribute.Key("gen_ai.request.top_k")

	// GenAIRequestTopPKey is the attribute Key conforming to the
	// "gen_ai.request.top_p" semantic conventions. It represents the top_p sampling
	// setting for the GenAI request.
	//
	// Type: double
	// Stability: Development
	//
	// Examples: 1.0
	GenAIRequestTopPKey = attribute.Key("gen_ai.request.top_p")

	// GenAIResponseFinishReasonsKey is the attribute Key conforming to the
	// "gen_ai.response.finish_reasons" semantic conventions. It represents the
	// array of reasons the model stopped generating tokens, corresponding to each
	// generation received.
	//
	// Type: string[]
	// Stability: Development
	//
	// Examples: "stop"], ["stop", "length"], ["stop", "length", "error"
	// Note: Values correspond to generations in the same order as the returned
	// choices/candidates.
	//
	// Each position SHOULD contain the finish reason for the corresponding
	// choice/candidate. If a finish reason was expected but not received, for
	// example because generation failed, was cancelled, or a stream ended before
	// the final event, instrumentations SHOULD report `error` for that position
	// instead of omitting it.
	//
	// `error` indicates that the generation ended abnormally, whether reported
	// by the provider or inferred by the instrumentation.
	GenAIResponseFinishReasonsKey = attribute.Key("gen_ai.response.finish_reasons")

	// GenAIResponseIDKey is the attribute Key conforming to the
	// "gen_ai.response.id" semantic conventions. It represents the unique
	// identifier for the completion.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "chatcmpl-123"
	GenAIResponseIDKey = attribute.Key("gen_ai.response.id")

	// GenAIResponseModelKey is the attribute Key conforming to the
	// "gen_ai.response.model" semantic conventions. It represents the name of the
	// model that generated the response.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "gpt-4-0613"
	GenAIResponseModelKey = attribute.Key("gen_ai.response.model")

	// GenAIResponseStatusKey is the attribute Key conforming to the
	// "gen_ai.response.status" semantic conventions. It represents the lifecycle
	// status of a generated response, as reported by the provider when the response
	// is fetched or polled.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples: "completed", "in_progress"
	// Note: This attribute captures the lifecycle state of a (possibly background
	// or long-running) generation, such as whether it is queued, still running, or
	// has reached a terminal state. It is distinct from
	// `gen_ai.response.finish_reasons`, which describes why the model stopped once
	// it began producing output.
	// The value SHOULD be the provider's response status mapped onto the closest
	// member of this enum or a provider-specific value when none of these apply.
	// Semantic conventions for individual GenAI providers SHOULD document how their
	// status values map to this attribute.
	GenAIResponseStatusKey = attribute.Key("gen_ai.response.status")

	// GenAIResponseTimeToFirstChunkKey is the attribute Key conforming to the
	// "gen_ai.response.time_to_first_chunk" semantic conventions. It represents the
	// time to first chunk in a streaming response, measured from request issuance,
	// in seconds. The value is measured from when the client issues the generation
	// request to when the first chunk is received in the response stream.
	//
	// Type: double
	// Stability: Development
	//
	// Examples: 0.5, 1.2
	GenAIResponseTimeToFirstChunkKey = attribute.Key("gen_ai.response.time_to_first_chunk")

	// GenAIRetrievalDocumentsKey is the attribute Key conforming to the
	// "gen_ai.retrieval.documents" semantic conventions. It represents the
	// documents retrieved.
	//
	// Type: any
	// Stability: Development
	//
	// Examples: "[\n {\n "id": "doc_123",\n "score": 0.95\n },\n {\n "id":
	// "doc_456",\n "score": 0.87\n },\n {\n "id": "doc_789",\n "score": 0.82\n
	// }\n]\n"
	// Note: Each document object SHOULD contain the following properties when
	// available:
	// `id` (string): A unique identifier for the document, `score` (double): The
	// relevance score of the document
	GenAIRetrievalDocumentsKey = attribute.Key("gen_ai.retrieval.documents")

	// GenAIRetrievalQueryTextKey is the attribute Key conforming to the
	// "gen_ai.retrieval.query.text" semantic conventions. It represents the query
	// text used for retrieval.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "What is the capital of France?", "weather in Paris"
	// Note: > [!Warning]
	//
	// > This attribute may contain sensitive information.
	GenAIRetrievalQueryTextKey = attribute.Key("gen_ai.retrieval.query.text")

	// GenAIRetrievalTopKKey is the attribute Key conforming to the
	// "gen_ai.retrieval.top_k" semantic conventions. It represents the maximum
	// number of documents the retriever was asked to return for the query (also
	// known as `k`, `limit`, or `max_num_results`).
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 5
	GenAIRetrievalTopKKey = attribute.Key("gen_ai.retrieval.top_k")

	// GenAISkillDescriptionKey is the attribute Key conforming to the
	// "gen_ai.skill.description" semantic conventions. It represents the
	// [description of the Agent Skill].
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "Review a changelist against the team's review policy."
	// Note: > [!WARNING]
	//
	// > This attribute may contain sensitive information.
	//
	// [description of the Agent Skill]: https://agentskills.io/specification#skill-md-format
	GenAISkillDescriptionKey = attribute.Key("gen_ai.skill.description")

	// GenAISkillNameKey is the attribute Key conforming to the "gen_ai.skill.name"
	// semantic conventions. It represents the name of the [Agent Skill].
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "code-review", "pdf-processing"
	//
	// [Agent Skill]: https://agentskills.io/specification#skill-md-format
	GenAISkillNameKey = attribute.Key("gen_ai.skill.name")

	// GenAISkillResourceNameKey is the attribute Key conforming to the
	// "gen_ai.skill.resource.name" semantic conventions. It represents the
	// skill-relative name of the resource (script, reference, or asset) being
	// accessed or executed.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "references/review_policy.md", "assets/report_template.html"
	GenAISkillResourceNameKey = attribute.Key("gen_ai.skill.resource.name")

	// GenAISkillSourceURIKey is the attribute Key conforming to the
	// "gen_ai.skill.source.uri" semantic conventions. It represents the source URI
	// for loading the skill.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "https://skills.example.com/code-review",
	// "gs://example-skills/code-review", "file:///opt/skills/code-review"
	// Note: This SHOULD identify the skill source, not a temporary location where
	// the
	// skill was materialized for execution.
	//
	// > [!WARNING]
	// > This attribute may contain sensitive information.
	//
	// Sensitive URI components, such as user information and access tokens in query
	// parameters, SHOULD be scrubbed when instrumentations can identify them.
	GenAISkillSourceURIKey = attribute.Key("gen_ai.skill.source.uri")

	// GenAISystemInstructionsKey is the attribute Key conforming to the
	// "gen_ai.system_instructions" semantic conventions. It represents the system
	// message or instructions provided to the GenAI model separately from the chat
	// history.
	//
	// Type: any
	// Stability: Development
	//
	// Examples: "[\n {\n "type": "text",\n "content": "You are an Agent that greet
	// users, always use greetings tool to respond"\n }\n]\n", "[\n {\n "type":
	// "text",\n "content": "You are a language translator."\n },\n {\n "type":
	// "text",\n "content": "Your mission is to translate text in English to
	// French."\n }\n]\n"
	// Note: This attribute SHOULD be used when the corresponding provider or API
	// allows to provide system instructions or messages separately from the
	// chat history.
	//
	// Instructions that are part of the chat history SHOULD be recorded in
	// `gen_ai.input.messages` attribute instead.
	//
	// Instrumentations MAY provide a way for users to filter or truncate
	// system instructions.
	//
	// > [!Warning]
	// > This attribute may contain sensitive information.
	//
	// See [Recording content on attributes]
	// section for more details.
	//
	// [Recording content on attributes]: /docs/gen-ai/gen-ai-spans.md#recording-content-on-attributes
	GenAISystemInstructionsKey = attribute.Key("gen_ai.system_instructions")

	// GenAITokenModalityKey is the attribute Key conforming to the
	// "gen_ai.token.modality" semantic conventions. It represents the modality of
	// the tokens being counted.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples: "text", "image", "audio"
	// Note: When the provider does not break usage down by modality and the
	// modality cannot be reliably determined, instrumentations SHOULD report the
	// tokens under the `unknown` modality.
	GenAITokenModalityKey = attribute.Key("gen_ai.token.modality")

	// GenAIToolCallArgumentsKey is the attribute Key conforming to the
	// "gen_ai.tool.call.arguments" semantic conventions. It represents the
	// parameters passed to the tool call.
	//
	// Type: any
	// Stability: Development
	//
	// Examples: "{\n "location": "San Francisco?",\n "date": "2025-10-01"\n}\n"
	// Note: > [!WARNING]
	//
	// > This attribute may contain sensitive information.
	//
	// It's expected to be an object - in case a serialized string is available
	// to the instrumentation, the instrumentation SHOULD do the best effort to
	// deserialize it to an object.
	GenAIToolCallArgumentsKey = attribute.Key("gen_ai.tool.call.arguments")

	// GenAIToolCallIDKey is the attribute Key conforming to the
	// "gen_ai.tool.call.id" semantic conventions. It represents the tool call
	// identifier supplied by the model or provider.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "call_mszuSIzqtI65i1wAUOE8w5H4"
	// Note: This attribute SHOULD be set to the identifier received from the model
	// or provider
	// if and only if one was supplied.
	GenAIToolCallIDKey = attribute.Key("gen_ai.tool.call.id")

	// GenAIToolCallResultKey is the attribute Key conforming to the
	// "gen_ai.tool.call.result" semantic conventions. It represents the result
	// returned by the tool call (if any and if execution was successful).
	//
	// Type: any
	// Stability: Development
	//
	// Examples: "{\n "temperature_range": {\n "high": 75,\n "low": 60\n },\n
	// "conditions": "sunny"\n}\n"
	// Note: > [!WARNING]
	//
	// > This attribute may contain sensitive information.
	//
	// It's expected to be an object - in case a serialized string is available
	// to the instrumentation, the instrumentation SHOULD do the best effort to
	// deserialize it to an object.
	GenAIToolCallResultKey = attribute.Key("gen_ai.tool.call.result")

	// GenAIToolDefinitionsKey is the attribute Key conforming to the
	// "gen_ai.tool.definitions" semantic conventions. It represents the list of
	// tool definitions available to the GenAI agent or model.
	//
	// Type: any
	// Stability: Development
	//
	// Examples: "[\n {\n "type": "function",\n "name": "get_current_weather",\n
	// "description": "Get the current weather in a given location",\n "parameters":
	// {\n "type": "object",\n "properties": {\n "location": {\n "type": "string",\n
	// "description": "The city and state, e.g. San Francisco, CA"\n },\n "unit":
	// {\n "type": "string",\n "enum": [\n "celsius",\n "fahrenheit"\n ]\n }\n },\n
	// "required": [\n "location",\n "unit"\n ]\n }\n }\n]\n"
	// Note: > [!WARNING]
	//
	// > This attribute may contain sensitive information.
	//
	// Since this attribute could be large, it's NOT RECOMMENDED to populate
	// non-required properties by default. Instrumentations MAY provide a way
	// to enable populating optional properties.
	GenAIToolDefinitionsKey = attribute.Key("gen_ai.tool.definitions")

	// GenAIToolDescriptionKey is the attribute Key conforming to the
	// "gen_ai.tool.description" semantic conventions. It represents the tool
	// description.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "Multiply two numbers"
	// Note: > [!WARNING]
	//
	// > This attribute may contain sensitive information.
	GenAIToolDescriptionKey = attribute.Key("gen_ai.tool.description")

	// GenAIToolNameKey is the attribute Key conforming to the "gen_ai.tool.name"
	// semantic conventions. It represents the name of the tool utilized by the
	// agent.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "Flights"
	GenAIToolNameKey = attribute.Key("gen_ai.tool.name")

	// GenAIToolTypeKey is the attribute Key conforming to the "gen_ai.tool.type"
	// semantic conventions. It represents the type of the tool utilized by the
	// agent.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "function", "extension", "datastore"
	// Note: Extension: A tool executed on the agent-side to directly call external
	// APIs, bridging the gap between the agent and real-world systems.
	// Agent-side operations involve actions that are performed by the agent on the
	// server or within the agent's controlled environment.
	// Function: A tool executed on the client-side, where the agent generates
	// parameters for a predefined function, and the client executes the logic.
	// Client-side operations are actions taken on the user's end or within the
	// client application.
	// Datastore: A tool used by the agent to access and query structured or
	// unstructured external data for retrieval-augmented tasks or knowledge
	// updates.
	GenAIToolTypeKey = attribute.Key("gen_ai.tool.type")

	// GenAITransferModeKey is the attribute Key conforming to the
	// "gen_ai.transfer.mode" semantic conventions. It represents the describes how
	// control passes to the transfer target.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples:
	GenAITransferModeKey = attribute.Key("gen_ai.transfer.mode")

	// GenAITransferTargetNameKey is the attribute Key conforming to the
	// "gen_ai.transfer.target.name" semantic conventions. It represents the
	// human-readable name of the agent targeted by the transfer.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "weather_agent"
	GenAITransferTargetNameKey = attribute.Key("gen_ai.transfer.target.name")

	// GenAIUsageAudioCacheReadInputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.audio.cache_read.input_tokens" semantic conventions. It
	// represents the number of audio input tokens served from a provider-managed
	// cache.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 60
	// Note: The value SHOULD be included in `gen_ai.usage.cache_read.input_tokens`
	// and in `gen_ai.usage.audio.input_tokens`.
	GenAIUsageAudioCacheReadInputTokensKey = attribute.Key("gen_ai.usage.audio.cache_read.input_tokens")

	// GenAIUsageAudioInputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.audio.input_tokens" semantic conventions. It represents the
	// number of audio input tokens.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 120
	// Note: The value SHOULD be included in `gen_ai.usage.input_tokens`.
	GenAIUsageAudioInputTokensKey = attribute.Key("gen_ai.usage.audio.input_tokens")

	// GenAIUsageAudioOutputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.audio.output_tokens" semantic conventions. It represents the
	// number of audio output tokens.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 240
	// Note: The value SHOULD be included in `gen_ai.usage.output_tokens`.
	GenAIUsageAudioOutputTokensKey = attribute.Key("gen_ai.usage.audio.output_tokens")

	// GenAIUsageCacheReadInputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.cache_read.input_tokens" semantic conventions. It represents
	// the number of input tokens served from a provider-managed cache.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 50
	// Note: The value SHOULD be included in `gen_ai.usage.input_tokens`.
	GenAIUsageCacheReadInputTokensKey = attribute.Key("gen_ai.usage.cache_read.input_tokens")

	// GenAIUsageCacheWriteInputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.cache_write.input_tokens" semantic conventions. It represents
	// the number of input tokens written to a provider-managed cache.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 25
	// Note: The value SHOULD be included in `gen_ai.usage.input_tokens`.
	GenAIUsageCacheWriteInputTokensKey = attribute.Key("gen_ai.usage.cache_write.input_tokens")

	// GenAIUsageImageCacheReadInputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.image.cache_read.input_tokens" semantic conventions. It
	// represents the number of image input tokens served from a provider-managed
	// cache.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 128
	// Note: The value SHOULD be included in `gen_ai.usage.cache_read.input_tokens`
	// and in `gen_ai.usage.image.input_tokens`.
	GenAIUsageImageCacheReadInputTokensKey = attribute.Key("gen_ai.usage.image.cache_read.input_tokens")

	// GenAIUsageImageInputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.image.input_tokens" semantic conventions. It represents the
	// number of image input tokens.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 258
	// Note: The value SHOULD be included in `gen_ai.usage.input_tokens`.
	GenAIUsageImageInputTokensKey = attribute.Key("gen_ai.usage.image.input_tokens")

	// GenAIUsageImageOutputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.image.output_tokens" semantic conventions. It represents the
	// number of image output tokens.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 1290
	// Note: The value SHOULD be included in `gen_ai.usage.output_tokens`.
	GenAIUsageImageOutputTokensKey = attribute.Key("gen_ai.usage.image.output_tokens")

	// GenAIUsageInputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.input_tokens" semantic conventions. It represents the number of
	// tokens used in the GenAI input (prompt).
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 100
	// Note: This value SHOULD include all types of input tokens, including cached
	// tokens.
	// Instrumentations SHOULD make a best effort to populate this value, using a
	// total
	// provided by the provider when available or, depending on the provider API,
	// by summing different token types parsed from the provider output.
	//
	// When the provider reports both billed token counts and model-consumed
	// token counts (for example, Cohere exposes both `usage.billed_units` and
	// `usage.tokens`), instrumentations SHOULD report the billed count so the
	// value matches the units the customer is charged for.
	//
	// Detailed usage attributes are subsets of total and aggregate counts. For
	// example,
	// if a request has 100 text tokens (40 cached) and 200 image tokens:
	//
	//   - `gen_ai.usage.input_tokens`: 300
	//   - `gen_ai.usage.cache_read.input_tokens`: 40
	//   - `gen_ai.usage.text.input_tokens`: 100
	//   - `gen_ai.usage.text.cache_read.input_tokens`: 40
	//   - `gen_ai.usage.image.input_tokens`: 200
	GenAIUsageInputTokensKey = attribute.Key("gen_ai.usage.input_tokens")

	// GenAIUsageOutputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.output_tokens" semantic conventions. It represents the number
	// of tokens used in the GenAI response (completion).
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 180
	// Note: When the provider reports both billed token counts and model-consumed
	// token counts (for example, Cohere exposes both `usage.billed_units` and
	// `usage.tokens`), instrumentations SHOULD report the billed count so the
	// value matches the units the customer is charged for.
	GenAIUsageOutputTokensKey = attribute.Key("gen_ai.usage.output_tokens")

	// GenAIUsageReasoningOutputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.reasoning.output_tokens" semantic conventions. It represents
	// the number of output tokens used for reasoning (e.g. chain-of-thought,
	// extended thinking).
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 50
	// Note: The value SHOULD be included in `gen_ai.usage.output_tokens`.
	GenAIUsageReasoningOutputTokensKey = attribute.Key("gen_ai.usage.reasoning.output_tokens")

	// GenAIUsageTextCacheReadInputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.text.cache_read.input_tokens" semantic conventions. It
	// represents the number of text input tokens served from a provider-managed
	// cache.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 40
	// Note: The value SHOULD be included in `gen_ai.usage.cache_read.input_tokens`
	// and in `gen_ai.usage.text.input_tokens`.
	GenAIUsageTextCacheReadInputTokensKey = attribute.Key("gen_ai.usage.text.cache_read.input_tokens")

	// GenAIUsageTextInputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.text.input_tokens" semantic conventions. It represents the
	// number of text input tokens.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 100
	// Note: The value SHOULD be included in `gen_ai.usage.input_tokens`.
	GenAIUsageTextInputTokensKey = attribute.Key("gen_ai.usage.text.input_tokens")

	// GenAIUsageTextOutputTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.text.output_tokens" semantic conventions. It represents the
	// number of text output tokens.
	//
	// Type: int
	// Stability: Development
	//
	// Examples: 180
	// Note: The value SHOULD be included in `gen_ai.usage.output_tokens`.
	GenAIUsageTextOutputTokensKey = attribute.Key("gen_ai.usage.text.output_tokens")

	// GenAIWorkflowNameKey is the attribute Key conforming to the
	// "gen_ai.workflow.name" semantic conventions. It represents the human-readable
	// name of the GenAI workflow provided by the application.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "multi_agent_rag", "customer_support_pipeline"
	// Note: The workflow name is usually a static, application-unique identifier
	// defined
	// in a framework-specific way.
	//
	// For example, it can be the name of the first chain in LangChain,
	// the name of the crew in CrewAI, or the entry point agent in ADK or
	// OpenAI Agents when no explicit workflow name is provided.
	//
	// This attribute MUST have low cardinality. It is NOT RECOMMENDED to use
	// instrumentation-time constants or names of types representing the workflow,
	// such as "StateGraph". When no meaningful, low-cardinality workflow name is
	// available for a given framework, this attribute MUST NOT be captured by
	// default.
	//
	// Semantic conventions for individual Generative AI frameworks SHOULD document
	// what `gen_ai.workflow.name` means in the context of that framework.
	GenAIWorkflowNameKey = attribute.Key("gen_ai.workflow.name")
)

// GenAIAgentDescription returns an attribute KeyValue conforming to the
// "gen_ai.agent.description" semantic conventions. It represents the free-form
// description of the GenAI agent provided by the application.
func GenAIAgentDescription(val string) attribute.KeyValue {
	return GenAIAgentDescriptionKey.String(val)
}

// GenAIAgentID returns an attribute KeyValue conforming to the "gen_ai.agent.id"
// semantic conventions. It represents the unique and stable identifier of the
// GenAI hosted agent resource.
func GenAIAgentID(val string) attribute.KeyValue {
	return GenAIAgentIDKey.String(val)
}

// GenAIAgentName returns an attribute KeyValue conforming to the
// "gen_ai.agent.name" semantic conventions. It represents the human-readable
// name of the GenAI agent provided by the application.
func GenAIAgentName(val string) attribute.KeyValue {
	return GenAIAgentNameKey.String(val)
}

// GenAIAgentVersion returns an attribute KeyValue conforming to the
// "gen_ai.agent.version" semantic conventions. It represents the version of the
// GenAI agent.
func GenAIAgentVersion(val string) attribute.KeyValue {
	return GenAIAgentVersionKey.String(val)
}

// GenAIConversationCompacted returns an attribute KeyValue conforming to the
// "gen_ai.conversation.compacted" semantic conventions. It represents the
// indicates whether the effective conversation context used for this operation
// is a compacted view of a prior conversation.
func GenAIConversationCompacted(val bool) attribute.KeyValue {
	return GenAIConversationCompactedKey.Bool(val)
}

// GenAIConversationID returns an attribute KeyValue conforming to the
// "gen_ai.conversation.id" semantic conventions. It represents the unique
// identifier for a conversation (session, thread), used to store and correlate
// messages within this conversation.
func GenAIConversationID(val string) attribute.KeyValue {
	return GenAIConversationIDKey.String(val)
}

// GenAIDataSourceID returns an attribute KeyValue conforming to the
// "gen_ai.data_source.id" semantic conventions. It represents the data source
// identifier.
func GenAIDataSourceID(val string) attribute.KeyValue {
	return GenAIDataSourceIDKey.String(val)
}

// GenAIEmbeddingsDimensionCount returns an attribute KeyValue conforming to the
// "gen_ai.embeddings.dimension.count" semantic conventions. It represents the
// number of dimensions the resulting output embeddings should have.
func GenAIEmbeddingsDimensionCount(val int) attribute.KeyValue {
	return GenAIEmbeddingsDimensionCountKey.Int(val)
}

// GenAIEvaluationExplanation returns an attribute KeyValue conforming to the
// "gen_ai.evaluation.explanation" semantic conventions. It represents a
// free-form explanation for the assigned score provided by the evaluator.
func GenAIEvaluationExplanation(val string) attribute.KeyValue {
	return GenAIEvaluationExplanationKey.String(val)
}

// GenAIEvaluationName returns an attribute KeyValue conforming to the
// "gen_ai.evaluation.name" semantic conventions. It represents the name of the
// evaluation metric used for the GenAI response.
func GenAIEvaluationName(val string) attribute.KeyValue {
	return GenAIEvaluationNameKey.String(val)
}

// GenAIEvaluationScoreLabel returns an attribute KeyValue conforming to the
// "gen_ai.evaluation.score.label" semantic conventions. It represents the human
// readable label for evaluation.
func GenAIEvaluationScoreLabel(val string) attribute.KeyValue {
	return GenAIEvaluationScoreLabelKey.String(val)
}

// GenAIEvaluationScoreValue returns an attribute KeyValue conforming to the
// "gen_ai.evaluation.score.value" semantic conventions. It represents the
// evaluation score returned by the evaluator.
func GenAIEvaluationScoreValue(val float64) attribute.KeyValue {
	return GenAIEvaluationScoreValueKey.Float64(val)
}

// GenAIMainAgentDescription returns an attribute KeyValue conforming to the
// "gen_ai.main_agent.description" semantic conventions. It represents the
// free-form description of the top-level Generative AI agent in the process.
func GenAIMainAgentDescription(val string) attribute.KeyValue {
	return GenAIMainAgentDescriptionKey.String(val)
}

// GenAIMainAgentID returns an attribute KeyValue conforming to the
// "gen_ai.main_agent.id" semantic conventions. It represents the unique and
// stable identifier of the top-level Generative AI agent in the process.
func GenAIMainAgentID(val string) attribute.KeyValue {
	return GenAIMainAgentIDKey.String(val)
}

// GenAIMainAgentName returns an attribute KeyValue conforming to the
// "gen_ai.main_agent.name" semantic conventions. It represents the
// human-readable name of the top-level Generative AI agent in the process.
func GenAIMainAgentName(val string) attribute.KeyValue {
	return GenAIMainAgentNameKey.String(val)
}

// GenAIMemoryQueryText returns an attribute KeyValue conforming to the
// "gen_ai.memory.query.text" semantic conventions. It represents the search
// query used to retrieve memories.
func GenAIMemoryQueryText(val string) attribute.KeyValue {
	return GenAIMemoryQueryTextKey.String(val)
}

// GenAIMemoryRecordCount returns an attribute KeyValue conforming to the
// "gen_ai.memory.record.count" semantic conventions. It represents the number of
// memory records relevant to the operation.
func GenAIMemoryRecordCount(val int) attribute.KeyValue {
	return GenAIMemoryRecordCountKey.Int(val)
}

// GenAIMemoryRecordID returns an attribute KeyValue conforming to the
// "gen_ai.memory.record.id" semantic conventions. It represents the unique
// identifier of the memory record.
func GenAIMemoryRecordID(val string) attribute.KeyValue {
	return GenAIMemoryRecordIDKey.String(val)
}

// GenAIMemoryStoreID returns an attribute KeyValue conforming to the
// "gen_ai.memory.store.id" semantic conventions. It represents the unique
// identifier of the memory store.
func GenAIMemoryStoreID(val string) attribute.KeyValue {
	return GenAIMemoryStoreIDKey.String(val)
}

// GenAIPromptName returns an attribute KeyValue conforming to the
// "gen_ai.prompt.name" semantic conventions. It represents the name of the
// prompt that uniquely identifies it.
func GenAIPromptName(val string) attribute.KeyValue {
	return GenAIPromptNameKey.String(val)
}

// GenAIPromptVariable returns an attribute KeyValue conforming to the
// "gen_ai.prompt.variable" semantic conventions. It represents the variables
// supplied to the prompt template, the `<key>` being the variable name, the
// value being the variable value.
func GenAIPromptVariable(key string, val string) attribute.KeyValue {
	return attribute.String("gen_ai.prompt.variable."+key, val)
}

// GenAIPromptVersion returns an attribute KeyValue conforming to the
// "gen_ai.prompt.version" semantic conventions. It represents the version of the
// prompt template used.
func GenAIPromptVersion(val string) attribute.KeyValue {
	return GenAIPromptVersionKey.String(val)
}

// GenAIRequestChoiceCount returns an attribute KeyValue conforming to the
// "gen_ai.request.choice.count" semantic conventions. It represents the target
// number of candidate completions to return.
func GenAIRequestChoiceCount(val int) attribute.KeyValue {
	return GenAIRequestChoiceCountKey.Int(val)
}

// GenAIRequestEncodingFormats returns an attribute KeyValue conforming to the
// "gen_ai.request.encoding_formats" semantic conventions. It represents the
// encoding formats requested in an embeddings operation, if specified.
func GenAIRequestEncodingFormats(val ...string) attribute.KeyValue {
	return GenAIRequestEncodingFormatsKey.StringSlice(val)
}

// GenAIRequestFrequencyPenalty returns an attribute KeyValue conforming to the
// "gen_ai.request.frequency_penalty" semantic conventions. It represents the
// frequency penalty setting for the GenAI request.
func GenAIRequestFrequencyPenalty(val float64) attribute.KeyValue {
	return GenAIRequestFrequencyPenaltyKey.Float64(val)
}

// GenAIRequestMaxTokens returns an attribute KeyValue conforming to the
// "gen_ai.request.max_tokens" semantic conventions. It represents the maximum
// number of tokens the model generates for a request.
func GenAIRequestMaxTokens(val int) attribute.KeyValue {
	return GenAIRequestMaxTokensKey.Int(val)
}

// GenAIRequestModel returns an attribute KeyValue conforming to the
// "gen_ai.request.model" semantic conventions. It represents the name of the
// GenAI model a request is being made to.
func GenAIRequestModel(val string) attribute.KeyValue {
	return GenAIRequestModelKey.String(val)
}

// GenAIRequestPresencePenalty returns an attribute KeyValue conforming to the
// "gen_ai.request.presence_penalty" semantic conventions. It represents the
// presence penalty setting for the GenAI request.
func GenAIRequestPresencePenalty(val float64) attribute.KeyValue {
	return GenAIRequestPresencePenaltyKey.Float64(val)
}

// GenAIRequestPreviousResponseID returns an attribute KeyValue conforming to the
// "gen_ai.request.previous_response.id" semantic conventions. It represents the
// unique identifier of a previous response or interaction used to provide
// context for the current operation.
func GenAIRequestPreviousResponseID(val string) attribute.KeyValue {
	return GenAIRequestPreviousResponseIDKey.String(val)
}

// GenAIRequestReasoningLevel returns an attribute KeyValue conforming to the
// "gen_ai.request.reasoning.level" semantic conventions. It represents the
// reasoning or thinking effort level requested for a GenAI model.
func GenAIRequestReasoningLevel(val string) attribute.KeyValue {
	return GenAIRequestReasoningLevelKey.String(val)
}

// GenAIRequestSeed returns an attribute KeyValue conforming to the
// "gen_ai.request.seed" semantic conventions. It represents the requests with
// same seed value more likely to return same result.
func GenAIRequestSeed(val int) attribute.KeyValue {
	return GenAIRequestSeedKey.Int(val)
}

// GenAIRequestStopSequences returns an attribute KeyValue conforming to the
// "gen_ai.request.stop_sequences" semantic conventions. It represents the list
// of sequences that the model will use to stop generating further tokens.
func GenAIRequestStopSequences(val ...string) attribute.KeyValue {
	return GenAIRequestStopSequencesKey.StringSlice(val)
}

// GenAIRequestStream returns an attribute KeyValue conforming to the
// "gen_ai.request.stream" semantic conventions. It represents the indicates
// whether the GenAI request was made in streaming mode.
func GenAIRequestStream(val bool) attribute.KeyValue {
	return GenAIRequestStreamKey.Bool(val)
}

// GenAIRequestStreamCursor returns an attribute KeyValue conforming to the
// "gen_ai.request.stream_cursor" semantic conventions. It represents the cursor
// identifying the last streamed event already received, used to resume a
// streamed response from that position.
func GenAIRequestStreamCursor(val string) attribute.KeyValue {
	return GenAIRequestStreamCursorKey.String(val)
}

// GenAIRequestTemperature returns an attribute KeyValue conforming to the
// "gen_ai.request.temperature" semantic conventions. It represents the
// temperature setting for the GenAI request.
func GenAIRequestTemperature(val float64) attribute.KeyValue {
	return GenAIRequestTemperatureKey.Float64(val)
}

// GenAIRequestTopK returns an attribute KeyValue conforming to the
// "gen_ai.request.top_k" semantic conventions. It represents the top-K sampling
// setting for the GenAI request: restricts token generation at each step to the
// K most likely next tokens.
func GenAIRequestTopK(val int) attribute.KeyValue {
	return GenAIRequestTopKKey.Int(val)
}

// GenAIRequestTopP returns an attribute KeyValue conforming to the
// "gen_ai.request.top_p" semantic conventions. It represents the top_p sampling
// setting for the GenAI request.
func GenAIRequestTopP(val float64) attribute.KeyValue {
	return GenAIRequestTopPKey.Float64(val)
}

// GenAIResponseFinishReasons returns an attribute KeyValue conforming to the
// "gen_ai.response.finish_reasons" semantic conventions. It represents the array
// of reasons the model stopped generating tokens, corresponding to each
// generation received.
func GenAIResponseFinishReasons(val ...string) attribute.KeyValue {
	return GenAIResponseFinishReasonsKey.StringSlice(val)
}

// GenAIResponseID returns an attribute KeyValue conforming to the
// "gen_ai.response.id" semantic conventions. It represents the unique identifier
// for the completion.
func GenAIResponseID(val string) attribute.KeyValue {
	return GenAIResponseIDKey.String(val)
}

// GenAIResponseModel returns an attribute KeyValue conforming to the
// "gen_ai.response.model" semantic conventions. It represents the name of the
// model that generated the response.
func GenAIResponseModel(val string) attribute.KeyValue {
	return GenAIResponseModelKey.String(val)
}

// GenAIResponseTimeToFirstChunk returns an attribute KeyValue conforming to the
// "gen_ai.response.time_to_first_chunk" semantic conventions. It represents the
// time to first chunk in a streaming response, measured from request issuance,
// in seconds. The value is measured from when the client issues the generation
// request to when the first chunk is received in the response stream.
func GenAIResponseTimeToFirstChunk(val float64) attribute.KeyValue {
	return GenAIResponseTimeToFirstChunkKey.Float64(val)
}

// GenAIRetrievalQueryText returns an attribute KeyValue conforming to the
// "gen_ai.retrieval.query.text" semantic conventions. It represents the query
// text used for retrieval.
func GenAIRetrievalQueryText(val string) attribute.KeyValue {
	return GenAIRetrievalQueryTextKey.String(val)
}

// GenAIRetrievalTopK returns an attribute KeyValue conforming to the
// "gen_ai.retrieval.top_k" semantic conventions. It represents the maximum
// number of documents the retriever was asked to return for the query (also
// known as `k`, `limit`, or `max_num_results`).
func GenAIRetrievalTopK(val int) attribute.KeyValue {
	return GenAIRetrievalTopKKey.Int(val)
}

// GenAISkillDescription returns an attribute KeyValue conforming to the
// "gen_ai.skill.description" semantic conventions. It represents the
// [description of the Agent Skill].
//
// [description of the Agent Skill]: https://agentskills.io/specification#skill-md-format
func GenAISkillDescription(val string) attribute.KeyValue {
	return GenAISkillDescriptionKey.String(val)
}

// GenAISkillName returns an attribute KeyValue conforming to the
// "gen_ai.skill.name" semantic conventions. It represents the name of the
// [Agent Skill].
//
// [Agent Skill]: https://agentskills.io/specification#skill-md-format
func GenAISkillName(val string) attribute.KeyValue {
	return GenAISkillNameKey.String(val)
}

// GenAISkillResourceName returns an attribute KeyValue conforming to the
// "gen_ai.skill.resource.name" semantic conventions. It represents the
// skill-relative name of the resource (script, reference, or asset) being
// accessed or executed.
func GenAISkillResourceName(val string) attribute.KeyValue {
	return GenAISkillResourceNameKey.String(val)
}

// GenAISkillSourceURI returns an attribute KeyValue conforming to the
// "gen_ai.skill.source.uri" semantic conventions. It represents the source URI
// for loading the skill.
func GenAISkillSourceURI(val string) attribute.KeyValue {
	return GenAISkillSourceURIKey.String(val)
}

// GenAIToolCallID returns an attribute KeyValue conforming to the
// "gen_ai.tool.call.id" semantic conventions. It represents the tool call
// identifier supplied by the model or provider.
func GenAIToolCallID(val string) attribute.KeyValue {
	return GenAIToolCallIDKey.String(val)
}

// GenAIToolDescription returns an attribute KeyValue conforming to the
// "gen_ai.tool.description" semantic conventions. It represents the tool
// description.
func GenAIToolDescription(val string) attribute.KeyValue {
	return GenAIToolDescriptionKey.String(val)
}

// GenAIToolName returns an attribute KeyValue conforming to the
// "gen_ai.tool.name" semantic conventions. It represents the name of the tool
// utilized by the agent.
func GenAIToolName(val string) attribute.KeyValue {
	return GenAIToolNameKey.String(val)
}

// GenAIToolType returns an attribute KeyValue conforming to the
// "gen_ai.tool.type" semantic conventions. It represents the type of the tool
// utilized by the agent.
func GenAIToolType(val string) attribute.KeyValue {
	return GenAIToolTypeKey.String(val)
}

// GenAITransferTargetName returns an attribute KeyValue conforming to the
// "gen_ai.transfer.target.name" semantic conventions. It represents the
// human-readable name of the agent targeted by the transfer.
func GenAITransferTargetName(val string) attribute.KeyValue {
	return GenAITransferTargetNameKey.String(val)
}

// GenAIUsageAudioCacheReadInputTokens returns an attribute KeyValue conforming
// to the "gen_ai.usage.audio.cache_read.input_tokens" semantic conventions. It
// represents the number of audio input tokens served from a provider-managed
// cache.
func GenAIUsageAudioCacheReadInputTokens(val int) attribute.KeyValue {
	return GenAIUsageAudioCacheReadInputTokensKey.Int(val)
}

// GenAIUsageAudioInputTokens returns an attribute KeyValue conforming to the
// "gen_ai.usage.audio.input_tokens" semantic conventions. It represents the
// number of audio input tokens.
func GenAIUsageAudioInputTokens(val int) attribute.KeyValue {
	return GenAIUsageAudioInputTokensKey.Int(val)
}

// GenAIUsageAudioOutputTokens returns an attribute KeyValue conforming to the
// "gen_ai.usage.audio.output_tokens" semantic conventions. It represents the
// number of audio output tokens.
func GenAIUsageAudioOutputTokens(val int) attribute.KeyValue {
	return GenAIUsageAudioOutputTokensKey.Int(val)
}

// GenAIUsageCacheReadInputTokens returns an attribute KeyValue conforming to the
// "gen_ai.usage.cache_read.input_tokens" semantic conventions. It represents the
// number of input tokens served from a provider-managed cache.
func GenAIUsageCacheReadInputTokens(val int) attribute.KeyValue {
	return GenAIUsageCacheReadInputTokensKey.Int(val)
}

// GenAIUsageCacheWriteInputTokens returns an attribute KeyValue conforming to
// the "gen_ai.usage.cache_write.input_tokens" semantic conventions. It
// represents the number of input tokens written to a provider-managed cache.
func GenAIUsageCacheWriteInputTokens(val int) attribute.KeyValue {
	return GenAIUsageCacheWriteInputTokensKey.Int(val)
}

// GenAIUsageImageCacheReadInputTokens returns an attribute KeyValue conforming
// to the "gen_ai.usage.image.cache_read.input_tokens" semantic conventions. It
// represents the number of image input tokens served from a provider-managed
// cache.
func GenAIUsageImageCacheReadInputTokens(val int) attribute.KeyValue {
	return GenAIUsageImageCacheReadInputTokensKey.Int(val)
}

// GenAIUsageImageInputTokens returns an attribute KeyValue conforming to the
// "gen_ai.usage.image.input_tokens" semantic conventions. It represents the
// number of image input tokens.
func GenAIUsageImageInputTokens(val int) attribute.KeyValue {
	return GenAIUsageImageInputTokensKey.Int(val)
}

// GenAIUsageImageOutputTokens returns an attribute KeyValue conforming to the
// "gen_ai.usage.image.output_tokens" semantic conventions. It represents the
// number of image output tokens.
func GenAIUsageImageOutputTokens(val int) attribute.KeyValue {
	return GenAIUsageImageOutputTokensKey.Int(val)
}

// GenAIUsageInputTokens returns an attribute KeyValue conforming to the
// "gen_ai.usage.input_tokens" semantic conventions. It represents the number of
// tokens used in the GenAI input (prompt).
func GenAIUsageInputTokens(val int) attribute.KeyValue {
	return GenAIUsageInputTokensKey.Int(val)
}

// GenAIUsageOutputTokens returns an attribute KeyValue conforming to the
// "gen_ai.usage.output_tokens" semantic conventions. It represents the number of
// tokens used in the GenAI response (completion).
func GenAIUsageOutputTokens(val int) attribute.KeyValue {
	return GenAIUsageOutputTokensKey.Int(val)
}

// GenAIUsageReasoningOutputTokens returns an attribute KeyValue conforming to
// the "gen_ai.usage.reasoning.output_tokens" semantic conventions. It represents
// the number of output tokens used for reasoning (e.g. chain-of-thought,
// extended thinking).
func GenAIUsageReasoningOutputTokens(val int) attribute.KeyValue {
	return GenAIUsageReasoningOutputTokensKey.Int(val)
}

// GenAIUsageTextCacheReadInputTokens returns an attribute KeyValue conforming to
// the "gen_ai.usage.text.cache_read.input_tokens" semantic conventions. It
// represents the number of text input tokens served from a provider-managed
// cache.
func GenAIUsageTextCacheReadInputTokens(val int) attribute.KeyValue {
	return GenAIUsageTextCacheReadInputTokensKey.Int(val)
}

// GenAIUsageTextInputTokens returns an attribute KeyValue conforming to the
// "gen_ai.usage.text.input_tokens" semantic conventions. It represents the
// number of text input tokens.
func GenAIUsageTextInputTokens(val int) attribute.KeyValue {
	return GenAIUsageTextInputTokensKey.Int(val)
}

// GenAIUsageTextOutputTokens returns an attribute KeyValue conforming to the
// "gen_ai.usage.text.output_tokens" semantic conventions. It represents the
// number of text output tokens.
func GenAIUsageTextOutputTokens(val int) attribute.KeyValue {
	return GenAIUsageTextOutputTokensKey.Int(val)
}

// GenAIWorkflowName returns an attribute KeyValue conforming to the
// "gen_ai.workflow.name" semantic conventions. It represents the human-readable
// name of the GenAI workflow provided by the application.
func GenAIWorkflowName(val string) attribute.KeyValue {
	return GenAIWorkflowNameKey.String(val)
}

// Enum values for gen_ai.operation.name
var (
	// Chat completion operation such as [OpenAI Chat API]
	// Stability: development
	//
	// [OpenAI Chat API]: https://platform.openai.com/docs/api-reference/chat
	GenAIOperationNameChat = GenAIOperationNameKey.String("chat")
	// Multimodal content generation operation such as [Gemini Generate Content]
	// Stability: development
	//
	// [Gemini Generate Content]: https://ai.google.dev/api/generate-content
	GenAIOperationNameGenerateContent = GenAIOperationNameKey.String("generate_content")
	// Text completions operation such as [OpenAI Completions API (Legacy)]
	// Stability: development
	//
	// [OpenAI Completions API (Legacy)]: https://platform.openai.com/docs/api-reference/completions
	GenAIOperationNameTextCompletion = GenAIOperationNameKey.String("text_completion")
	// Embeddings operation such as [OpenAI Create embeddings API]
	// Stability: development
	//
	// [OpenAI Create embeddings API]: https://platform.openai.com/docs/api-reference/embeddings/create
	GenAIOperationNameEmbeddings = GenAIOperationNameKey.String("embeddings")
	// Retrieval operation such as [OpenAI Search Vector Store API]
	// Stability: development
	//
	// [OpenAI Search Vector Store API]: https://platform.openai.com/docs/api-reference/vector-stores/search
	GenAIOperationNameRetrieval = GenAIOperationNameKey.String("retrieval")
	// Fetch a previously generated model response by its identifier, without
	// performing inference, such as [OpenAI Get a model response]
	// Stability: development
	//
	// [OpenAI Get a model response]: https://platform.openai.com/docs/api-reference/responses/get
	GenAIOperationNameFetchResponse = GenAIOperationNameKey.String("fetch_response")
	// Create GenAI agent
	// Stability: development
	GenAIOperationNameCreateAgent = GenAIOperationNameKey.String("create_agent")
	// Invoke GenAI agent
	// Stability: development
	GenAIOperationNameInvokeAgent = GenAIOperationNameKey.String("invoke_agent")
	// Execute a tool
	// Stability: development
	GenAIOperationNameExecuteTool = GenAIOperationNameKey.String("execute_tool")
	// Invoke GenAI workflow
	// Stability: development
	GenAIOperationNameInvokeWorkflow = GenAIOperationNameKey.String("invoke_workflow")
	// Agent planning or task decomposition phase
	// Stability: development
	GenAIOperationNamePlan = GenAIOperationNameKey.String("plan")
	// Search/query memories from a memory store
	// Stability: development
	GenAIOperationNameSearchMemory = GenAIOperationNameKey.String("search_memory")
	// Create new memory records
	// Stability: development
	GenAIOperationNameCreateMemory = GenAIOperationNameKey.String("create_memory")
	// Update existing memory records
	// Stability: development
	GenAIOperationNameUpdateMemory = GenAIOperationNameKey.String("update_memory")
	// Create or update memory records without the caller choosing which
	// Stability: development
	GenAIOperationNameUpsertMemory = GenAIOperationNameKey.String("upsert_memory")
	// Delete memory records
	// Stability: development
	GenAIOperationNameDeleteMemory = GenAIOperationNameKey.String("delete_memory")
	// Create or initialize a memory store
	// Stability: development
	GenAIOperationNameCreateMemoryStore = GenAIOperationNameKey.String("create_memory_store")
	// Delete or deprovision a memory store
	// Stability: development
	GenAIOperationNameDeleteMemoryStore = GenAIOperationNameKey.String("delete_memory_store")
)

// Enum values for gen_ai.output.type
var (
	// Plain text
	// Stability: development
	GenAIOutputTypeText = GenAIOutputTypeKey.String("text")
	// JSON object with known or unknown schema
	// Stability: development
	GenAIOutputTypeJSON = GenAIOutputTypeKey.String("json")
	// Image
	// Stability: development
	GenAIOutputTypeImage = GenAIOutputTypeKey.String("image")
	// Speech
	// Stability: development
	GenAIOutputTypeSpeech = GenAIOutputTypeKey.String("speech")
)

// Enum values for gen_ai.provider.name
var (
	// [OpenAI]
	// Stability: development
	//
	// [OpenAI]: https://openai.com/
	GenAIProviderNameOpenAI = GenAIProviderNameKey.String("openai")
	// Any Google generative AI endpoint
	// Stability: development
	GenAIProviderNameGCPGenAI = GenAIProviderNameKey.String("gcp.gen_ai")
	// [Vertex AI]
	// Stability: development
	//
	// [Vertex AI]: https://cloud.google.com/vertex-ai
	GenAIProviderNameGCPVertexAI = GenAIProviderNameKey.String("gcp.vertex_ai")
	// [Gemini]
	// Stability: development
	//
	// [Gemini]: https://cloud.google.com/products/gemini
	GenAIProviderNameGCPGemini = GenAIProviderNameKey.String("gcp.gemini")
	// [Anthropic]
	// Stability: development
	//
	// [Anthropic]: https://www.anthropic.com/
	GenAIProviderNameAnthropic = GenAIProviderNameKey.String("anthropic")
	// [Cohere]
	// Stability: development
	//
	// [Cohere]: https://cohere.com/
	GenAIProviderNameCohere = GenAIProviderNameKey.String("cohere")
	// Azure AI Inference
	// Stability: development
	GenAIProviderNameAzureAIInference = GenAIProviderNameKey.String("azure.ai.inference")
	// [Azure OpenAI]
	// Stability: development
	//
	// [Azure OpenAI]: https://learn.microsoft.com/en-us/azure/ai-services/openai/overview
	GenAIProviderNameAzureAIOpenAI = GenAIProviderNameKey.String("azure.ai.openai")
	// [IBM Watsonx AI]
	// Stability: development
	//
	// [IBM Watsonx AI]: https://www.ibm.com/products/watsonx-ai
	GenAIProviderNameIBMWatsonxAI = GenAIProviderNameKey.String("ibm.watsonx.ai")
	// [AWS Bedrock]
	// Stability: development
	//
	// [AWS Bedrock]: https://aws.amazon.com/bedrock
	GenAIProviderNameAWSBedrock = GenAIProviderNameKey.String("aws.bedrock")
	// [Perplexity]
	// Stability: development
	//
	// [Perplexity]: https://www.perplexity.ai/
	GenAIProviderNamePerplexity = GenAIProviderNameKey.String("perplexity")
	// [xAI]
	// Stability: development
	//
	// [xAI]: https://x.ai/
	GenAIProviderNameXAI = GenAIProviderNameKey.String("x_ai")
	// [DeepSeek]
	// Stability: development
	//
	// [DeepSeek]: https://www.deepseek.com/
	GenAIProviderNameDeepseek = GenAIProviderNameKey.String("deepseek")
	// [Groq]
	// Stability: development
	//
	// [Groq]: https://groq.com/
	GenAIProviderNameGroq = GenAIProviderNameKey.String("groq")
	// [Mistral]
	// Stability: development
	//
	// [Mistral]: https://mistral.ai/
	GenAIProviderNameMistral = GenAIProviderNameKey.String("mistral")
	// [Moonshot AI]
	// Stability: development
	//
	// [Moonshot AI]: https://www.moonshot.ai/
	GenAIProviderNameMoonshotAI = GenAIProviderNameKey.String("moonshot_ai")
)

// Enum values for gen_ai.response.status
var (
	// The response has been accepted by the provider but generation has not started
	// yet.
	// Stability: development
	GenAIResponseStatusQueued = GenAIResponseStatusKey.String("queued")
	// The response is still being generated, for example a background or streamed
	// response that has not finished.
	// Stability: development
	GenAIResponseStatusInProgress = GenAIResponseStatusKey.String("in_progress")
	// The response finished generating successfully.
	// Stability: development
	GenAIResponseStatusCompleted = GenAIResponseStatusKey.String("completed")
	// The response stopped before generation completed, for example because a token
	// limit or content filter was reached.
	// Stability: development
	GenAIResponseStatusIncomplete = GenAIResponseStatusKey.String("incomplete")
	// The response generation failed with an error.
	// Stability: development
	GenAIResponseStatusFailed = GenAIResponseStatusKey.String("failed")
	// The response generation was cancelled before it completed.
	// Stability: development
	GenAIResponseStatusCancelled = GenAIResponseStatusKey.String("cancelled")
)

// Enum values for gen_ai.token.modality
var (
	// Text tokens.
	// Stability: development
	GenAITokenModalityText = GenAITokenModalityKey.String("text")
	// Image tokens.
	// Stability: development
	GenAITokenModalityImage = GenAITokenModalityKey.String("image")
	// Audio tokens.
	// Stability: development
	GenAITokenModalityAudio = GenAITokenModalityKey.String("audio")
	// The modality is not known.
	// Stability: development
	GenAITokenModalityUnknown = GenAITokenModalityKey.String("unknown")
)

// Enum values for gen_ai.transfer.mode
var (
	// The initiator waits for the target to finish and then resumes execution.
	// Stability: development
	GenAITransferModeReturnToCaller = GenAITransferModeKey.String("return_to_caller")
	// The initiator does not wait for the target to finish. The target takes
	// control of the remaining work.
	// Stability: development
	GenAITransferModePassControl = GenAITransferModeKey.String("pass_control")
)

// Namespace: mcp
const (
	// McpMethodNameKey is the attribute Key conforming to the "mcp.method.name"
	// semantic conventions. It represents the name of the request or notification
	// method.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples:
	McpMethodNameKey = attribute.Key("mcp.method.name")

	// McpProtocolVersionKey is the attribute Key conforming to the
	// "mcp.protocol.version" semantic conventions. It represents the [version] of
	// the Model Context Protocol used.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "2025-06-18"
	//
	// [version]: https://modelcontextprotocol.io/specification/versioning
	McpProtocolVersionKey = attribute.Key("mcp.protocol.version")

	// McpResourceURIKey is the attribute Key conforming to the "mcp.resource.uri"
	// semantic conventions. It represents the value of the resource uri.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "postgres://database/customers/schema",
	// "file:///home/user/documents/report.pdf"
	// Note: This is a URI of the resource provided in the following requests or
	// notifications: `resources/read`, `resources/subscribe`,
	// `resources/unsubscribe`, or `notifications/resources/updated`.
	McpResourceURIKey = attribute.Key("mcp.resource.uri")

	// McpSessionIDKey is the attribute Key conforming to the "mcp.session.id"
	// semantic conventions. It represents the identifies [MCP session].
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "191c4850af6c49e08843a3f6c80e5046"
	//
	// [MCP session]: https://modelcontextprotocol.io/specification/2025-06-18/basic/transports#session-management
	McpSessionIDKey = attribute.Key("mcp.session.id")
)

// McpProtocolVersion returns an attribute KeyValue conforming to the
// "mcp.protocol.version" semantic conventions. It represents the [version] of
// the Model Context Protocol used.
//
// [version]: https://modelcontextprotocol.io/specification/versioning
func McpProtocolVersion(val string) attribute.KeyValue {
	return McpProtocolVersionKey.String(val)
}

// McpResourceURI returns an attribute KeyValue conforming to the
// "mcp.resource.uri" semantic conventions. It represents the value of the
// resource uri.
func McpResourceURI(val string) attribute.KeyValue {
	return McpResourceURIKey.String(val)
}

// McpSessionID returns an attribute KeyValue conforming to the "mcp.session.id"
// semantic conventions. It represents the identifies [MCP session].
//
// [MCP session]: https://modelcontextprotocol.io/specification/2025-06-18/basic/transports#session-management
func McpSessionID(val string) attribute.KeyValue {
	return McpSessionIDKey.String(val)
}

// Enum values for mcp.method.name
var (
	// Notification cancelling a previously-issued request.
	//
	// Stability: development
	McpMethodNameNotificationsCancelled = McpMethodNameKey.String("notifications/cancelled")
	// Request to initialize the MCP client.
	//
	// Stability: development
	McpMethodNameInitialize = McpMethodNameKey.String("initialize")
	// Notification indicating that the MCP client has been initialized.
	//
	// Stability: development
	McpMethodNameNotificationsInitialized = McpMethodNameKey.String("notifications/initialized")
	// Notification indicating the progress for a long-running operation.
	//
	// Stability: development
	McpMethodNameNotificationsProgress = McpMethodNameKey.String("notifications/progress")
	// Request to check that the other party is still alive.
	//
	// Stability: development
	McpMethodNamePing = McpMethodNameKey.String("ping")
	// Request to list resources available on server.
	//
	// Stability: development
	McpMethodNameResourcesList = McpMethodNameKey.String("resources/list")
	// Request to list resource templates available on server.
	//
	// Stability: development
	McpMethodNameResourcesTemplatesList = McpMethodNameKey.String("resources/templates/list")
	// Request to read a resource.
	//
	// Stability: development
	McpMethodNameResourcesRead = McpMethodNameKey.String("resources/read")
	// Notification indicating that the list of resources has changed.
	//
	// Stability: development
	McpMethodNameNotificationsResourcesListChanged = McpMethodNameKey.String("notifications/resources/list_changed")
	// Request to subscribe to a resource.
	//
	// Stability: development
	McpMethodNameResourcesSubscribe = McpMethodNameKey.String("resources/subscribe")
	// Request to unsubscribe from resource updates.
	//
	// Stability: development
	McpMethodNameResourcesUnsubscribe = McpMethodNameKey.String("resources/unsubscribe")
	// Notification indicating that a resource has been updated.
	//
	// Stability: development
	McpMethodNameNotificationsResourcesUpdated = McpMethodNameKey.String("notifications/resources/updated")
	// Request to list prompts available on server.
	//
	// Stability: development
	McpMethodNamePromptsList = McpMethodNameKey.String("prompts/list")
	// Request to get a prompt.
	//
	// Stability: development
	McpMethodNamePromptsGet = McpMethodNameKey.String("prompts/get")
	// Notification indicating that the list of prompts has changed.
	//
	// Stability: development
	McpMethodNameNotificationsPromptsListChanged = McpMethodNameKey.String("notifications/prompts/list_changed")
	// Request to list tools available on server.
	//
	// Stability: development
	McpMethodNameToolsList = McpMethodNameKey.String("tools/list")
	// Request to call a tool.
	//
	// Stability: development
	McpMethodNameToolsCall = McpMethodNameKey.String("tools/call")
	// Notification indicating that the list of tools has changed.
	//
	// Stability: development
	McpMethodNameNotificationsToolsListChanged = McpMethodNameKey.String("notifications/tools/list_changed")
	// Request to set the logging level.
	//
	// Stability: development
	McpMethodNameLoggingSetLevel = McpMethodNameKey.String("logging/setLevel")
	// Notification indicating that a message has been received.
	//
	// Stability: development
	McpMethodNameNotificationsMessage = McpMethodNameKey.String("notifications/message")
	// Request to create a sampling message.
	//
	// Stability: development
	McpMethodNameSamplingCreateMessage = McpMethodNameKey.String("sampling/createMessage")
	// Request to complete a prompt.
	//
	// Stability: development
	McpMethodNameCompletionComplete = McpMethodNameKey.String("completion/complete")
	// Request to list roots available on server.
	//
	// Stability: development
	McpMethodNameRootsList = McpMethodNameKey.String("roots/list")
	// Notification indicating that the list of roots has changed.
	//
	// Stability: development
	McpMethodNameNotificationsRootsListChanged = McpMethodNameKey.String("notifications/roots/list_changed")
	// Request from the server to elicit additional information from the user via
	// the client
	//
	// Stability: development
	McpMethodNameElicitationCreate = McpMethodNameKey.String("elicitation/create")
)

// Namespace: openai
const (
	// OpenAIAPITypeKey is the attribute Key conforming to the "openai.api.type"
	// semantic conventions. It represents the type of OpenAI API being used.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples:
	OpenAIAPITypeKey = attribute.Key("openai.api.type")

	// OpenAIRequestServiceTierKey is the attribute Key conforming to the
	// "openai.request.service_tier" semantic conventions. It represents the service
	// tier requested. May be a specific tier, default, or auto.
	//
	// Type: Enum
	// Stability: Development
	//
	// Examples: "auto", "default"
	OpenAIRequestServiceTierKey = attribute.Key("openai.request.service_tier")

	// OpenAIResponseServiceTierKey is the attribute Key conforming to the
	// "openai.response.service_tier" semantic conventions. It represents the
	// service tier used for the response.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "scale", "default"
	OpenAIResponseServiceTierKey = attribute.Key("openai.response.service_tier")

	// OpenAIResponseSystemFingerprintKey is the attribute Key conforming to the
	// "openai.response.system_fingerprint" semantic conventions. It represents a
	// fingerprint to track any eventual change in the Generative AI environment.
	//
	// Type: string
	// Stability: Development
	//
	// Examples: "fp_44709d6fcb"
	OpenAIResponseSystemFingerprintKey = attribute.Key("openai.response.system_fingerprint")
)

// OpenAIResponseServiceTier returns an attribute KeyValue conforming to the
// "openai.response.service_tier" semantic conventions. It represents the service
// tier used for the response.
func OpenAIResponseServiceTier(val string) attribute.KeyValue {
	return OpenAIResponseServiceTierKey.String(val)
}

// OpenAIResponseSystemFingerprint returns an attribute KeyValue conforming to
// the "openai.response.system_fingerprint" semantic conventions. It represents a
// fingerprint to track any eventual change in the Generative AI environment.
func OpenAIResponseSystemFingerprint(val string) attribute.KeyValue {
	return OpenAIResponseSystemFingerprintKey.String(val)
}

// Enum values for openai.api.type
var (
	// The OpenAI [Chat Completions API].
	// Stability: development
	//
	// [Chat Completions API]: https://developers.openai.com/api/reference/chat-completions/overview
	OpenAIAPITypeChatCompletions = OpenAIAPITypeKey.String("chat_completions")
	// The OpenAI [Responses API].
	// Stability: development
	//
	// [Responses API]: https://developers.openai.com/api/reference/responses/overview
	OpenAIAPITypeResponses = OpenAIAPITypeKey.String("responses")
)

// Enum values for openai.request.service_tier
var (
	// The system will utilize scale tier credits until they are exhausted.
	// Stability: development
	OpenAIRequestServiceTierAuto = OpenAIRequestServiceTierKey.String("auto")
	// The system will utilize the default scale tier.
	// Stability: development
	OpenAIRequestServiceTierDefault = OpenAIRequestServiceTierKey.String("default")
)
