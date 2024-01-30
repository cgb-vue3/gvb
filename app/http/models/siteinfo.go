package models

type WebInfo struct {
	Name        string `json:"name"`
	Addr        string `json:"addr"`
	Title       string `json:"title"`
	QQImage     string `json:"QQImage"`
	WeChatImage string `json:"weChatImage"`
	Email       string `json:"email"`
	GitHubUrl   string `json:"gitHubUrl"`
}
