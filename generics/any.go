package generics

//go:generate bash -c "go run ../generator/builtins/*.go > builtins.gen.go"

type (
	// Any is any type
	Any interface{}

	// T1 is a generic type
	T1 interface{}
)
