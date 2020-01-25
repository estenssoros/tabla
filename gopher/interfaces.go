package gopher

type Dialect interface {
	DropIfExists(*GoStruct) string
	Create(*GoStruct) (string, error)
}
