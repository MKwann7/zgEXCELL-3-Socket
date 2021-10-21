package dtos

type Messages struct {
}

func (messages *Messages) getNotifications() string {

	return ""
}

type Message struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}
