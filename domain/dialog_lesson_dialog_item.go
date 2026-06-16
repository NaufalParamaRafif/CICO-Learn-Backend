package domain

import "context"

type DialogLessonDialogItem struct {
	ID             uint
	Order          uint16
	Dialog         string
	DialogLessonID uint
}

type DialogLessonDialogItemRepository interface {
	GetByID(ctx context.Context, dialogLessonDialogItemID uint) (DialogLessonDialogItem, error)
	Insert(ctx context.Context, dialogLessonDialogItem *DialogLessonDialogItem) error
	Update(ctx context.Context, dialogLessonDialogItem *DialogLessonDialogItem) error
	Delete(ctx context.Context, dialogLessonDialogItemID uint) error
}
