package gopher

type dialect interface {
	DropIfExists(*GoStruct) string
	Create(*GoStruct) string
}
