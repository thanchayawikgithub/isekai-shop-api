package exceptons

import "fmt"

type PlayerCreating struct {
	PlayerID string
}

func (e *PlayerCreating) Error() string {
	return fmt.Sprintf("Failed to creating playerID: %s", e.PlayerID)
}
