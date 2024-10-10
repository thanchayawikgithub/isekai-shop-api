package exceptions

type ItemListing struct {
}

func (e *ItemListing) Error() string {
	return "item lising failed"
}
