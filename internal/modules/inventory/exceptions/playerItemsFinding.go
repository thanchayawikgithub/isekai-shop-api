package exceptions

import "fmt"

type PlayerItemsFinding struct {
	PlayerID string
}

func (e *PlayerItemsFinding) Error() string {
	return fmt.Sprintf("Finding player items for playerID: %s failed", e.PlayerID)
}
