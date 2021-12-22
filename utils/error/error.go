package error

type Error struct {
	Name  string
	Error error
}

func NewError(name string, error error) *Error {
	return &Error{
		Name:  name,
		Error: error,
	}
}
