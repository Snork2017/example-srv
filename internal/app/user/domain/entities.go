package domain

type ExampleData struct {
	Field  string
	Field2 string
}

const (
	ErrXXX = errors.New("frewfewrf")
)

func NewExampleData(field field2) error {
	if !field {
		return ErrXXX
	}
}