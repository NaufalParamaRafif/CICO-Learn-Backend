package response

type UserDataResponse struct {
	ID                 uint     `json:"id"`
	FirstName          string   `json:"first_name"`
	LastName           string   `json:"last_name"`
	Username           string   `json:"username"`
	Email              string   `json:"email"`
	Address            string   `json:"address"`
	Gender             string   `json:"gender"`
	ProfilePictureURL  string   `json:"profile_picture_url"`
	BackgroundImageURL string   `json:"background_image_url"`
	Level              string   `json:"level"`
	Language           string   `json:"language"`
	Target             string   `json:"target"`
	DailyReminder      []string `json:"daily_reminder"`
	TimeTarget         string   `json:"time_target"`
}

type BackgroundImageResponse struct {
	Message            string `json:"message"`
	BackgroundImageURL string `json:"background_image_url"`
}

type ProfilePictureResponse struct {
	Message           string `json:"message"`
	ProfilePictureURL string `json:"profile_picture_url"`
}
