package dto

type DialogLessonDialog struct {
	OrderDialog uint8  `json:"order_dialog"`
	Dialog      string `json:"dialog"`
	Meaning     string `json:"meaning"`
	IsForUser   string `json:"is_for_user"`
	// TODO: Add sound UK and US
}

type DialogLessonMisspronounciatedWord struct {
	Word         string  `json:"word"`
	Accuracy     float32 `json:"accuracy"`
	UserShouldDo string  `json:"user_should_do"`
}

type DialogLessonGetStatusOKResponse struct {
	Dialogs []DialogLessonDialog `json:"dialogs"`
}

type DialogLessonPostStatusOKResponse struct {
	Dialog                 string                              `json:"dialog"`
	MisspronounciatedWords []DialogLessonMisspronounciatedWord `json:"misspronounciated_words"`
	OverallAccuracy        float32                             `json:"overall_accuracy"`
	CanContinue            bool                                `json:"can_continue"`
}
