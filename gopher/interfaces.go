package gopher

// Dialect interface for sql dialects
type Dialect interface {
	DropIfExists(*GoStruct) string
	Create(*GoStruct) (string, error)
}
