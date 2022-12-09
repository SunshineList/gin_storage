package request

type UserInfo struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
}

