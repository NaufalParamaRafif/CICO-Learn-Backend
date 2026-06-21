package domain

import "context"

type DialogLessonDialogItem struct {
	ID             uint   `db:"id"`
	Order          uint16 `db:"dialog_order"` // TODO:
	Dialog         string `db:"dialog"`
	DialogLessonID uint   `db:"dialog_lesson_id"`
}

type DialogLessonDialogItemRepository interface {
	GetByID(ctx context.Context, dialogLessonDialogItemID uint) (DialogLessonDialogItem, error)
	Insert(ctx context.Context, dialogLessonDialogItem *DialogLessonDialogItem) error
	Update(ctx context.Context, dialogLessonDialogItem *DialogLessonDialogItem) error
	Delete(ctx context.Context, dialogLessonDialogItemID uint) error
}
