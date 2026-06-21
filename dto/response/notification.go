package response

type Notification struct {
	Title         string `json:"title"`
	Text          string `json:"text"`
	HasBeenOpened bool   `json:"has_been_opened"`
	Emoticon      string `json:"emoticon"`
	SendAt        string `json:"send_at"`
}

type NotificationGetStatusOKResponse struct {
	Notifications []Notification `json:"notifications"`
}
