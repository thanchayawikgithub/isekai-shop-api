package exceptions

type Logout struct{}

func (e *Logout) Error() string {
	return "Logout Failed"
}
