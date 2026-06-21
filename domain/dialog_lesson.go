package domain

import "context"

type DialogLesson struct {
	ID          uint   `db:"id"`
	Topic       string `db:"topic"`
	Description string `db:"description"`
}

type DialogLessonRepository interface {
	GetByID(ctx context.Context, dialogLessonID uint) (DialogLesson, error)
	Insert(ctx context.Context, dialogLesson *DialogLesson) error
	Update(ctx context.Context, dialogLesson *DialogLesson) error
	Delete(ctx context.Context, dialogLessonID uint) error

	GetDialogItems(ctx context.Context, dialogLessonID uint) ([]DialogLessonDialogItem, error)

	GetUsers(ctx context.Context, dialogLessonID uint) ([]User, error)
	AddUser(ctx context.Context, userID uint) error
	RemoveUser(ctx context.Context, userID uint) error
}
