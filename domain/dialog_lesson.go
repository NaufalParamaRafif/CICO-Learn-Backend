package domain

type DialogLesson struct {
	ID          uint
	Topic       string
	Description string
	TotalDialog uint16
	Dialogues   []DialogLessonDialogItem
}
