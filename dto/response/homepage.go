package response

type Announcement struct {
	Emoticon        string `json:"emoticon"`
	Text            string `json:"text"`
	HighlightedWord string `json:"highlighted_word"`
}

type Streak struct {
	DaysStreak    []string `json:"days"`
	LessonLearned uint16   `json:"lesson_learned"`
}

type RecentLessons struct {
	Lessons []string `json:"lessons"`
}

type HomepageGetStatusOKResponse struct {
	Announcement  Announcement
	Streak        Streak
	RecentLessons RecentLessons
}
