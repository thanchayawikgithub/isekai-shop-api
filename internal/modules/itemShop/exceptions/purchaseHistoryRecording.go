package exceptions

type PurchaseHistoryRecording struct{}

func (e *PurchaseHistoryRecording) Error() string {
	return "Failed to recoring purchase history"
}
