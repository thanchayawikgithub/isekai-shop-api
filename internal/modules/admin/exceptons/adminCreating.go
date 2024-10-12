package exceptons

import "fmt"

type AdminCreating struct {
	AdminID string
}

func (e *AdminCreating) Error() string {
	return fmt.Sprintf("Failed to creating adminID: %s", e.AdminID)
}
