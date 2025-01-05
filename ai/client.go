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

	SetModelSize(size ModelSize) error
	SetTools(tools []Tool) error
}

type Client interface {
	NewDialog() Dialog
	Query(query string, dialog Dialog) error
}
