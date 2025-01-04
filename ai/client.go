package ai

type MessageType string

const (
	MessageTypeUser      MessageType = "user"
	MessageTypeAi        MessageType = "ai"
	MessageTypeTool      MessageType = "tool"
	MessageTypeUndefined MessageType = "undefined"
)

type Message interface {
	Content() string
	Type() MessageType
}

type Dialog interface {
	AppendUserMessage(message string)
	GetMessages() []Message
}

type Client interface {
	NewDialog() Dialog
	Query(query string, dialog Dialog) error

	SetTools(tools []Tool)
}
