package ai

type DataType string

const (
	Integer  DataType = "integer"
	Real     DataType = "real"
	String   DataType = "string"
	Boolean  DataType = "boolean"
	Datetime DataType = "datetime"
)

type ToolParameter struct {
	Name        string
	Description string
	Type        DataType
}

type Tool struct {
	Name        string
	Description string
	Parameters  []ToolParameter
	Function    func(map[string]any) (string, error)
}
