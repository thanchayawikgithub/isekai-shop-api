package models

type (
	Item struct {
		ID          uint64 `json:id`
		Name        string `json:name`
		Description string `json:description`
		Picture     string `json:picture`
		Price       uint   `json:price`
	}
)
