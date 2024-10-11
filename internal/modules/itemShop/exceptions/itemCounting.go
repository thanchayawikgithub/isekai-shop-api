package exceptions

type ItemCounting struct{}

func (e *ItemCounting) Error() string {
	return "Failed to counting item"
}
