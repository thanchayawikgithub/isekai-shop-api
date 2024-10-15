package exceptions

import "fmt"

type InventoryFilling struct {
	PlayerID string
	ItemID   uint64
}

func (e *InventoryFilling) Error() string {
	return fmt.Sprintf("Inventory filling for playerID: %s and itemID: %d failed", e.PlayerID, e.ItemID)
}
