package exceptions

type InvalidState struct{}

func (e *InvalidState) Error() string {
	return "Invalid state"
}
