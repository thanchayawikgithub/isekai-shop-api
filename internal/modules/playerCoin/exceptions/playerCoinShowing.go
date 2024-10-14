package exceptions

type PlayerCoinShowing struct{}

func (e *PlayerCoinShowing) Error() string {
	return "Failed to show player coin"
}
