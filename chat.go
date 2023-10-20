package staringos

import (
	"github.com/darwinOrg/go-openai"
)

const (
	host             = "https://api.ai.staringos.com"
	chatByCorpusUrl  = host + "/docs-ai/chatByCorpus"
	messageUrlPrefix = host + "/message?id="
)

type ChatByCorpusRequest struct {
	Model        string                         `json:"model,omitempty"`
	Prompt       []openai.ChatCompletionMessage `json:"prompt,omitempty"`
	Temperature  float32                        `json:"temperature,omitempty"`
	TopP         float32                        `json:"top_p,omitempty"`
	Functions    []openai.FunctionDefinition    `json:"functions,omitempty"`
	FunctionCall any                            `json:"function_call,omitempty"`
}

type MessageIdResponse struct {
	Id string `json:"id"`
}

type MessageResponse struct {
	Content      string `json:"content,omitempty"`
	IsFinish     bool   `json:"isFinish,omitempty"`
	FunctionCall any    `json:"function_call,omitempty"`
	Model        string `json:"model,omitempty"`
	AppId        int    `json:"appId,omitempty"`
	At           int64  `json:"at,omitempty"`
}
