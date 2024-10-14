package domain

type Event struct {
	OrderType  string `json:"orderType"`
	SessionId  string `json:"sessionId"`
	Card       string `json:"card"`
	EventDate  string `json:"eventDate"`
	WebsiteURL string `json:"websiteUrl"`
}
