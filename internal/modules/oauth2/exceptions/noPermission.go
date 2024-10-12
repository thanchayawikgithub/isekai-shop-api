package exceptions

type NoPermission struct{}

func (e *NoPermission) Error() string {
	return "no permission"
}
