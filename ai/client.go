package ai

type MessageType string

const (
	MessageTypeUser      MessageType = "user"
	MessageTypeSystem    MessageType = "system"
	MessageTypeAi        MessageType = "ai"
	MessageTypeTool      MessageType = "tool"
	MessageTypeUndefined MessageType = "undefined"
)

type ModelSize string

const (
	ModelS  ModelSize = "s"
	ModelM  ModelSize = "m"
	ModelL  ModelSize = "l"
	ModelXL ModelSize = "xl"
)

type Message interface {
	Content() string
	Type() MessageType
}

type Dialog interface {
	AppendUserMessage(message string)
	AppendSystemMessage(message string)
	GetMessages() []Message
	GetLastMessage() Message

	SetModelSize(size ModelSize) error
	SetTools(tools []Tool) error

	Duplicate() Dialog
}

type Client interface {
	NewDialog() Dialog
	RequestCompletion(dialog Dialog) error
	Query(query string, dialog Dialog) error
}
