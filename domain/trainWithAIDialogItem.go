package domain

// TODO: FIX MIGRATIONS

type TrainWithAIDialogItem struct {
	ID                  uint
	OrderDialog         uint16 // (?)
	Dialog              string
	IsFromUser          bool // TODO: FIX MIGRATIONS
	TrainWithAIDialogID uint
}
