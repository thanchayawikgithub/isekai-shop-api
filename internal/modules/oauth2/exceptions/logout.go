package exceptions

type Logout struct{}

func (e *Logout) Error() string {
	return "logout failed"
}
